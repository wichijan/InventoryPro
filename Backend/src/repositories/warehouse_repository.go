package repositories

import (
	"database/sql"

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
	GetWarehouses() (*[]model.Warehouses, *models.INVError)
	GetWarehouseById(id *uuid.UUID) (*model.Warehouses, *models.INVError)
	GetWarehousesWithRooms() (*[]models.WarehouseWithRooms, *models.INVError)
	GetWarehouseByIdWithRooms(id *uuid.UUID) (*models.WarehouseWithRooms, *models.INVError)
	CreateWarehouse(tx *sql.Tx, warehouse *models.WarehousesODT) (*uuid.UUID, *models.INVError)
	UpdateWarehouse(tx *sql.Tx, Warehouse *model.Warehouses) *models.INVError
	DeleteWarehouse(tx *sql.Tx, warehouseId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type WarehouseRepository struct {
	managers.DatabaseManagerI
}

func (wr *WarehouseRepository) GetWarehouses() (*[]model.Warehouses, *models.INVError) {
	var warehouses []model.Warehouses

	// Create the query
	stmt := mysql.SELECT(
		table.Warehouses.AllColumns,
	).FROM(
		table.Warehouses,
	)

	// Execute the query
	err := stmt.Query(wr.GetDatabaseConnection(), &warehouses)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(warehouses) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &warehouses, nil
}

func (wr *WarehouseRepository) GetWarehouseById(id *uuid.UUID) (*model.Warehouses, *models.INVError) {
	var warehouse model.Warehouses

	// Create the query
	stmt := mysql.SELECT(
		table.Warehouses.AllColumns,
	).FROM(
		table.Warehouses,
	).WHERE(
		table.Warehouses.ID.EQ(utils.MySqlString(id.String())),
	)

	// Execute the query
	err := stmt.Query(wr.GetDatabaseConnection(), &warehouse)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &warehouse, nil
}

func (wr *WarehouseRepository) CreateWarehouse(tx *sql.Tx, warehouse *models.WarehousesODT) (*uuid.UUID, *models.INVError) {
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
	rows, err := insertQuery.Exec(tx)
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

func (wr *WarehouseRepository) UpdateWarehouse(tx *sql.Tx, warehouse *model.Warehouses) *models.INVError {

	// Create the update statement
	updateQuery := table.Warehouses.UPDATE(
		table.Warehouses.Name,
		table.Warehouses.Description,
	).SET(
		warehouse.Name,
		warehouse.Description,
	).WHERE(table.Warehouses.ID.EQ(mysql.String(warehouse.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(tx)
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

func (wr *WarehouseRepository) DeleteWarehouse(tx *sql.Tx, warehouseId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}

func (wr *WarehouseRepository) GetWarehousesWithRooms() (*[]models.WarehouseWithRooms, *models.INVError) {
	var warehouse []models.WarehouseWithRooms

	// Create the query
	// Create the query
	stmt := mysql.SELECT(
		table.Warehouses.AllColumns,
		table.Rooms.AllColumns,
	).FROM(
		table.Warehouses.
			LEFT_JOIN(table.Rooms, table.Rooms.WarehouseID.EQ(table.Warehouses.ID)),
	)

	// Execute the query
	err := stmt.Query(wr.GetDatabaseConnection(), &warehouse)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &warehouse, nil
}

func (wr *WarehouseRepository) GetWarehouseByIdWithRooms(id *uuid.UUID) (*models.WarehouseWithRooms, *models.INVError) {
	var warehouse models.WarehouseWithRooms

	// Create the query
	stmt := mysql.SELECT(
		table.Warehouses.AllColumns,
		table.Rooms.AllColumns,
	).FROM(
		table.Warehouses.
			LEFT_JOIN(table.Rooms, table.Rooms.WarehouseID.EQ(table.Warehouses.ID)),
	).WHERE(
		table.Warehouses.ID.EQ(utils.MySqlString(id.String())),
	)

	// Execute the query
	err := stmt.Query(wr.GetDatabaseConnection(), &warehouse)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &warehouse, nil
}
