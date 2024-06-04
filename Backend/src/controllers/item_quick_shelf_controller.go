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
	GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]models.ItemQuickShelfInsert, *models.INVError)
	InsertItemInQuickShelf(itemQuickShelf *models.ItemQuickShelfInsert) *models.INVError // Check user amount & total amount

	RemoveItemFromQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) *models.INVError // Insert in regular shelf
	ClearQuickShelf(quickShelfId *uuid.UUID) *models.INVError                                                   // Return to regular shelf & clear quick shelf
}

type ItemQuickShelfController struct {
	ItemQuickShelfRepo repositories.ItemQuickShelfRepositoryI
	UserItemRepo       repositories.UserItemRepositoryI
	ItemRepo           repositories.ItemRepositoryI
	ItemsInShelfRepo   repositories.ItemInShelveRepositoryI
}

func (qsc *ItemQuickShelfController) GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]models.ItemQuickShelfInsert, *models.INVError) {
	return qsc.ItemQuickShelfRepo.GetItemsInQuickShelf(quickShelfId)
}

func (qsc *ItemQuickShelfController) InsertItemInQuickShelf(itemQuickShelf *models.ItemQuickShelfInsert) *models.INVError {
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
	userItems, inv_error := qsc.ItemQuickShelfRepo.GetItemsFromUserInQuickShelf(&itemQuickShelf.UserID)
	if inv_error != nil {
		return inv_error
	}
	if len(*userItems) >= int(utils.MAX_AMOUNT_OF_ITEMS_FOR_USER_IN_QUICK_SHELF) {
		return inv_errors.INV_QUICK_SHELF_USER_LIMIT_FULL
	}

	// is quick shelf full?
	allItems, inv_error := qsc.ItemQuickShelfRepo.GetItemsInQuickShelf(&itemQuickShelf.QuickShelfID)
	if inv_error != nil {
		return inv_error
	}
	if len(*allItems) >= utils.MAX_AMOUNT_OF_ITEMS_IN_QUICK_SHELF {
		return inv_errors.INV_QUICK_SHELF_FULL
	}

	// Check if item is already in quick shelf
	isInQuickShelf, inv_error := qsc.ItemQuickShelfRepo.CheckIfItemAlreadyInQuickShelf(&itemQuickShelf.ItemID, &itemQuickShelf.QuickShelfID)
	if inv_error != nil {
		return inv_error
	}

	quickShelfItem := model.ItemQuickShelf{
		QuickShelfID: itemQuickShelf.QuickShelfID.String(),
		UserID:       itemQuickShelf.UserID.String(),
		ItemID:       itemQuickShelf.ItemID.String(),
		Quantity:     &itemQuickShelf.Quantity,
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

func (qsc *ItemQuickShelfController) ClearQuickShelf(quickShelfId *uuid.UUID) *models.INVError {
	tx, err := qsc.ItemQuickShelfRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	// Get all items in quick shelf
	quickShelfItems, inv_error := qsc.ItemQuickShelfRepo.GetItemsInQuickShelf(quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	// Insert all items in regular shelf
	for _, quickShelfItem := range *quickShelfItems {
		item, inv_error := qsc.ItemRepo.GetItemById(&quickShelfItem.ItemID)
		if inv_error != nil {
			return inv_error
		}

		// Get Quantity in regular shelf
		quantity, inv_error := qsc.ItemsInShelfRepo.GetQuantityInShelve(&quickShelfItem.ItemID)
		if inv_error != nil {
			return inv_error
		}
		newQuantity := *quantity + quickShelfItem.Quantity

		// Update Quantity in regular shelf
		inv_error = qsc.ItemsInShelfRepo.UpdateItemInShelve(tx, &model.ItemsInShelf{
			ItemID:   quickShelfItem.ItemID.String(),
			ShelfID:  item.RegularShelfId,
			Quantity: &newQuantity,
		})
		if inv_error != nil {
			return inv_error
		}
	}
	// Clear quick shelf
	inv_error = qsc.ItemQuickShelfRepo.ClearQuickShelf(tx, quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	return nil
}

func (qsc *ItemQuickShelfController) RemoveItemFromQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) *models.INVError {
	tx, err := qsc.ItemQuickShelfRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	// Get Quantity in quick shelf
	quantity, inv_error := qsc.ItemQuickShelfRepo.GetQuantityOfItemInQuickShelf(itemId, quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	// Get Item
	item, inv_error := qsc.ItemRepo.GetItemById(itemId)
	if inv_error != nil {
		return inv_error
	}

	// Get Quantity in regular shelf
	quantityInRegularShelf, inv_error := qsc.ItemsInShelfRepo.GetQuantityInShelve(itemId)
	if inv_error != nil {
		return inv_error
	}
	newQuantity := *quantityInRegularShelf + *quantity

	// Update Quantity in regular shelf
	inv_error = qsc.ItemsInShelfRepo.UpdateItemInShelve(tx, &model.ItemsInShelf{
		ItemID:   itemId.String(),
		ShelfID:  item.RegularShelfId,
		Quantity: &newQuantity,
	})
	if inv_error != nil {
		return inv_error
	}

	// Remove item from quick shelf
	inv_error = qsc.ItemQuickShelfRepo.RemoveItemFromQuickShelf(tx, itemId, quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	return nil
}
