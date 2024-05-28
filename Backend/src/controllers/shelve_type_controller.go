package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type ShelveTypeControllerI interface {
	GetShelveTypes() (*[]model.ShelveTypes, *models.INVError)
	CreateShelveType(shelveTypeName *string) (*uuid.UUID, *models.INVError)
	UpdateShelveType(shelveType *model.ShelveTypes) *models.INVError
	DeleteShelveType(shelveTypeId *uuid.UUID) *models.INVError
}

type ShelveTypeController struct {
	ShelveTypeRepo repositories.ShelveTypeRepositoryI
}

func (stc *ShelveTypeController) GetShelveTypes() (*[]model.ShelveTypes, *models.INVError) {
	shelveTypes, inv_error := stc.ShelveTypeRepo.GetShelveTypes()
	if inv_error != nil {
		return nil, inv_error
	}
	return shelveTypes, nil
}

func (stc *ShelveTypeController) CreateShelveType(shelveTypeName *string) (*uuid.UUID, *models.INVError) {
	tx, err := stc.ShelveTypeRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	shelveTypeId, inv_error := stc.ShelveTypeRepo.CreateShelveType(tx, shelveTypeName)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return shelveTypeId, nil
}

func (stc *ShelveTypeController) UpdateShelveType(shelveType *model.ShelveTypes) *models.INVError {
	tx, err := stc.ShelveTypeRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()
	
	inv_error := stc.ShelveTypeRepo.UpdateShelveType(tx, shelveType)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (stc *ShelveTypeController) DeleteShelveType(shelveTypeId *uuid.UUID) *models.INVError {
	tx, err := stc.ShelveTypeRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	inv_error := stc.ShelveTypeRepo.DeleteShelveType(tx, shelveTypeId)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	
	return nil
}
