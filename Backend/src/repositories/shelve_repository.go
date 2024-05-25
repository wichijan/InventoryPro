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
)

type ShelveRepositoryI interface {
	GetShelves() (*[]models.OwnShelve, *models.INVError)
	GetShelveById(id *uuid.UUID) (*models.OwnShelve, *models.INVError)
	GetShelvesWithItems() (*[]models.ShelveWithItems, *models.INVError)
	GetShelveByIdWithItems(id *uuid.UUID) (*models.ShelveWithItems, *models.INVError)
	CreateShelve(tx *sql.Tx, shelve *model.Shelves) (*uuid.UUID, *models.INVError)
	UpdateShelve(shelve *model.Shelves) *models.INVError
	DeleteShelve(shelveId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type ShelveRepository struct {
	managers.DatabaseManagerI
}

func (sr *ShelveRepository) GetShelves() (*[]models.OwnShelve, *models.INVError) {
	var shelves []models.OwnShelve

	// Create the query
	stmt := mysql.SELECT(
		table.Shelves.ID,
		table.ShelveTypes.TypeName,
		table.Shelves.RoomID,
	).FROM(
		table.Shelves.LEFT_JOIN(table.ShelveTypes, table.ShelveTypes.ID.EQ(table.Shelves.ShelveTypeID)),
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &shelves)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(shelves) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &shelves, nil
}

func (sr *ShelveRepository) GetShelveById(id *uuid.UUID) (*models.OwnShelve, *models.INVError) {
	var shelve models.OwnShelve

	// Create the query
	stmt := mysql.SELECT(
		table.Shelves.ID,
		table.ShelveTypes.TypeName,
		table.Shelves.RoomID,
	).FROM(
		table.Shelves.LEFT_JOIN(table.ShelveTypes, table.ShelveTypes.ID.EQ(table.Shelves.ShelveTypeID)),
	).WHERE(
		table.Shelves.ID.EQ(mysql.String(id.String())),
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &shelve)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &shelve, nil
}

func (sr *ShelveRepository) CreateShelve(tx *sql.Tx, shelve *model.Shelves) (*uuid.UUID, *models.INVError) {
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

func (sr *ShelveRepository) UpdateShelve(shelve *model.Shelves) *models.INVError {
	// Create the update statement
	updateQuery := table.Shelves.UPDATE(
		table.Shelves.ShelveTypeID,
		table.Shelves.RoomID,
	).SET(
		shelve.ShelveTypeID,
		shelve.RoomID,
	).WHERE(table.Shelves.ID.EQ(mysql.String(shelve.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(sr.GetDatabaseConnection())
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

func (sr *ShelveRepository) DeleteShelve(shelveId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}

func (sr *ShelveRepository) GetShelvesWithItems() (*[]models.ShelveWithItems, *models.INVError) {
	var shelvesWithItems []models.ShelveWithItems

	// Create the query
	stmt := mysql.SELECT(
		table.Shelves.ID,
		table.ShelveTypes.TypeName,
		table.Shelves.RoomID,
		table.Items.AllColumns,
	).FROM(
		table.Shelves.
			LEFT_JOIN(table.ShelveTypes, table.ShelveTypes.ID.EQ(table.Shelves.ShelveTypeID)).
			LEFT_JOIN(table.ItemsInShelve, table.ItemsInShelve.ShelveID.EQ(table.Shelves.ID)).
			LEFT_JOIN(table.Items, table.Items.ID.EQ(table.ItemsInShelve.ItemID)),
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &shelvesWithItems)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &shelvesWithItems, nil
}

func (sr *ShelveRepository) GetShelveByIdWithItems(id *uuid.UUID) (*models.ShelveWithItems, *models.INVError) {
	var shelveWithItems models.ShelveWithItems

	// Create the query
	stmt := mysql.SELECT(
		table.Shelves.ID,
		table.ShelveTypes.TypeName,
		table.Shelves.RoomID,
		table.Items.AllColumns,
	).FROM(
		table.Shelves.
			LEFT_JOIN(table.ShelveTypes, table.ShelveTypes.ID.EQ(table.Shelves.ShelveTypeID)).
			LEFT_JOIN(table.ItemsInShelve, table.ItemsInShelve.ShelveID.EQ(table.Shelves.ID)).
			LEFT_JOIN(table.Items, table.Items.ID.EQ(table.ItemsInShelve.ItemID)),
	).WHERE(
		table.Shelves.ID.EQ(mysql.String(id.String())),
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &shelveWithItems)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &shelveWithItems, nil
}
