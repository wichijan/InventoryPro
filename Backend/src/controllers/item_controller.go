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
}

type ItemController struct {
	ItemRepo        repositories.ItemRepositoryI
	ItemKeywordRepo repositories.ItemKeywordRepositoryI
	ItemSubjectRepo repositories.ItemSubjectRepositoryI
	ItemPictureRepo repositories.ItemPictureRepositoryI
	ItemStatus      repositories.ItemStatusRepositoryI
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
	log.Print(item)

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

	log.Print(item.Status)

	statusId, inv_error := ic.ItemStatus.GetStatusIdByName(&item.Status)
	if inv_error != nil {
		return nil, inv_error
	}

	statusString := statusId.String()
	pureItem.StatusID = &statusString

	id, inv_error := ic.ItemRepo.CreateItem(&pureItem)
	if inv_error != nil {
		return nil, inv_error
	}

	return id, nil
}
