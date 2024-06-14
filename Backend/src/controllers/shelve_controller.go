package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type ShelveControllerI interface {
	GetShelves() (*[]model.Shelves, *models.INVError)
	GetShelveById(id *uuid.UUID) (*model.Shelves, *models.INVError)
	GetShelvesWithItems() (*[]models.ShelveWithItems, *models.INVError)
	GetShelveByIdWithItems(id *uuid.UUID) (*models.ShelveWithItems, *models.INVError)
	CreateShelve(shelve *models.ShelveOTD) (*uuid.UUID, *models.INVError)
	UpdateShelve(shelve *model.Shelves) *models.INVError
	DeleteShelve(shelveId *uuid.UUID) *models.INVError
}

type ShelveController struct {
	ShelveRepo   repositories.ShelveRepositoryI
	ItemRepo     repositories.ItemRepositoryI
	ItemInShelve repositories.ItemInShelveRepositoryI
}

func (sc *ShelveController) GetShelves() (*[]model.Shelves, *models.INVError) {
	shelves, inv_errors := sc.ShelveRepo.GetShelves()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return shelves, nil
}

func (sc *ShelveController) CreateShelve(shelve *models.ShelveOTD) (*uuid.UUID, *models.INVError) {
	tx, err := sc.ShelveRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	if shelve == nil {
		return nil, inv_errors.INV_BAD_REQUEST.WithDetails("invalid shelve data")
	}

	var newShelve model.Shelves
	newShelve.RoomID = &shelve.RoomID

	shelveId, inv_error := sc.ShelveRepo.CreateShelve(tx, &newShelve)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return shelveId, nil
}

func (sc *ShelveController) UpdateShelve(shelve *model.Shelves) *models.INVError {
	tx, err := sc.ShelveRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	if shelve == nil {
		return inv_errors.INV_BAD_REQUEST.WithDetails("invalid shelve data")
	}

	inv_error := sc.ShelveRepo.UpdateShelve(tx, shelve)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (sc *ShelveController) DeleteShelve(shelveId *uuid.UUID) *models.INVError {
	tx, err := sc.ShelveRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	// Check if shelve exists in items
	if err := sc.ItemRepo.CheckIfShelfIdExists(shelveId); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting shelve")
	}

	// Check if shelve exists in items_in_shelf
	if err := sc.ItemInShelve.CheckIfShelfIdExists(shelveId); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting shelve")
	}

	// Delete
	inv_error := sc.ShelveRepo.DeleteShelve(tx, shelveId)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (sc *ShelveController) GetShelveById(id *uuid.UUID) (*model.Shelves, *models.INVError) {
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
