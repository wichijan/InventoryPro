package controllers

import (
	"log"

	"github.com/google/uuid"
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
	RemoveKeywordFromItem(itemKeyword models.ItemWithKeyword) *models.INVError
}

type ItemController struct {
	ItemRepo        repositories.ItemRepositoryI
	ItemStatusRepo  repositories.ItemStatusRepositoryI
	KeywordRepo     repositories.KeywordRepositoryI
	ItemKeywordRepo repositories.ItemKeywordRepositoryI
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
	pureItem.Quantity = &item.Quantity

	if item.Status != "" {
		statusId, inv_error := ic.ItemStatusRepo.GetStatusIdByName(&item.Status)
		if inv_error != nil {
			return nil, inv_error
		}

		log.Print(statusId.String())
		statusString := statusId.String()
		pureItem.StatusID = &statusString
	} else {
		pureItem.StatusID = nil
	}

	id, inv_error := ic.ItemRepo.CreateItem(&pureItem)
	if inv_error != nil {
		return nil, inv_error
	}

	return id, nil
}

func (ic *ItemController) UpdateItem(item *models.ItemWithStatus) *models.INVError {
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
	pureItem.Quantity = &item.Quantity

	if item.Status != "" {
		statusId, inv_error := ic.ItemStatusRepo.GetStatusIdByName(&item.Status)
		if inv_error != nil {
			return inv_error
		}

		log.Print(statusId.String())
		statusString := statusId.String()
		pureItem.StatusID = &statusString
	} else {
		pureItem.StatusID = nil
	}

	inv_error := ic.ItemRepo.UpdateItem(&pureItem)
	if inv_error != nil {
		return inv_error
	}

	return nil
}

func (ic *ItemController) AddKeywordToItem(itemKeyword models.ItemWithKeywordName) *models.INVError {
	keyword, inv_error := ic.KeywordRepo.GetKeywordByName(&itemKeyword.KeywordName)
	if inv_error != nil {
		return inv_error
	}

	itemKeywordWithID := models.ItemWithKeyword{
		ItemID:    itemKeyword.ItemID,
		KeywordID: keyword.ID,
	}

	// TODO Move to handler
	inv_err := ic.ItemKeywordRepo.CheckIfKeywordAndItemExists(itemKeywordWithID)
	if inv_err != nil {
		return inv_err
	}

	_, inv_error = ic.ItemKeywordRepo.CreateKeywordForItem(&itemKeywordWithID)
	if inv_error != nil {
		return inv_error
	}

	return nil
}

func (ic *ItemController) RemoveKeywordFromItem(itemKeyword models.ItemWithKeyword) *models.INVError {
	keyword, inv_error := ic.KeywordRepo.GetKeywordByName(&itemKeyword.KeywordID)
	if inv_error != nil {
		return inv_error
	}

	// TODO Move to handler
	inv_err := ic.KeywordRepo.CheckIfKeywordExists(&keyword.ID)
	if inv_err != nil {
		return inv_err
	}

	inv_error = ic.ItemKeywordRepo.DeleteKeywordForItem(&itemKeyword)
	if inv_error != nil {
		return inv_error
	}

	return nil
}
