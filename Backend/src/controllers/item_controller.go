package controllers

import (
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
	"github.com/wichijan/InventoryPro/src/utils"
)

type ItemControllerI interface {
	GetItems() (*[]models.ItemWithEverything, *models.INVError)
	GetItemById(itemId *uuid.UUID) (*models.ItemWithEverything, *models.INVError)
	CreateItem(item *models.ItemCreate) (*uuid.UUID, *models.INVError)
	UpdateItem(item *models.ItemUpdate) *models.INVError
	AddKeywordToItem(itemKeyword models.ItemWithKeywordName) *models.INVError
	RemoveKeywordFromItem(itemKeyword models.ItemWithKeywordName) *models.INVError
	AddSubjectToItem(itemSubject models.ItemWithSubjectName) *models.INVError
	RemoveSubjectFromItem(itemSubject models.ItemWithSubjectName) *models.INVError

	BorrowItem(itemReserve models.ItemBorrowCreate) *models.INVError
	ReturnItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError

	MoveItemRequest(itemMove models.ItemMove) *models.INVError
	MoveItemAccepted(itemMove models.ItemMove) *models.INVError

	UploadImage(itemId *uuid.UUID) (*uuid.UUID, *models.INVError)
	GetImageIdFromItem(itemId *uuid.UUID) (*uuid.UUID, *models.INVError)
	RemoveImageIdFromItem(itemId *uuid.UUID) *models.INVError
}

type ItemController struct {
	ItemRepo            repositories.ItemRepositoryI
	ItemInShelveRepo    repositories.ItemInShelveRepositoryI
	UserItemRepo        repositories.UserItemRepositoryI
	KeywordRepo         repositories.KeywordRepositoryI
	SubjectRepo         repositories.SubjectRepositoryI
	ItemKeywordRepo     repositories.ItemKeywordRepositoryI
	ItemSubjectRepo     repositories.ItemSubjectRepositoryI
	ReservationRepo     repositories.ReservationRepositoryI
	ItemTypeRepo        repositories.ItemTypeRepositoryI
	TransactionRepo     repositories.TransactionRepositoryI
	TransferRequestRepo repositories.TransferRequestRepositoryI
	ShelveRepo          repositories.ShelveRepositoryI
}

func (ic *ItemController) GetItems() (*[]models.ItemWithEverything, *models.INVError) {
	items, inv_error := ic.ItemRepo.GetItems()
	if inv_error != nil {
		return nil, inv_error
	}

	return items, nil
}

func (ic *ItemController) GetItemById(itemId *uuid.UUID) (*models.ItemWithEverything, *models.INVError) {
	item, inv_error := ic.ItemRepo.GetItemById(itemId)
	if inv_error != nil {
		return nil, inv_error
	}

	return item, nil
}

func (ic *ItemController) CreateItem(item *models.ItemCreate) (*uuid.UUID, *models.INVError) {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	item.ItemTypeName = strings.ToLower(item.ItemTypeName)

	// Get item type id
	itemTypeId, inv_err := ic.ItemTypeRepo.GetItemTypesByName(&item.ItemTypeName)
	if inv_err != nil {
		return nil, inv_err
	}

	// Check if shelf id exists
	inv_error := ic.ShelveRepo.CheckIfShelveExists(&item.RegularShelfId)
	if inv_error != nil {
		return nil, inv_error
	}

	var pureItem model.Items
	pureItem.Name = &item.Name
	pureItem.ItemTypeID = &itemTypeId.ID
	pureItem.RegularShelfID = utils.GetStringPointer(item.RegularShelfId.String())
	pureItem.HintText = &item.HintText
	pureItem.Description = &item.Description
	pureItem.ClassOne = &item.ClassOne
	pureItem.ClassTwo = &item.ClassTwo
	pureItem.ClassThree = &item.ClassThree
	pureItem.ClassFour = &item.ClassFour
	pureItem.Damaged = &item.Damaged
	pureItem.DamagedDescription = &item.DamagedDesc

	id, inv_error := ic.ItemRepo.CreateItem(tx, &pureItem)
	if inv_error != nil {
		return nil, inv_error
	}

	inv_error = ic.ItemInShelveRepo.CreateItemInShelve(tx, &model.ItemsInShelf{
		ItemID:   id.String(),
		ShelfID:  item.RegularShelfId.String(),
		Quantity: &item.BaseQuantityInShelf,
	})
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return id, nil
}

