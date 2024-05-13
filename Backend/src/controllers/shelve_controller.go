package controllers

import (
	"log"

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
	CreateShelve(shelve *models.OwnShelve) (*uuid.UUID, *models.INVError)
	UpdateShelve(shelve *models.OwnShelve) *models.INVError
	DeleteShelve(shelveId *uuid.UUID) *models.INVError
}

type ShelveController struct {
	ShelveRepo repositories.ShelveRepositoryI
	ShelveType repositories.ShelveTypeRepositoryI
}

func (sc *ShelveController) GetShelves() (*[]models.OwnShelve, *models.INVError) {
	shelves, inv_errors := sc.ShelveRepo.GetShelves()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return shelves, nil
}

func (sc *ShelveController) CreateShelve(shelve *models.OwnShelve) (*uuid.UUID, *models.INVError) {
	log.Print("Creating shelve")
	
	if shelve == nil {
		return nil, inv_errors.INV_BAD_REQUEST
	}

	log.Print("Shelve not null")
	log.Printf("ShelveType is %v ", &shelve.ShelveTypeName)

	shelveType, inv_error := sc.ShelveType.GetShelveTypeByName(&shelve.ShelveTypeName)
	if inv_error != nil {
		return nil, inv_error
	}

	log.Print("Type name ist %v and theirs id is %v", shelve.ShelveTypeName, shelveType.ID)

	var newShelve model.Shelves
	newShelve.ShelveTypeID = &shelveType.ID
	newShelve.RoomID = &shelve.RoomID

	shelveId, inv_error := sc.ShelveRepo.CreateShelve(&newShelve)
	if inv_error != nil {
		return nil, inv_error
	}
	return shelveId, nil
}

func (sc *ShelveController) UpdateShelve(shelve *models.OwnShelve) *models.INVError {
	if shelve == nil {
		return inv_errors.INV_BAD_REQUEST
	}

	shelveType, inv_error := sc.ShelveType.GetShelveTypeByName(&shelve.ShelveTypeName)
	if inv_error != nil {
		return inv_error
	}

	var newShelve model.Shelves
	newShelve.ShelveTypeID = &shelveType.ID
	newShelve.RoomID = &shelve.RoomID
	newShelve.ID = shelve.ID

	inv_errors := sc.ShelveRepo.UpdateShelve(&newShelve)
	if inv_errors != nil {
		return inv_errors
	}
	return nil
}

func (sc *ShelveController) DeleteShelve(shelveId *uuid.UUID) *models.INVError {
	// TODO Needs to be implemented
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
