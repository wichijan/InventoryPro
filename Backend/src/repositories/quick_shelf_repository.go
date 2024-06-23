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

type QuickShelfRepositoryI interface {
	GetQuickShelves() (*[]models.QuickShelfWithItems, *models.INVError)
	CreateQuickShelf(tx *sql.Tx, quickShelf *models.QuickShelfCreate) (*uuid.UUID, *models.INVError)
	UpdateQuickShelf(tx *sql.Tx, quickShelf *model.QuickShelves) *models.INVError
	DeleteQuickShelf(tx *sql.Tx, quickShelfId *uuid.UUID) *models.INVError

	CheckIfRoomIdExists(roomId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type QuickShelfRepository struct {
	managers.DatabaseManagerI
}

func (qsr *QuickShelfRepository) GetQuickShelves() (*[]models.QuickShelfWithItems, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.QuickShelves.AllColumns,
		table.Items.AllColumns,
	).FROM(
		table.QuickShelves.
			LEFT_JOIN(table.ItemQuickShelf, table.ItemQuickShelf.QuickShelfID.EQ(table.QuickShelves.QuickShelfID)).
			LEFT_JOIN(table.Items, table.Items.ID.EQ(table.ItemQuickShelf.ItemID)),
	)

	// Execute the query
	var quickShelves []models.QuickShelfWithItems
	err := stmt.Query(qsr.GetDatabaseConnection(), &quickShelves)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading quick shelves")
	}

	return &quickShelves, nil
}

func (qsr *QuickShelfRepository) CreateQuickShelf(tx *sql.Tx, quickShelf *models.QuickShelfCreate) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the query
	stmt := table.QuickShelves.INSERT(
		table.QuickShelves.QuickShelfID,
		table.QuickShelves.RoomID,
	).VALUES(
		uuid,
		quickShelf.RoomId,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating quick shelf")
	}

	return &uuid, nil
}

func (qsr *QuickShelfRepository) UpdateQuickShelf(tx *sql.Tx, quickShelf *model.QuickShelves) *models.INVError {
	// Create the query
	stmt := table.QuickShelves.UPDATE(
		table.QuickShelves.RoomID,
	).SET(
		quickShelf.RoomID,
	).WHERE(
		table.QuickShelves.QuickShelfID.EQ(mysql.String(quickShelf.QuickShelfID)),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating quick shelf")
	}

	return nil
}

func (qsr *QuickShelfRepository) DeleteQuickShelf(tx *sql.Tx, quickShelfId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.QuickShelves.DELETE().WHERE(
		table.QuickShelves.QuickShelfID.EQ(mysql.String(quickShelfId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting quick shelf")
	}

	return nil
}

func (qsr *QuickShelfRepository) CheckIfRoomIdExists(roomId *uuid.UUID) *models.INVError {
	count, err := utils.CountStatement(table.QuickShelves, table.QuickShelves.RoomID.EQ(mysql.String(roomId.String())), qsr.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if roomId exists in QuickShelves table")
	}
	if count < 0 {
		return inv_errors.INV_CONFLICT.WithDetails("QuickShelves still has roomId in it")
	}
	return nil
}