func (ic *ItemController) UpdateItem(item *models.ItemUpdate) *models.INVError {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	item.ItemTypeName = strings.ToLower(item.ItemTypeName)
	itemTypeId, inv_err := ic.ItemTypeRepo.GetItemTypesByName(&item.ItemTypeName)
	if inv_err != nil {
		return inv_err
	}

	// Check if shelf id exists
	inv_error := ic.ShelveRepo.CheckIfShelveExists(&item.RegularShelfId)
	if inv_error != nil {
		return inv_error
	}

	var pureItem model.Items
	pureItem.ID = item.ID
	pureItem.Name = &item.Name
	pureItem.ItemTypeID = &itemTypeId.ID
	pureItem.RegularShelfID = utils.GetStringPointer(item.RegularShelfId.String())
	pureItem.HintText = &item.HintText
	pureItem.Description = &item.Description
	pureItem.ClassOne = &item.ClassOne
	pureItem.ClassTwo = &item.ClassTwo
	pureItem.ClassThree = &item.ClassThree
	pureItem.ClassFour = &item.ClassFour
	pureItem.Damaged = &item.Damaged
	pureItem.DamagedDescription = &item.DamagedDesc

	inv_error = ic.ItemRepo.UpdateItem(tx, &pureItem)
	if inv_error != nil {
		return inv_error
	}

	inv_error = ic.ItemInShelveRepo.UpdateItemInShelve(tx, &model.ItemsInShelf{
		ItemID:   item.ID,
		ShelfID:  item.RegularShelfId.String(),
		Quantity: &item.QuantityInShelf,
	})
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) AddKeywordToItem(itemKeyword models.ItemWithKeywordName) *models.INVError {
	tx, err := ic.ItemKeywordRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	keyword, inv_error := ic.KeywordRepo.GetKeywordByName(&itemKeyword.KeywordName)
	if inv_error != nil {
		return inv_error
	}

	itemKeywordWithID := models.ItemWithKeyword{
		ItemID:    itemKeyword.ItemID,
		KeywordID: keyword.ID,
	}

	inv_err := ic.ItemKeywordRepo.CheckIfKeywordAndItemExists(itemKeywordWithID)
	if inv_err != nil {
		return inv_err
	}

	inv_error = ic.ItemKeywordRepo.CreateKeywordForItem(tx, &itemKeywordWithID)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) RemoveKeywordFromItem(itemKeyword models.ItemWithKeywordName) *models.INVError {
	tx, err := ic.ItemKeywordRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	keyword, inv_error := ic.KeywordRepo.GetKeywordByName(&itemKeyword.KeywordName)
	if inv_error != nil {
		return inv_error
	}

	itemKeywordWithID := models.ItemWithKeyword{
		ItemID:    itemKeyword.ItemID,
		KeywordID: keyword.ID,
	}

	inv_error = ic.ItemKeywordRepo.DeleteKeywordForItem(tx, &itemKeywordWithID)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) AddSubjectToItem(itemSubject models.ItemWithSubjectName) *models.INVError {
	tx, err := ic.ItemSubjectRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	subject, inv_error := ic.SubjectRepo.GetSubjectByName(&itemSubject.SubjectName)
	if inv_error != nil {
		return inv_error
	}

	itemSubjectWithID := models.ItemWithSubject{
		ItemID:    itemSubject.ItemID,
		SubjectID: subject.ID,
	}

	inv_err := ic.ItemSubjectRepo.CheckIfSubjectAndItemExists(itemSubjectWithID)
	if inv_err != nil {
		return inv_err
	}

	inv_error = ic.ItemSubjectRepo.CreateSubjectForItem(tx, &itemSubjectWithID)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) RemoveSubjectFromItem(itemSubject models.ItemWithSubjectName) *models.INVError {
	tx, err := ic.ItemSubjectRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	subject, inv_error := ic.SubjectRepo.GetSubjectByName(&itemSubject.SubjectName)
	if inv_error != nil {
		return inv_error
	}

	itemSubjectWithID := models.ItemWithSubject{
		ItemID:    itemSubject.ItemID,
		SubjectID: subject.ID,
	}

	inv_error = ic.ItemSubjectRepo.DeleteSubjectForItem(tx, &itemSubjectWithID)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) UploadImage(itemId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	pictureId, inv_error := ic.ItemRepo.StoreItemPicture(tx, itemId)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return pictureId, nil
}

func (ic *ItemController) GetImageIdFromItem(itemId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	pictureId, inv_error := ic.ItemRepo.GetPictureIdFromItem(itemId)
	if inv_error != nil {
		return nil, inv_error
	}

	return pictureId, nil
}

