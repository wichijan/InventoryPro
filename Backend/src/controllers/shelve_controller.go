package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type ShelveControllerI interface {
	GetShelves() (*[]models.OwnShelve, *models.INVError)
	GetShelveById(id *uuid.UUID) (*models.OwnShelve, *models.INVError)
	GetShelvesWithItems() (*[]models.ShelveWithItems, *models.INVError)
	GetShelveByIdWithItems(id *uuid.UUID) (*models.ShelveWithItems, *models.INVError)
	CreateShelve(shelve *models.ShelveOTD) (*uuid.UUID, *models.INVError)
	UpdateShelve(shelve *models.OwnShelve) *models.INVError
	DeleteShelve(shelveId *uuid.UUID) *models.INVError
}

type ShelveController struct {
	ShelveRepo     repositories.ShelveRepositoryI
	ShelveTypeRepo repositories.ShelveTypeRepositoryI
}

func (sc *ShelveController) GetShelves() (*[]models.OwnShelve, *models.INVError) {
	shelves, inv_errors := sc.ShelveRepo.GetShelves()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return shelves, nil
}

func (sc *ShelveController) CreateShelve(shelve *models.ShelveOTD) (*uuid.UUID, *models.INVError) {
	tx, err := sc.ShelveRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	if shelve == nil {
		return nil, inv_errors.INV_BAD_REQUEST
	}

	var newShelve model.Shelves
	newShelve.RoomID = &shelve.RoomID

	shelveTypeObj, inv_error := sc.ShelveTypeRepo.GetShelveTypeByName(&shelve.ShelveTypeName)
	if inv_error == inv_errors.INV_NOT_FOUND {
		shelveTypeId, inv_error_Create := sc.ShelveTypeRepo.CreateShelveType(tx, &shelve.ShelveTypeName)
		if inv_error_Create != nil {
			return nil, inv_error
		}
		shelveTypeIdString := shelveTypeId.String()
		newShelve.ShelveTypeID = &shelveTypeIdString
	} else if inv_error != nil {
		return nil, inv_error
	} else {
		newShelve.ShelveTypeID = &shelveTypeObj.ID
	}

	shelveId, inv_error := sc.ShelveRepo.CreateShelve(tx, &newShelve)
	if inv_error != nil {
		return nil, inv_error
	}
	return shelveId, nil
}

func (sc *ShelveController) UpdateShelve(shelve *models.OwnShelve) *models.INVError {
	tx, err := sc.ShelveRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	if shelve == nil {
		return inv_errors.INV_BAD_REQUEST
	}

	shelveType, inv_error := sc.ShelveTypeRepo.GetShelveTypeByName(&shelve.ShelveTypeName)
	if inv_error != nil {
		return inv_error
	}

	var newShelve model.Shelves
	newShelve.ID = shelve.ID
	newShelve.RoomID = &shelve.RoomID
	newShelve.ShelveTypeID = &shelveType.ID

	inv_errors := sc.ShelveRepo.UpdateShelve(tx, &newShelve)
	if inv_errors != nil {
		return inv_errors
	}

	return nil
}

func (sc *ShelveController) DeleteShelve(shelveId *uuid.UUID) *models.INVError {
	tx, err := sc.ShelveRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	return nil
}

func (sc *ShelveController) GetShelveById(id *uuid.UUID) (*models.OwnShelve, *models.INVError) {
	shelve, inv_errors := sc.ShelveRepo.GetShelveById(id)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return shelve, nil
}

func (sc *ShelveController) GetShelvesWithItems() (*[]models.ShelveWithItems, *models.INVError) {
	shelvesWithItems, inv_errors := sc.ShelveRepo.GetShelvesWithItems()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return shelvesWithItems, nil
}

func (sc *ShelveController) GetShelveByIdWithItems(id *uuid.UUID) (*models.ShelveWithItems, *models.INVError) {
	shelveWithItems, inv_errors := sc.ShelveRepo.GetShelveByIdWithItems(id)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return shelveWithItems, nil
}
