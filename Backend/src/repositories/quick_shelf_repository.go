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

type QuickShelfRepositoryI interface {
	GetQuickShelves() (*model.QuickShelves, *models.INVError)
	CreateQuickShelf(tx *sql.Tx, book *model.QuickShelves) (*string, *models.INVError)
	DeleteQuickShelf(tx *sql.Tx, shelfId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type QuickShelfRepository struct {
	managers.DatabaseManagerI
}

func (qsr *QuickShelfRepository) GetQuickShelves() (*model.QuickShelves, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.QuickShelves.AllColumns,
	).FROM(
		table.QuickShelves,
	)

	// Execute the query
	var shelf model.QuickShelves
	err := stmt.Query(qsr.GetDatabaseConnection(), &shelf)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading quick shelves")
	}

	return &shelf, nil
}

func (qsr *QuickShelfRepository) CreateQuickShelf(tx *sql.Tx, shelf *model.QuickShelves) (*string, *models.INVError) {
	// Create the query
	stmt := table.QuickShelves.INSERT(
		table.QuickShelves.QuickShelfID,
		table.QuickShelves.RoomID,
	).VALUES(
		shelf.QuickShelfID,
		shelf.RoomID,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating quick shelf")
	}

	return &shelf.QuickShelfID, nil
}

func (qsr *QuickShelfRepository) DeleteQuickShelf(tx *sql.Tx, shelfId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.QuickShelves.DELETE().WHERE(
		table.QuickShelves.QuickShelfID.EQ(mysql.String(shelfId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting quick shelf")
	}

	return nil
}
