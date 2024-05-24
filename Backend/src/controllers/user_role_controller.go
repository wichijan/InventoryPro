package controllers

import (
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
	return urc.UserRoleRepo.CreateUserRole(user_role)
}

func (urc *UserRoleController) DeleteUserRole(userRole *model.UserRoles) *models.INVError {
	return urc.UserRoleRepo.DeleteUserRole(userRole)
}
