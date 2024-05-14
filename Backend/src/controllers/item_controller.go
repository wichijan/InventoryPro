package controllers

import (
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type ItemControllerI interface {
	GetItems() (*[]models.ItemWithEverything, *models.INVError)
}

type ItemController struct {
	ItemRepo        repositories.ItemRepositoryI
	ItemKeywordRepo repositories.ItemKeywordRepositoryI
	ItemSubjectRepo repositories.ItemSubjectRepositoryI
	ItemPictureRepo repositories.ItemPictureRepositoryI
}

func (ic *ItemController) GetItems() (*[]models.ItemWithEverything, *models.INVError) {
	var items []models.ItemWithEverything

	pureItems, inv_errors := ic.ItemRepo.GetItems()
	if inv_errors != nil {
		return nil, inv_errors
	}

	for _, item := range *pureItems {
		itemWithEverything := models.ItemWithEverything{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			ClassOne:    item.ClassOne,
			ClassTwo:    item.ClassTwo,
			ClassThree:  item.ClassThree,
			ClassFour:   item.ClassFour,
			Damaged:     item.Damaged,
			DamagedDesc: item.DamagedDesc,
			Quantity:    item.Quantity,
			Status:      item.Status,
		}

		keywords, inv_errors := ic.ItemKeywordRepo.GetKeywordsForItem(&item.ID)
		if inv_errors != nil {
			return nil, inv_errors
		}
		itemWithEverything.Keywords = keywords

		subjects, inv_errors := ic.ItemSubjectRepo.GetSubjectsForItem(&item.ID)
		if inv_errors != nil {
			return nil, inv_errors
		}
		itemWithEverything.Subject = subjects

		pictures, inv_errors := ic.ItemPictureRepo.GetItemPictures(&item.ID)
		if inv_errors != nil {
			return nil, inv_errors
		}
		itemWithEverything.Pictures = pictures

		items = append(items, itemWithEverything)
	}

	return &items, nil
}
