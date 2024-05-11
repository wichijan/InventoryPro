package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

type WarehouseRepositoryI interface {
	// Warehouse
	GetWarehouses() (*[]model.Warehouses, *models.INVError)
	GetWarehouseByName(name *string) (*model.Warehouses, *models.INVError)
	CreateWarehouse(warehouse *model.Warehouses) (*uuid.UUID, *models.INVError)
	UpdateWarehouse(Warehouse *model.Warehouses) *models.INVError
	DeleteWarehouse(warehouseId *uuid.UUID) *models.INVError
}

type WarehouseRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

// Warehouse
func (mr *WarehouseRepository) GetWarehouses() (*[]model.Warehouses, *models.INVError) {
	var warehouses []model.Warehouses

	// Create the query
	stmt := mysql.SELECT(
		table.Warehouses.AllColumns,
	).FROM(
		table.Warehouses,
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &warehouses)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(warehouses) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &warehouses, nil
}

func (mr *WarehouseRepository) GetWarehouseByName(name *string) (*model.Warehouses, *models.INVError) {
	var warehouse model.Warehouses

	// Create the query
	stmt := mysql.SELECT(
		table.Warehouses.AllColumns,
	).FROM(
		table.Warehouses,
	).WHERE(
		table.Warehouses.Name.EQ(utils.MySqlString(*name)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &warehouse)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &warehouse, nil
}

func (mr *WarehouseRepository) CreateWarehouse(warehouse *model.Warehouses) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Warehouses.INSERT(
		table.Warehouses.ID,
		table.Warehouses.Name,
		table.Warehouses.Description,
	).
		VALUES(
			uuid.String(),
			warehouse.Name,
			warehouse.Description,
		)

	// Execute the query
	rows, err := insertQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &uuid, nil
}

func (mr *WarehouseRepository) UpdateWarehouse(warehouse *model.Warehouses) *models.INVError {

	// Create the update statement
	updateQuery := table.Warehouses.UPDATE(
		table.Warehouses.Name,
		table.Warehouses.Description,
	).SET(
		warehouse.Name,
		warehouse.Description,
	).WHERE(table.Warehouses.ID.EQ(mysql.String(warehouse.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND
	}

	return nil
}

func (mr *WarehouseRepository) DeleteWarehouse(warehouseId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
