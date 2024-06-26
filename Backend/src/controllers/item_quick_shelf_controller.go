package controllers

import (
	"fmt"

	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
	"github.com/wichijan/InventoryPro/src/utils"
)

type ItemQuickShelfControllerI interface {
	GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]models.GetQuickShelf, *models.INVError)
	InsertItemInQuickShelf(itemQuickShelf *models.ItemQuickShelfInsert) *models.INVError // Check user amount & total amount

	RemoveItemFromQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) *models.INVError // Insert in regular shelf
	ClearQuickShelf(quickShelfId *uuid.UUID) *models.INVError                             // Return to regular shelf & clear quick shelf
}

type ItemQuickShelfController struct {
	ItemQuickShelfRepo repositories.ItemQuickShelfRepositoryI
	UserItemRepo       repositories.UserItemRepositoryI
	ItemRepo           repositories.ItemRepositoryI
	ItemsInShelfRepo   repositories.ItemInShelveRepositoryI
}

func (iqsc *ItemQuickShelfController) GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]models.GetQuickShelf, *models.INVError) {
	return iqsc.ItemQuickShelfRepo.GetItemsInQuickShelf(quickShelfId)
}

func (iqsc *ItemQuickShelfController) InsertItemInQuickShelf(itemQuickShelf *models.ItemQuickShelfInsert) *models.INVError {
	tx, err := iqsc.ItemQuickShelfRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Remove item from user => User_item_table
	currentQuantityOfUserItem, inv_error := iqsc.UserItemRepo.GetQuantityFromUserItem(&itemQuickShelf.ItemID, &itemQuickShelf.UserID)
	if inv_error != nil {
		return inv_error
	}

	number := *currentQuantityOfUserItem - itemQuickShelf.Quantity
	if number < 0 {
		return inv_errors.INV_CONFLICT.WithDetails("Not enough quantity")
	} else if number == 0 {
		inv_error := iqsc.UserItemRepo.DeleteItemUser(tx, &itemQuickShelf.UserID, &itemQuickShelf.ItemID)
		if inv_error != nil {
			return inv_error
		}
	} else {
		inv_error := iqsc.UserItemRepo.ReduceQuantityOfUserItem(tx, &itemQuickShelf.UserID, &itemQuickShelf.ItemID, &number)
		if inv_error != nil {
			return inv_error
		}
	}

	// Check if user has more than three items already in quick shelf
	userItems, inv_error := iqsc.ItemQuickShelfRepo.GetItemsFromUserInQuickShelf(&itemQuickShelf.UserID)
	if inv_error != nil {
		return inv_error
	}
	if len(*userItems) >= int(utils.MAX_AMOUNT_OF_ITEMS_FOR_USER_IN_QUICK_SHELF) {
		return inv_errors.INV_CONFLICT.WithDetails(fmt.Sprintf("User has reached the limit of items in quick shelf. Limit is %v", utils.MAX_AMOUNT_OF_ITEMS_FOR_USER_IN_QUICK_SHELF))
	}

	// is quick shelf full?
	allItems, inv_error := iqsc.ItemQuickShelfRepo.GetItemsInQuickShelf(&itemQuickShelf.QuickShelfID)
	if inv_error != nil {
		return inv_error
	}
	if len(*allItems) >= utils.MAX_AMOUNT_OF_ITEMS_IN_QUICK_SHELF {
		return inv_errors.INV_CONFLICT.WithDetails("Quick shelf is full")
	}

	// Check if item is already in quick shelf
	isInQuickShelf, inv_error := iqsc.ItemQuickShelfRepo.CheckIfItemAlreadyInQuickShelf(&itemQuickShelf.ItemID, &itemQuickShelf.QuickShelfID)
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
		quantity, inv_error := iqsc.ItemQuickShelfRepo.GetQuantityOfItemInQuickShelf(&itemQuickShelf.ItemID, &itemQuickShelf.QuickShelfID)
		if inv_error != nil {
			return inv_error
		}
		newQuantity := *quantity + itemQuickShelf.Quantity
		quickShelfItem.Quantity = &newQuantity

		// Update
		inv_error = iqsc.ItemQuickShelfRepo.UpdateQuantityOfItemInQuickShelf(tx, &quickShelfItem)
		if inv_error != nil {
			return inv_error
		}
	} else {
		// Insert
		inv_error = iqsc.ItemQuickShelfRepo.InsertNewItemInQuickShelf(tx, &quickShelfItem)
		if inv_error != nil {
			return inv_error
		}
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return nil
}

func (iqsc *ItemQuickShelfController) ClearQuickShelf(quickShelfId *uuid.UUID) *models.INVError {
	tx, err := iqsc.ItemQuickShelfRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Get all items in quick shelf
	quickShelfItems, inv_error := iqsc.ItemQuickShelfRepo.GetItemsInQuickShelf(quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	// Insert all items in regular shelf
	for _, itemInQuickShelf := range *quickShelfItems {
		itemId, err := uuid.Parse(itemInQuickShelf.Items.ID)
		if err != nil {
			return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error parsing item id from quick shelf items")
		}

		item, inv_error := iqsc.ItemRepo.GetItemById(&itemId)
		if inv_error != nil {
			return inv_error
		}

		// Get Quantity in regular shelf
		quantity, inv_error := iqsc.ItemsInShelfRepo.GetQuantityInShelve(&itemId)
		if inv_error != nil {
			return inv_error
		}
		newQuantity := *quantity + itemInQuickShelf.Quantity

		// Update Quantity in regular shelf
		inv_error = iqsc.ItemsInShelfRepo.UpdateItemInShelve(tx, &model.ItemsInShelf{
			ItemID:   item.ID,
			ShelfID:  *item.RegularShelfID,
			Quantity: &newQuantity,
		})
		if inv_error != nil {
			return inv_error
		}
	}
	// Clear quick shelf
	inv_error = iqsc.ItemQuickShelfRepo.ClearQuickShelf(tx, quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return nil
}

func (iqsc *ItemQuickShelfController) RemoveItemFromQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) *models.INVError {
	tx, err := iqsc.ItemQuickShelfRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Get Quantity in quick shelf
	quantity, inv_error := iqsc.ItemQuickShelfRepo.GetQuantityOfItemInQuickShelf(itemId, quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	// Get Item
	item, inv_error := iqsc.ItemRepo.GetItemById(itemId)
	if inv_error != nil {
		return inv_error
	}

	// Get Quantity in regular shelf
	quantityInRegularShelf, inv_error := iqsc.ItemsInShelfRepo.GetQuantityInShelve(itemId)
	if inv_error != nil {
		return inv_error
	}
	newQuantity := *quantityInRegularShelf + *quantity

	// Update Quantity in regular shelf
	inv_error = iqsc.ItemsInShelfRepo.UpdateItemInShelve(tx, &model.ItemsInShelf{
		ItemID:   itemId.String(),
		ShelfID:  *item.RegularShelfID,
		Quantity: &newQuantity,
	})
	if inv_error != nil {
		return inv_error
	}

	// Remove item from quick shelf
	inv_error = iqsc.ItemQuickShelfRepo.RemoveItemFromQuickShelf(tx, itemId, quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return nil
}
