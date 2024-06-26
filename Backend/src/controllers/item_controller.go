package controllers

import (
	"log"
	"os"
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

	GetBookById(itemId *uuid.UUID) (*model.Books, *models.INVError)
	GetSingleObjectById(itemId *uuid.UUID) (*model.SingleObject, *models.INVError)
	GetSetOfObjectsById(itemId *uuid.UUID) (*model.SetsOfObjects, *models.INVError)

	CreateItemWithBook(item *models.ItemCreateWithBook) (*uuid.UUID, *models.INVError)
	CreateItemWithSingleObject(item *models.ItemCreateWithSingleObject) (*uuid.UUID, *models.INVError)
	CreateItemWithSetOfObject(item *models.ItemCreateWithSetOfObject) (*uuid.UUID, *models.INVError)

	UpdateItemWithBook(item *models.ItemUpdateWithBook) *models.INVError
	UpdateItemWithSingleObject(item *models.ItemUpdateWithSingleObject) *models.INVError
	UpdateItemWithSetOfObject(item *models.ItemUpdateWithSetsOfObjects) *models.INVError

	DeleteItem(itemId *uuid.UUID) *models.INVError
	AddKeywordToItem(itemKeyword models.ItemWithKeywordName) *models.INVError
	RemoveKeywordFromItem(itemKeyword models.ItemWithKeywordName) *models.INVError
	AddSubjectToItem(itemSubject models.ItemWithSubjectName) *models.INVError
	RemoveSubjectFromItem(itemSubject models.ItemWithSubjectName) *models.INVError

	BorrowItem(itemReserve models.ItemBorrowCreate) *models.INVError
	ReturnItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError

	MoveItemRequest(itemMove models.ItemMove) (*uuid.UUID, *models.INVError)
	MoveItemAccepted(transferAccept models.TransferAccept) (*models.TransferRequestSelect, *models.INVError)
	GetTransferRequestByUserId(userId uuid.UUID) (*[]models.TransferRequestSelect, *models.INVError)

	UploadItemImage(itemId *uuid.UUID) (*uuid.UUID, *models.INVError)
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
	TransactionRepo     repositories.TransactionRepositoryI
	TransferRequestRepo repositories.TransferRequestRepositoryI
	ShelveRepo          repositories.ShelveRepositoryI

	ItemsQuickShelfRepo repositories.ItemQuickShelfRepositoryI

	BookRepo         repositories.BookRepositoryI
	SingleObjectRepo repositories.SingleObjectRepositoryI
	SetOfObjectsRepo repositories.SetsOfObjectsRepositoryI
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

func (ic *ItemController) GetBookById(itemId *uuid.UUID) (*model.Books, *models.INVError) {
	book, inv_error := ic.BookRepo.GetBookById(itemId)
	if inv_error != nil {
		return nil, inv_error
	}

	return book, nil
}

func (ic *ItemController) GetSingleObjectById(itemId *uuid.UUID) (*model.SingleObject, *models.INVError) {
	singleObject, inv_error := ic.SingleObjectRepo.GetSingleObjectById(itemId)
	if inv_error != nil {
		return nil, inv_error
	}

	return singleObject, nil
}

func (ic *ItemController) GetSetOfObjectsById(itemId *uuid.UUID) (*model.SetsOfObjects, *models.INVError) {
	setOfObjects, inv_error := ic.SetOfObjectsRepo.GetSetsOfObjectsById(itemId)
	if inv_error != nil {
		return nil, inv_error
	}

	return setOfObjects, nil
}

func (ic *ItemController) CreateItemWithBook(item *models.ItemCreateWithBook) (*uuid.UUID, *models.INVError) {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check if shelf id exists
	inv_error := ic.ShelveRepo.CheckIfShelveExists(&item.RegularShelfId)
	if inv_error != nil {
		return nil, inv_error
	}

	var pureItem model.Items
	pureItem.Name = &item.Name
	pureItem.ItemTypes = item.ItemTypeName
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

	// Create Book
	inv_error = ic.BookRepo.CreateBook(tx, &model.Books{
		ItemID:    id.String(),
		Isbn:      item.Isbn,
		Author:    &item.Author,
		Publisher: &item.Publisher,
		Edition:   &item.Edition,
	})
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
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return id, nil
}

func (ic *ItemController) CreateItemWithSingleObject(item *models.ItemCreateWithSingleObject) (*uuid.UUID, *models.INVError) {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check if shelf id exists
	inv_error := ic.ShelveRepo.CheckIfShelveExists(&item.RegularShelfId)
	if inv_error != nil {
		return nil, inv_error
	}

	var pureItem model.Items
	pureItem.Name = &item.Name
	pureItem.ItemTypes = item.ItemTypeName
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

	// Create SingleObject
	inv_error = ic.SingleObjectRepo.CreateSingleObject(tx, &model.SingleObject{
		ItemID: id.String(),
	})
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
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return id, nil
}

func (ic *ItemController) CreateItemWithSetOfObject(item *models.ItemCreateWithSetOfObject) (*uuid.UUID, *models.INVError) {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check if shelf id exists
	inv_error := ic.ShelveRepo.CheckIfShelveExists(&item.RegularShelfId)
	if inv_error != nil {
		return nil, inv_error
	}

	var pureItem model.Items
	pureItem.Name = &item.Name
	pureItem.ItemTypes = item.ItemTypeName
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

	// Create SetOfObjects
	inv_error = ic.SetOfObjectsRepo.CreateSetsOfObjects(tx, &model.SetsOfObjects{
		ItemID:        id.String(),
		TotalObjects:  item.TotalObjects,
		UsefulObjects: item.UsefulObjects,
		BrokenObjects: item.BrokenObjects,
		LostObjects:   item.LostObjects,
	})
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
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return id, nil
}

func (ic *ItemController) UpdateItemWithBook(item *models.ItemUpdateWithBook) *models.INVError {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check if shelf id exists
	convertedUUID, inv_error := utils.ConvertStringToUUID(*item.RegularShelfID)
	if inv_error != nil {
		return inv_error
	}
	inv_error = ic.ShelveRepo.CheckIfShelveExists(convertedUUID)
	if inv_error != nil {
		return inv_error
	}

	var pureItem model.Items
	pureItem.ID = item.ID
	pureItem.Name = item.Name
	pureItem.ItemTypes = item.ItemTypes
	pureItem.RegularShelfID = item.RegularShelfID
	pureItem.HintText = item.HintText
	pureItem.Description = item.Description
	pureItem.ClassOne = item.ClassOne
	pureItem.ClassTwo = item.ClassTwo
	pureItem.ClassThree = item.ClassThree
	pureItem.ClassFour = item.ClassFour
	pureItem.Damaged = item.Damaged
	pureItem.DamagedDescription = item.DamagedDescription

	inv_error = ic.ItemRepo.UpdateItem(tx, &pureItem)
	if inv_error != nil {
		return inv_error
	}

	inv_error = ic.ItemInShelveRepo.UpdateItemInShelve(tx, &model.ItemsInShelf{
		ItemID:   item.ID,
		ShelfID:  *item.RegularShelfID,
		Quantity: &item.QuantityInShelf,
	})
	if inv_error != nil {
		return inv_error
	}

	book := model.Books{
		ItemID:    item.ID,
		Isbn:      item.Isbn,
		Author:    item.Author,
		Publisher: item.Publisher,
		Edition:   item.Edition,
	}

	inv_error = ic.BookRepo.UpdateBook(tx, &book)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) UpdateItemWithSingleObject(item *models.ItemUpdateWithSingleObject) *models.INVError {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check if shelf id exists
	convertedUUID, inv_error := utils.ConvertStringToUUID(*item.RegularShelfID)
	if inv_error != nil {
		return inv_error
	}
	inv_error = ic.ShelveRepo.CheckIfShelveExists(convertedUUID)
	if inv_error != nil {
		return inv_error
	}

	var pureItem model.Items
	pureItem.ID = item.ID
	pureItem.Name = item.Name
	pureItem.ItemTypes = item.ItemTypes
	pureItem.RegularShelfID = item.RegularShelfID
	pureItem.HintText = item.HintText
	pureItem.Description = item.Description
	pureItem.ClassOne = item.ClassOne
	pureItem.ClassTwo = item.ClassTwo
	pureItem.ClassThree = item.ClassThree
	pureItem.ClassFour = item.ClassFour
	pureItem.Damaged = item.Damaged
	pureItem.DamagedDescription = item.DamagedDescription

	inv_error = ic.ItemRepo.UpdateItem(tx, &pureItem)
	if inv_error != nil {
		return inv_error
	}

	inv_error = ic.ItemInShelveRepo.UpdateItemInShelve(tx, &model.ItemsInShelf{
		ItemID:   item.ID,
		ShelfID:  *item.RegularShelfID,
		Quantity: &item.QuantityInShelf,
	})
	if inv_error != nil {
		return inv_error
	}

	// If singleObject has more attributes we need to update them here

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) UpdateItemWithSetOfObject(item *models.ItemUpdateWithSetsOfObjects) *models.INVError {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check if shelf id exists
	convertedUUID, inv_error := utils.ConvertStringToUUID(*item.RegularShelfID)
	if inv_error != nil {
		return inv_error
	}
	inv_error = ic.ShelveRepo.CheckIfShelveExists(convertedUUID)
	if inv_error != nil {
		return inv_error
	}

	var pureItem model.Items
	pureItem.ID = item.ID
	pureItem.Name = item.Name
	pureItem.ItemTypes = item.ItemTypes
	pureItem.RegularShelfID = item.RegularShelfID
	pureItem.HintText = item.HintText
	pureItem.Description = item.Description
	pureItem.ClassOne = item.ClassOne
	pureItem.ClassTwo = item.ClassTwo
	pureItem.ClassThree = item.ClassThree
	pureItem.ClassFour = item.ClassFour
	pureItem.Damaged = item.Damaged
	pureItem.DamagedDescription = item.DamagedDescription

	inv_error = ic.ItemRepo.UpdateItem(tx, &pureItem)
	if inv_error != nil {
		return inv_error
	}

	inv_error = ic.ItemInShelveRepo.UpdateItemInShelve(tx, &model.ItemsInShelf{
		ItemID:   item.ID,
		ShelfID:  *item.RegularShelfID,
		Quantity: &item.QuantityInShelf,
	})
	if inv_error != nil {
		return inv_error
	}

	setOfObjects := model.SetsOfObjects{
		ItemID:        item.ID,
		TotalObjects:  item.TotalObjects,
		UsefulObjects: item.UsefulObjects,
		BrokenObjects: item.BrokenObjects,
		LostObjects:   item.LostObjects,
	}

	inv_error = ic.SetOfObjectsRepo.UpdateSetsOfObjects(tx, &setOfObjects)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) DeleteItem(itemId *uuid.UUID) *models.INVError {
	tx, err := ic.ItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check if item exists in other tables
	// Book
	ok, inv_error := ic.BookRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.BookRepo.DeleteBook(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// SingleObject
	ok, inv_error = ic.SingleObjectRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.SingleObjectRepo.DeleteSingleObject(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// SetOfObjects
	ok, inv_error = ic.SetOfObjectsRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.SetOfObjectsRepo.DeleteSetsOfObjects(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// ItemKeyword
	ok, inv_error = ic.ItemKeywordRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.ItemKeywordRepo.DeleteKeywordsForItem(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// ItemSubject
	ok, inv_error = ic.ItemSubjectRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.ItemSubjectRepo.DeleteSubjectsForItem(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// ItemsInShelve
	ok, inv_error = ic.ItemInShelveRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.ItemInShelveRepo.DeleteItemsInShelve(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// Items Quick Shelf
	ok, inv_error = ic.ItemsQuickShelfRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.ItemsQuickShelfRepo.DeleteItemsQuickShelf(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// UserItem
	ok, inv_error = ic.UserItemRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.UserItemRepo.DeleteItemUsers(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// Reservation
	ok, inv_error = ic.ReservationRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.ReservationRepo.DeleteReservationForItems(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// Transfer Request
	ok, inv_error = ic.TransferRequestRepo.CheckIfItemIdExists(itemId)
	if inv_error != nil {
		return inv_error
	} else if ok {
		if inv_error := ic.TransferRequestRepo.DeleteTransferRequest(tx, itemId); inv_error != nil {
			return inv_error
		}
	}

	// Item itself
	if inv_error := ic.ItemRepo.DeleteItem(tx, itemId); inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return nil
}

func (ic *ItemController) AddKeywordToItem(itemKeyword models.ItemWithKeywordName) *models.INVError {
	tx, err := ic.ItemKeywordRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) RemoveKeywordFromItem(itemKeyword models.ItemWithKeywordName) *models.INVError {
	tx, err := ic.ItemKeywordRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) AddSubjectToItem(itemSubject models.ItemWithSubjectName) *models.INVError {
	tx, err := ic.ItemSubjectRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) RemoveSubjectFromItem(itemSubject models.ItemWithSubjectName) *models.INVError {
	tx, err := ic.ItemSubjectRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) UploadItemImage(itemId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	pictureId, inv_error := ic.ItemRepo.StoreItemPicture(tx, itemId)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error removing image")
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) BorrowItem(itemReserve models.ItemBorrowCreate) *models.INVError {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check items_in_shelve quantity
	quantityInShelve, inv_error := ic.ItemInShelveRepo.GetQuantityInShelve(&itemReserve.ItemID)
	if inv_error != nil {
		return inv_error
	}

	newQuantityInShelve := *quantityInShelve - itemReserve.Quantity
	if newQuantityInShelve < 0 {
		return inv_errors.INV_CONFLICT.WithDetails("Not enough quantity in shelve")
	}

	// insert into user_items
	var pureItemBorrow models.ItemBorrow
	pureItemBorrow.ItemID = itemReserve.ItemID.String()
	pureItemBorrow.UserID = itemReserve.UserID.String()
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

	// Transaction Logging
	var transaction model.Transactions
	transaction.ItemID = itemReserve.ItemID.String()
	transaction.UserID = itemReserve.UserID.String()
	transaction.TransactionType = "borrow"
	transaction.TargetUserID = nil
	transaction.OriginUserID = nil

	// Add Transaction
	inv_error = ic.TransactionRepo.CreateTransaction(tx, &transaction)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) ReturnItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Get quantity
	quantityInShelve, inv_error := ic.ItemInShelveRepo.GetQuantityInShelve(itemId)
	if inv_error != nil {
		return inv_error
	}

	// Get quantity from user_items
	quantityReservedItem, inv_error := ic.UserItemRepo.GetQuantityFromUserItem(itemId, userId)
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

	// Transaction Logging
	var transaction model.Transactions
	transaction.ItemID = itemId.String()
	transaction.UserID = userId.String()
	transaction.TransactionType = "return"
	transaction.TargetUserID = nil
	transaction.OriginUserID = nil

	// Add Transaction
	inv_error = ic.TransactionRepo.CreateTransaction(tx, &transaction)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (ic *ItemController) MoveItemRequest(itemMove models.ItemMove) (*uuid.UUID, *models.INVError) {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	var transferRequest models.TransferRequestCreate
	transferRequest.ItemID = itemMove.ItemID.String()
	transferRequest.UserID = itemMove.UserID.String()
	transferRequest.TargetUserID = itemMove.NewUserID.String()

	transferRequestId, inv_error := ic.TransferRequestRepo.CreateTransferRequest(tx, &transferRequest)
	if inv_error != nil {
		return nil, inv_error
	}

	// Transaction Logging
	targetUserId := itemMove.NewUserID.String()
	originUserId := itemMove.UserID.String()

	var transaction model.Transactions
	transaction.ItemID = itemMove.ItemID.String()
	transaction.UserID = itemMove.UserID.String()
	transaction.TransactionType = "transfer_request"
	transaction.TargetUserID = &targetUserId
	transaction.OriginUserID = &originUserId

	// Add Transaction
	inv_error = ic.TransactionRepo.CreateTransaction(tx, &transaction)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return transferRequestId, nil
}

func (ic *ItemController) MoveItemAccepted(transferAccept models.TransferAccept) (*models.TransferRequestSelect, *models.INVError) {
	tx, err := ic.UserItemRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// GET Transfer Request
	transferRequest, inv_error := ic.TransferRequestRepo.GetTransferRequestById(*transferAccept.TransferRequestID)
	if inv_error != nil {
		return nil, inv_error
	}

	log.Print("TransferRequest ", transferRequest.TargetUserID)
	log.Print("TransferAccept ", transferAccept.UserId)
	if transferRequest.TargetUserID.String() != transferAccept.UserId.String() {
		return nil, inv_errors.INV_CONFLICT.WithDetails("User ID does not match the target user ID")
	}

	inv_error = ic.UserItemRepo.MoveItemToNewUser(tx, transferRequest.UserID, transferRequest.ItemID, transferRequest.TargetUserID)
	if inv_error != nil {
		return nil, inv_error
	}

	// Delete transfer request
	inv_error = ic.TransferRequestRepo.DeleteTransferRequest(tx, transferAccept.TransferRequestID)
	if inv_error != nil {
		return nil, inv_error
	}

	transactionDate := time.Now()
	targetUserId := transferRequest.TargetUserID.String()
	originUserId := transferRequest.UserID.String()

	var transaction model.Transactions
	transaction.ItemID = transferRequest.ItemID.String()
	transaction.UserID = transferRequest.UserID.String()
	transaction.TransactionType = "transfer_accepted"
	transaction.TransactionDate = &transactionDate
	transaction.TargetUserID = &targetUserId
	transaction.OriginUserID = &originUserId

	// Add Transaction
	inv_error = ic.TransactionRepo.CreateTransaction(tx, &transaction)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return transferRequest, nil
}

func (ic *ItemController) GetTransferRequestByUserId(userId uuid.UUID) (*[]models.TransferRequestSelect, *models.INVError) {
	return ic.TransferRequestRepo.GetTransferRequestByUserId(userId)
}
