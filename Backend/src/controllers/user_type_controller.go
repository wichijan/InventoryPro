package controllers

import (
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type UserTypeControllerI interface {
	GetUserTypes() (*[]model.UserTypes, *models.INVError)
}

type UserTypeController struct {
	UserTypeRepo repositories.UserTypeRepositoryI
}

func (utc *UserTypeController) GetUserTypes() (*[]model.UserTypes, *models.INVError) {
	return utc.UserTypeRepo.GetUserTypes()
}
