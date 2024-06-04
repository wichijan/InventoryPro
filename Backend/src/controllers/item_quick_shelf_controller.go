package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
	"github.com/wichijan/InventoryPro/src/utils"
)


type ItemQuickShelfControllerI interface {
	GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]model.ItemQuickShelf, *models.INVError)
	InsertItemInQuickShelf(itemQuickShelf *models.ItemQuickShelfInsert) *models.INVError // Check user amount & total amount

	MoveItemOutOfQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) *models.INVError // Insert in regular shelf
	ClearQuickShelf() *models.INVError // Return to regular shelf & clear quick shelf
}

type ItemQuickShelfController struct {
	ItemQuickShelfRepo repositories.ItemQuickShelfRepositoryI
	UserItemRepo repositories.UserItemRepositoryI
}

func (qsc *ItemQuickShelfController) GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]model.ItemQuickShelf, *models.INVError) {
	return qsc.ItemQuickShelfRepo.GetItemsInQuickShelf(quickShelfId)
}

func (qsc *ItemQuickShelfController) InsertItemInQuickShelf(itemQuickShelf *models.ItemQuickShelfInsert)  *models.INVError {
	tx, err := qsc.ItemQuickShelfRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	// Remove item from user => User_item_table
	currentQuantityOfUserItem, inv_error := qsc.UserItemRepo.GetQuantityFromUserItem(&itemQuickShelf.ItemID)
	if inv_error != nil {
		return inv_error
	}
	number := *currentQuantityOfUserItem - itemQuickShelf.Quantity
	if number <= 0 {
		inv_error := qsc.UserItemRepo.DeleteItemUser(tx, &itemQuickShelf.UserID, &itemQuickShelf.ItemID)
		if inv_error != nil {
			return inv_error
		}
	} else {
		inv_error := qsc.UserItemRepo.ReduceQuantityOfUserItem(tx, &itemQuickShelf.UserID, &itemQuickShelf.ItemID, &number)
		if inv_error != nil {
			return inv_error
		}
	}


	// Check if user has more than three items already in quick shelf
	items, inv_error := qsc.ItemQuickShelfRepo.GetItemsFromUserInQuickShelf(&itemQuickShelf.UserID)
	if inv_error != nil {
		return inv_error
	}
	if len(*items) >= utils.AMOUNT_OF_ITEMS_IN_QUICK_SHELF {
		return inv_errors.INV_QUICK_SHELF_USER_LIMIT_FULL
	}

	// is quick shelf full?
	items, inv_error = qsc.ItemQuickShelfRepo.GetItemsInQuickShelf(&itemQuickShelf.QuickShelfID)
	if inv_error != nil {
		return inv_error
	}
	if len(*items) >= utils.AMOUNT_OF_ITEMS_IN_QUICK_SHELF {
		return inv_errors.INV_QUICK_SHELF_FULL
	}

	// Check if item is already in quick shelf
	isInQuickShelf, inv_error := qsc.ItemQuickShelfRepo.CheckIfItemAlreadyInQuickShelf(&itemQuickShelf.ItemID, &itemQuickShelf.QuickShelfID)
	if inv_error != nil {
		return inv_error
	}

	quickShelfItem := model.ItemQuickShelf{
		QuickShelfID: itemQuickShelf.QuickShelfID.String(),
		UserID: itemQuickShelf.UserID.String(),
		ItemID: itemQuickShelf.ItemID.String(),
		Quantity: &itemQuickShelf.Quantity,
	}

	// Insert new oder update quantity of item in quick shelf
	if *isInQuickShelf {
		// Get Quantity
		quantity, inv_error := qsc.ItemQuickShelfRepo.GetQuantityOfItemInQuickShelf(&itemQuickShelf.ItemID, &itemQuickShelf.QuickShelfID)
		if inv_error != nil {
			return inv_error
		}
		newQuantity := *quantity + itemQuickShelf.Quantity
		quickShelfItem.Quantity = &newQuantity

		// Update 
		inv_error = qsc.ItemQuickShelfRepo.UpdateQuantityOfItemInQuickShelf(tx, &quickShelfItem)
		if inv_error != nil {
			return inv_error
		}
	} else {
		// Insert
		inv_error = qsc.ItemQuickShelfRepo.InsertNewItemInQuickShelf(tx, &quickShelfItem)
		if inv_error != nil {
			return inv_error
		}
	}


	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	return nil	
}