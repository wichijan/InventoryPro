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
	return rc.RoleRepo.CreateRole(roleName)
}

func (rc *RoleController) UpdateRole(role *model.Roles) *models.INVError {
	if *role.RoleName == "Admin" {
		return inv_errors.INV_CONFLICT
	}

	return rc.RoleRepo.UpdateRole(role)
}
