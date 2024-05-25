package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type UserTypeControllerI interface {
	GetUserTypes() (*[]model.UserTypes, *models.INVError)
	CreateUserType(type_name *string) (*uuid.UUID, *models.INVError)
	UpdateUserType(userType *model.UserTypes) *models.INVError
	DeleteUserType(userTypeId *uuid.UUID) *models.INVError
}

type UserTypeController struct {
	UserTypeRepo repositories.UserTypeRepositoryI
}

func (utc *UserTypeController) GetUserTypes() (*[]model.UserTypes, *models.INVError) {
	return utc.UserTypeRepo.GetUserTypes()
}

func (utc *UserTypeController) CreateUserType(type_name *string) (*uuid.UUID, *models.INVError) {
	tx, err := utc.UserTypeRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	return utc.UserTypeRepo.CreateUserType(tx, type_name)
}

func (utc *UserTypeController) UpdateUserType(userType *model.UserTypes) *models.INVError {
	tx, err := utc.UserTypeRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	return utc.UserTypeRepo.UpdateUserType(tx, userType)
}

func (utc *UserTypeController) DeleteUserType(userTypeId *uuid.UUID) *models.INVError {
	tx, err := utc.UserTypeRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	return utc.UserTypeRepo.DeleteUserType(tx, userTypeId)
}
