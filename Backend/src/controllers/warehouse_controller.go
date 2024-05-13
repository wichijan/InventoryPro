package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type WarehouseControllerI interface {
	GetWarehouses() (*[]model.Warehouses, *models.INVError)
	GetWarehouseById(id *uuid.UUID) (*model.Warehouses, *models.INVError)
	GetWarehousesWithRooms() (*[]models.WarehouseWithRooms, *models.INVError)
	GetWarehouseByIdWithRooms(id *uuid.UUID) (*models.WarehouseWithRooms, *models.INVError)
	CreateWarehouse(warehouse *model.Warehouses) (*uuid.UUID, *models.INVError)
	UpdateWarehouse(warehouse *model.Warehouses) *models.INVError
	DeleteWarehouse(warehouse_id *uuid.UUID) *models.INVError
}

type WarehouseController struct {
	WarehouseRepo repositories.WarehouseRepositoryI
}

func (mc *WarehouseController) GetWarehouses() (*[]model.Warehouses, *models.INVError) {
	warehouses, inv_errors := mc.WarehouseRepo.GetWarehouses()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return warehouses, nil
}

func (mc *WarehouseController) CreateWarehouse(warehouse *model.Warehouses) (*uuid.UUID, *models.INVError) {
	if warehouse == nil {
		return nil, inv_errors.INV_BAD_REQUEST
	}

	warehouseId, inv_errors := mc.WarehouseRepo.CreateWarehouse(warehouse)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return warehouseId, nil
}

func (mc *WarehouseController) UpdateWarehouse(warehouse *model.Warehouses) *models.INVError {
	inv_errors := mc.WarehouseRepo.UpdateWarehouse(warehouse)
	if inv_errors != nil {
		return inv_errors
	}
	return nil
}

func (mc *WarehouseController) DeleteWarehouse(warehouse_id *uuid.UUID) *models.INVError {
	// TODO Needs to be implemented
	return nil
}

func (mc *WarehouseController) GetWarehouseById(id *uuid.UUID) (*model.Warehouses, *models.INVError) {
	warehouse, inv_errors := mc.WarehouseRepo.GetWarehouseById(id)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return warehouse, nil
}

func (mc *WarehouseController) GetWarehousesWithRooms() (*[]models.WarehouseWithRooms, *models.INVError) {
	warehouse, inv_errors := mc.WarehouseRepo.GetWarehousesWithRooms()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return warehouse, nil
}

func (mc *WarehouseController) GetWarehouseByIdWithRooms(id *uuid.UUID) (*models.WarehouseWithRooms, *models.INVError) {
	warehouse, inv_errors := mc.WarehouseRepo.GetWarehouseByIdWithRooms(id)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return warehouse, nil
}
