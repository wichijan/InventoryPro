package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type RoleControllerI interface {
	GetRoles() (*[]model.Roles, *models.INVError)
	CreateRole(roleName *string) (*uuid.UUID, *models.INVError)
	UpdateRole(role *model.Roles) *models.INVError
}

type RoleController struct {
	RoleRepo repositories.RoleRepositoryI
}

func (rc *RoleController) GetRoles() (*[]model.Roles, *models.INVError) {
	return rc.RoleRepo.GetRoles()
}

func (rc *RoleController) CreateRole(roleName *string) (*uuid.UUID, *models.INVError) {
	tx, err := rc.RoleRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	roleId, inv_error := rc.RoleRepo.CreateRole(tx, roleName)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return roleId, nil
}

func (rc *RoleController) UpdateRole(role *model.Roles) *models.INVError {
	tx, err := rc.RoleRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	if *role.RoleName == "Admin" {
		return inv_errors.INV_CONFLICT
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return rc.RoleRepo.UpdateRole(tx, role)
}
