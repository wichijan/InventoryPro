package controllers

import (
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
	RemoveKeywordFromItem(itemKeyword models.ItemWithKeywordName) *models.INVError
	AddSubjectToItem(itemSubject models.ItemWithSubjectName) *models.INVError
	RemoveSubjectFromItem(itemSubject models.ItemWithSubjectName) *models.INVError
}

type ItemController struct {
	ItemRepo        repositories.ItemRepositoryI
	KeywordRepo     repositories.KeywordRepositoryI
	SubjectRepo     repositories.SubjectRepositoryI
	ItemKeywordRepo repositories.ItemKeywordRepositoryI
	ItemSubjectRepo repositories.ItemSubjectRepositoryI
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

	inv_err := ic.ItemKeywordRepo.CheckIfKeywordAndItemExists(itemKeywordWithID)
	if inv_err != nil {
		return inv_err
	}

	inv_error = ic.ItemKeywordRepo.CreateKeywordForItem(&itemKeywordWithID)
	if inv_error != nil {
		return inv_error
	}

	return nil
}

func (ic *ItemController) RemoveKeywordFromItem(itemKeyword models.ItemWithKeywordName) *models.INVError {
	keyword, inv_error := ic.KeywordRepo.GetKeywordByName(&itemKeyword.KeywordName)
	if inv_error != nil {
		return inv_error
	}

	itemKeywordWithID := models.ItemWithKeyword{
		ItemID:    itemKeyword.ItemID,
		KeywordID: keyword.ID,
	}

	inv_error = ic.ItemKeywordRepo.DeleteKeywordForItem(&itemKeywordWithID)
	if inv_error != nil {
		return inv_error
	}

	return nil
}

func (ic *ItemController) AddSubjectToItem(itemSubject models.ItemWithSubjectName) *models.INVError {
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

	inv_error = ic.ItemSubjectRepo.CreateSubjectForItem(&itemSubjectWithID)
	if inv_error != nil {
		return inv_error
	}

	return nil
}

func (ic *ItemController) RemoveSubjectFromItem(itemSubject models.ItemWithSubjectName) *models.INVError {
	subject, inv_error := ic.SubjectRepo.GetSubjectByName(&itemSubject.SubjectName)
	if inv_error != nil {
		return inv_error
	}

	itemSubjectWithID := models.ItemWithSubject{
		ItemID:    itemSubject.ItemID,
		SubjectID: subject.ID,
	}

	inv_error = ic.ItemSubjectRepo.DeleteSubjectForItem(&itemSubjectWithID)
	if inv_error != nil {
		return inv_error
	}

	return nil
}
