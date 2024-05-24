package controllers

import (
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type UserRoleControllerI interface {
	CreateUserRole(user_role *model.UserRoles) *models.INVError
	UpdateUserRole(userRole *model.UserRoles) *models.INVError
	DeleteUserRole(userRoleId *uuid.UUID) *models.INVError
}

type UserRoleController struct {
	UserRoleRepo repositories.UserRoleRepositoryI
}

// TODO Implement the CreateUserRole, UpdateUserRole, and DeleteUserRole functions correctly
func (urc *UserRoleController) CreateUserRole(user_role *model.UserRoles) *models.INVError {
	return urc.UserRoleRepo.CreateUserRole(user_role)
}

func (urc *UserRoleController) UpdateUserRole(userRole *model.UserRoles) *models.INVError {
	return urc.UserRoleRepo.UpdateUserRole(userRole)
}

func (urc *UserRoleController) DeleteUserRole(userRoleId *uuid.UUID) *models.INVError {
	return urc.UserRoleRepo.DeleteUserRole(userRoleId)
}
