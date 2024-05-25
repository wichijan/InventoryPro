package controllers

import (
	"time"

	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type ItemControllerI interface {
	GetItems() (*[]models.ItemWithEverything, *models.INVError)
	GetItemById(itemId *uuid.UUID) (*models.ItemWithEverything, *models.INVError)
	CreateItem(item *models.ItemWithStatus) (*uuid.UUID, *models.INVError)
	UpdateItem(item *models.ItemWithStatus) *models.INVError
	AddKeywordToItem(itemKeyword models.ItemWithKeywordName) *models.INVError
	RemoveKeywordFromItem(itemKeyword models.ItemWithKeywordName) *models.INVError
	AddSubjectToItem(itemSubject models.ItemWithSubjectName) *models.INVError
	RemoveSubjectFromItem(itemSubject models.ItemWithSubjectName) *models.INVError

	ReserveItem(itemReserve models.ItemReserve) *models.INVError
	CancelReserveItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError

	BorrowItem(itemReserve models.ItemBorrow) *models.INVError
	ReturnItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError
}

type ItemController struct {
	ItemRepo         repositories.ItemRepositoryI
	ItemInShelveRepo repositories.ItemInShelveRepositoryI
	ItemStatusRepo   repositories.ItemStatusRepositoryI
	UserItemRepo     repositories.UserItemRepositoryI
	KeywordRepo      repositories.KeywordRepositoryI
	SubjectRepo      repositories.SubjectRepositoryI
	ItemKeywordRepo  repositories.ItemKeywordRepositoryI
	ItemSubjectRepo  repositories.ItemSubjectRepositoryI
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

func (ic *ItemController) CreateItem(item *models.ItemWithStatus) (*uuid.UUID, *models.INVError) {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	var pureItem model.Items
	pureItem.ID = item.ID
	pureItem.Name = &item.Name
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

	return id, nil
}

func (ic *ItemController) UpdateItem(item *models.ItemWithStatus) *models.INVError {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	var pureItem model.Items
	pureItem.ID = item.ID
	pureItem.Name = &item.Name
	pureItem.Description = &item.Description
	pureItem.ClassOne = &item.ClassOne
	pureItem.ClassTwo = &item.ClassTwo
	pureItem.ClassThree = &item.ClassThree
	pureItem.ClassFour = &item.ClassFour
	pureItem.Damaged = &item.Damaged
	pureItem.DamagedDescription = &item.DamagedDesc

	inv_error := ic.ItemRepo.UpdateItem(tx, &pureItem)
	if inv_error != nil {
		return inv_error
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

	return nil
}

func (ic *ItemController) ReserveItem(itemReserve models.ItemReserve) *models.INVError {
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
	// Get status_id
	statusName := "Reserved"
	statusID, inv_error := ic.ItemStatusRepo.GetStatusIdByName(&statusName)
	if inv_error != nil {
		return inv_error
	}
	itemReserve.StatusID = statusID.String()
	itemReserve.ReserveDate = time.Now()

	inv_error = ic.UserItemRepo.ReserveItem(tx, &itemReserve)
	if inv_error != nil {
		return inv_error
	}

	// update items_in_shelve quantity
	inv_error = ic.ItemInShelveRepo.UpdateQuantityInShelve(tx, &itemReserve.ItemID, &newQuantityInShelve)
	if inv_error != nil {
		return inv_error
	}

	return nil
}

func (ic *ItemController) CancelReserveItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError {
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
	quantityReservedItem, inv_error := ic.UserItemRepo.GetQuantityFromReservedItem(itemId)
	if inv_error != nil {
		return inv_error
	}

	// Get status_id
	statusName := "Reserved"
	statusID, inv_error := ic.ItemStatusRepo.GetStatusIdByName(&statusName)
	if inv_error != nil {
		return inv_error
	}
	statusIdString := statusID.String()

	// Update user_items
	inv_error = ic.UserItemRepo.DeleteReserveItem(tx, userId, itemId, &statusIdString)
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

	return nil
}

func (ic *ItemController) BorrowItem(itemReserve models.ItemBorrow) *models.INVError {
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
	// Get status_id
	statusName := "Borrowed"
	statusID, inv_error := ic.ItemStatusRepo.GetStatusIdByName(&statusName)
	if inv_error != nil {
		return inv_error
	}
	itemReserve.StatusID = statusID.String()
	itemReserve.BorrowDate = time.Now()

	inv_error = ic.UserItemRepo.BorrowItem(tx, &itemReserve)
	if inv_error != nil {
		return inv_error
	}

	// update items_in_shelve quantity
	inv_error = ic.ItemInShelveRepo.UpdateQuantityInShelve(tx, &itemReserve.ItemID, &newQuantityInShelve)
	if inv_error != nil {
		return inv_error
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
	quantityReservedItem, inv_error := ic.UserItemRepo.GetQuantityFromReservedItem(itemId)
	if inv_error != nil {
		return inv_error
	}

	// Get status_id
	statusName := "Borrowed"
	statusID, inv_error := ic.ItemStatusRepo.GetStatusIdByName(&statusName)
	if inv_error != nil {
		return inv_error
	}
	statusIdString := statusID.String()

	// Update user_items
	inv_error = ic.UserItemRepo.ReturnItem(tx, userId, itemId, &statusIdString)
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

	return nil
}
