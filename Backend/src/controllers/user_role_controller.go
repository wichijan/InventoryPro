package controllers

import (
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type UserRoleControllerI interface {
	GetRolesByUserId(userId *uuid.UUID) (*[]models.UserRoleWithName, *models.INVError)
}

type UserRoleController struct {
	UserRoleRepo repositories.UserRoleRepositoryI
}

func (urc *UserRoleController) GetRolesByUserId(userId *uuid.UUID) (*[]models.UserRoleWithName, *models.INVError) {
	return urc.UserRoleRepo.GetRolesByUserId(userId)
}
