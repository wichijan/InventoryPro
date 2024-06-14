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

type ShelveRepositoryI interface {
	GetShelves() (*[]model.Shelves, *models.INVError)
	GetShelveById(id *uuid.UUID) (*model.Shelves, *models.INVError)
	GetShelvesWithItems() (*[]models.ShelveWithItems, *models.INVError)
	GetShelveByIdWithItems(id *uuid.UUID) (*models.ShelveWithItems, *models.INVError)
	CreateShelve(tx *sql.Tx, shelve *model.Shelves) (*uuid.UUID, *models.INVError)
	UpdateShelve(tx *sql.Tx, shelve *model.Shelves) *models.INVError
	DeleteShelve(tx *sql.Tx, shelveId *uuid.UUID) *models.INVError
	CheckIfShelveExists(shelveId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type ShelveRepository struct {
	managers.DatabaseManagerI
}

func (sr *ShelveRepository) GetShelves() (*[]model.Shelves, *models.INVError) {
	var shelves []model.Shelves

	// Create the query
	stmt := mysql.SELECT(
		table.Shelves.ID,
		table.Shelves.RoomID,
	).FROM(
		table.Shelves,
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &shelves)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading shelves")
	}

	return &shelves, nil
}

func (sr *ShelveRepository) GetShelveById(id *uuid.UUID) (*model.Shelves, *models.INVError) {
	var shelve model.Shelves

	// Create the query
	stmt := mysql.SELECT(
		table.Shelves.ID,
		table.Shelves.RoomID,
	).FROM(
		table.Shelves,
	).WHERE(
		table.Shelves.ID.EQ(mysql.String(id.String())),
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &shelve)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading shelve")
	}

	return &shelve, nil
}

func (sr *ShelveRepository) CreateShelve(tx *sql.Tx, shelve *model.Shelves) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Shelves.INSERT(
		table.Shelves.ID,
		table.Shelves.RoomID,
	).VALUES(
		uuid.String(),
		shelve.RoomID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating shelve")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating shelve")
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Shelve already exists")
	}

	return &uuid, nil
}

func (sr *ShelveRepository) UpdateShelve(tx *sql.Tx, shelve *model.Shelves) *models.INVError {
	// Create the update statement
	updateQuery := table.Shelves.UPDATE(
		table.Shelves.RoomID,
	).SET(
		shelve.RoomID,
	).WHERE(table.Shelves.ID.EQ(mysql.String(shelve.ID)))

	// Execute the query
	_, err := updateQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating shelve")
	}

	return nil
}

func (sr *ShelveRepository) DeleteShelve(tx *sql.Tx, shelveId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.Shelves.DELETE().WHERE(
		table.Shelves.ID.EQ(mysql.String(shelveId.String())),
	)

	// Execute the query
	rows, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting shelve")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting shelve")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Shelve does not exist")
	}

	return nil
}

func (sr *ShelveRepository) GetShelvesWithItems() (*[]models.ShelveWithItems, *models.INVError) {
	var shelvesWithItems []models.ShelveWithItems

	// Create the query
	stmt := mysql.SELECT(
		table.Shelves.ID,
		table.Shelves.RoomID,
		table.Items.AllColumns,
		table.ItemsInShelf.Quantity,
	).FROM(
		table.Shelves.
			LEFT_JOIN(table.ItemsInShelf, table.ItemsInShelf.ShelfID.EQ(table.Shelves.ID)).
			LEFT_JOIN(table.Items, table.Items.ID.EQ(table.ItemsInShelf.ItemID)),
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &shelvesWithItems)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading shelves with items")
	}

	return &shelvesWithItems, nil
}

func (sr *ShelveRepository) GetShelveByIdWithItems(id *uuid.UUID) (*models.ShelveWithItems, *models.INVError) {
	var shelveWithItems models.ShelveWithItems

	// Create the query
	stmt := mysql.SELECT(
		table.Shelves.ID,
		table.Shelves.RoomID,
		table.Items.AllColumns,
	).FROM(
		table.Shelves.
			LEFT_JOIN(table.ItemsInShelf, table.ItemsInShelf.ShelfID.EQ(table.Shelves.ID)).
			LEFT_JOIN(table.Items, table.Items.ID.EQ(table.ItemsInShelf.ItemID)),
	).WHERE(
		table.Shelves.ID.EQ(mysql.String(id.String())),
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &shelveWithItems)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading shelve with items")
	}

	return &shelveWithItems, nil
}

func (sr *ShelveRepository) CheckIfShelveExists(shelveId *uuid.UUID) *models.INVError {
	count, err := utils.CountStatement(table.Shelves, table.Shelves.ID.EQ(mysql.String(shelveId.String())), sr.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if shelve exists")
	}
	if count <= 0 {
		return inv_errors.INV_CONFLICT.WithDetails("ShelveId does not exist")
	}
	return nil
}
