package controllers

import (
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type UserRoleControllerI interface {
	CreateUserRole(user_role *model.UserRoles) *models.INVError
	DeleteUserRole(userRole *model.UserRoles) *models.INVError
}

type UserRoleController struct {
	UserRoleRepo repositories.UserRoleRepositoryI
}

func (urc *UserRoleController) CreateUserRole(user_role *model.UserRoles) *models.INVError {
	tx, err := urc.UserRoleRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	inv_error := urc.UserRoleRepo.CreateUserRole(tx, user_role)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (urc *UserRoleController) DeleteUserRole(userRole *model.UserRoles) *models.INVError {
	tx, err := urc.UserRoleRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	inv_error := urc.UserRoleRepo.DeleteUserRole(tx, userRole)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}
