package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type RoororepositoryI interface {
	GetShelves() (*[]model.Shelves, *models.INVError)
	CreateShelve(shelve *model.Shelves) (*uuid.UUID, *models.INVError)
	UpdateShelve(shelve *model.Shelves) *models.INVError
	DeleteShelve(shelveId *uuid.UUID) *models.INVError
}

type Roororepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (ror *Roororepository) GetShelves() (*[]model.Shelves, *models.INVError) {
	var shelves []model.Shelves

	// Create the query 
	stmt := mysql.SELECT(
		table.Shelves.AllColumns,
	).FROM(
		table.Shelves,
	)

	// Execute the query
	err := stmt.Query(ror.DatabaseManager.GetDatabaseConnection(), &shelves)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(shelves) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &shelves, nil
}

func (ror *Roororepository) CreateShelve(shelve *model.Shelves) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Shelves.INSERT(
		table.Shelves.ID,
		table.Shelves.ShelveTypeID,
		table.Shelves.RoomID,
	).VALUES(
		uuid.String(),
		shelve.ShelveTypeID,
		shelve.RoomID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(ror.DatabaseManager.GetDatabaseConnection())
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

func (ror *Roororepository) UpdateShelve(shelve *model.Shelves) *models.INVError {
	// Create the update statement
	updateQuery := table.Shelves.UPDATE(
		table.Shelves.ShelveTypeID,
		table.Shelves.RoomID,
	).SET(
		shelve.ShelveTypeID,
		shelve.RoomID,
	).WHERE(table.Shelves.ID.EQ(mysql.String(shelve.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(ror.DatabaseManager.GetDatabaseConnection())
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

func (ror *Roororepository) DeleteShelve(shelveId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
