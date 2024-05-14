package controllers

import (
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type ItemControllerI interface {
	GetItems() (*[]models.ItemWithEverything, *models.INVError)
	GetItemById(itemId *uuid.UUID) (*models.ItemWithEverything, *models.INVError)
}

type ItemController struct {
	ItemRepo        repositories.ItemRepositoryI
	ItemKeywordRepo repositories.ItemKeywordRepositoryI
	ItemSubjectRepo repositories.ItemSubjectRepositoryI
	ItemPictureRepo repositories.ItemPictureRepositoryI
}

func (ic *ItemController) GetItems() (*[]models.ItemWithEverything, *models.INVError) {
	items, inv_errors := ic.ItemRepo.GetItems()
	if inv_errors != nil {
		return nil, inv_errors
	}

	return items, nil
}

func (ic *ItemController) GetItemById(itemId *uuid.UUID) (*models.ItemWithEverything, *models.INVError) {
	item, inv_errors := ic.ItemRepo.GetItemById(itemId)
	if inv_errors != nil {
		return nil, inv_errors
	}

	return item, nil
}