func (ic *ItemController) RemoveImageIdFromItem(itemId *uuid.UUID) *models.INVError {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	pictureId, inv_error := ic.ItemRepo.GetPictureIdFromItem(itemId)
	if inv_error != nil {
		return inv_error
	}

	inv_error = ic.ItemRepo.RemovePictureIdFromItem(tx, itemId)
	if inv_error != nil {
		return inv_error
	}

	imageName := "./../uploads/" + pictureId.String() + ".jpeg"
	inv_err := os.Remove(imageName)
	if inv_err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) BorrowItem(itemReserve models.ItemBorrowCreate) *models.INVError {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	// Check items_in_shelve quantity
	itemId, inv_err := uuid.Parse(itemReserve.ItemID)
	if inv_err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	quantityInShelve, inv_error := ic.ItemInShelveRepo.GetQuantityInShelve(&itemId)
	if inv_error != nil {
		return inv_error
	}

	newQuantityInShelve := *quantityInShelve - itemReserve.Quantity
	if newQuantityInShelve < 0 {
		return inv_errors.INV_NOT_ENOUGH_QUANTITY
	}

	// insert into user_items
	var pureItemBorrow models.ItemBorrow
	pureItemBorrow.ItemID = itemReserve.ItemID
	pureItemBorrow.UserID = itemReserve.UserID
	pureItemBorrow.Quantity = itemReserve.Quantity
	pureItemBorrow.TransactionDate = time.Now()

	inv_error = ic.UserItemRepo.InsertUserItem(tx, &pureItemBorrow)
	if inv_error != nil {
		return inv_error
	}

	// update items_in_shelve quantity
	inv_error = ic.ItemInShelveRepo.UpdateQuantityInShelve(tx, &pureItemBorrow.ItemID, &newQuantityInShelve)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) ReturnItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	// Get quantity
	quantityInShelve, inv_error := ic.ItemInShelveRepo.GetQuantityInShelve(itemId)
	if inv_error != nil {
		return inv_error
	}

	// Get quantity from user_items
	quantityReservedItem, inv_error := ic.UserItemRepo.GetQuantityFromUserItem(itemId)
	if inv_error != nil {
		return inv_error
	}

	// Update user_items
	inv_error = ic.UserItemRepo.DeleteItemUser(tx, userId, itemId)
	if inv_error != nil {
		return inv_error
	}

	// update items_in_shelve quantity
	itemIdStr := itemId.String()
	newQuantityInShelve := *quantityReservedItem + *quantityInShelve
	inv_error = ic.ItemInShelveRepo.UpdateQuantityInShelve(tx, &itemIdStr, &newQuantityInShelve)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) MoveItemRequest(itemMove models.ItemMove) *models.INVError {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	var transferRequest models.TransferRequestCreate
	transferRequest.ItemID = itemMove.ItemID.String()
	transferRequest.UserID = itemMove.UserID.String()
	transferRequest.TargetUserID = itemMove.NewUserID.String()

	_, inv_error := ic.TransferRequestRepo.CreateTransferRequest(tx, &transferRequest)
	if inv_error != nil {
		return inv_error
	}

	// Transaction Logging
	transactionDate := time.Now()
	targetUserId := itemMove.NewUserID.String()
	originUserId := itemMove.UserID.String()

	var transaction model.Transactions
	transaction.ItemID = itemMove.ItemID.String()
	transaction.UserID = itemMove.UserID.String()
	transaction.TransactionType = "transfer_request"
	transaction.TransactionDate = &transactionDate
	transaction.TargetUserID = &targetUserId
	transaction.OriginUserID = &originUserId

	// Add Transaction
	inv_error = ic.TransactionRepo.CreateTransaction(tx, &transaction)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ic *ItemController) MoveItemAccepted(itemMove models.ItemMove) *models.INVError {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	inv_error := ic.UserItemRepo.MoveItemToNewUser(tx, &itemMove.UserID, &itemMove.NewUserID, &itemMove.ItemID)
	if inv_error != nil {
		return inv_error
	}

	transactionDate := time.Now()
	targetUserId := itemMove.NewUserID.String()
	originUserId := itemMove.UserID.String()

	var transaction model.Transactions
	transaction.ItemID = itemMove.ItemID.String()
	transaction.UserID = itemMove.UserID.String()
	transaction.TransactionType = "transfer_accepted"
	transaction.TransactionDate = &transactionDate
	transaction.TargetUserID = &targetUserId
	transaction.OriginUserID = &originUserId

	// Add Transaction
	inv_error = ic.TransactionRepo.CreateTransaction(tx, &transaction)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}
