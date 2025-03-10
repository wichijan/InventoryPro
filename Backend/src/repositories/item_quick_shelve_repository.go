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

type ItemQuickShelfRepositoryI interface {
	DeleteItemsQuickShelf(tx *sql.Tx, itemId *uuid.UUID) *models.INVError

	GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]models.GetQuickShelf, *models.INVError)
	InsertNewItemInQuickShelf(tx *sql.Tx, itemQuickShelf *model.ItemQuickShelf) *models.INVError
	UpdateQuantityOfItemInQuickShelf(tx *sql.Tx, itemQuickShelf *model.ItemQuickShelf) *models.INVError
	RemoveItemFromQuickShelf(tx *sql.Tx, itemId *uuid.UUID, quickShelfId *uuid.UUID) *models.INVError

	ClearQuickShelf(tx *sql.Tx, quickShelfId *uuid.UUID) *models.INVError

	GetItemsFromUserInQuickShelf(userId *uuid.UUID) (*[]model.ItemQuickShelf, *models.INVError)
	GetQuantityOfItemInQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) (*int32, *models.INVError)

	CheckIfItemAlreadyInQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) (*bool, *models.INVError)

	CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError)

	managers.DatabaseManagerI
}

type ItemQuickShelfRepository struct {
	managers.DatabaseManagerI
}

func (qsr *ItemQuickShelfRepository) DeleteItemsQuickShelf(tx *sql.Tx, itemId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.QuickShelves.DELETE().WHERE(
		table.ItemQuickShelf.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting quick shelf")
	}

	return nil
}

func (qsr *ItemQuickShelfRepository) GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]models.GetQuickShelf, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.ItemQuickShelf.AllColumns,
		table.Items.AllColumns,
		table.Users.AllColumns,
	).FROM(
		table.ItemQuickShelf.
			LEFT_JOIN(table.Items, table.Items.ID.EQ(table.ItemQuickShelf.ItemID)).
			LEFT_JOIN(table.Users, table.Users.ID.EQ(table.ItemQuickShelf.UserID)),
	).WHERE(
		table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(quickShelfId.String())),
	)

	// Execute the query
	var items []models.GetQuickShelf
	err := stmt.Query(qsr.GetDatabaseConnection(), &items)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading items in quick shelf")
	}

	return &items, nil
}

func (qsr *ItemQuickShelfRepository) InsertNewItemInQuickShelf(tx *sql.Tx, itemQuickShelf *model.ItemQuickShelf) *models.INVError {
	// Create the query
	stmt := table.ItemQuickShelf.INSERT(
		table.ItemQuickShelf.QuickShelfID,
		table.ItemQuickShelf.UserID,
		table.ItemQuickShelf.ItemID,
		table.ItemQuickShelf.Quantity,
	).VALUES(
		itemQuickShelf.QuickShelfID,
		itemQuickShelf.UserID,
		itemQuickShelf.ItemID,
		itemQuickShelf.Quantity,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error inserting item in quick shelf")
	}

	return nil
}

func (qsr *ItemQuickShelfRepository) UpdateQuantityOfItemInQuickShelf(tx *sql.Tx, itemQuickShelf *model.ItemQuickShelf) *models.INVError {
	// Create the query
	stmt := table.ItemQuickShelf.UPDATE(
		table.ItemQuickShelf.UserID,
		table.ItemQuickShelf.Quantity,
	).SET(
		itemQuickShelf.UserID,
		itemQuickShelf.Quantity,
	).WHERE(
		table.ItemQuickShelf.ItemID.EQ(mysql.String(itemQuickShelf.ItemID)).
			AND(table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(itemQuickShelf.QuickShelfID))),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating quantity of item in quick shelf")
	}

	return nil
}

func (qsr *ItemQuickShelfRepository) RemoveItemFromQuickShelf(tx *sql.Tx, itemId *uuid.UUID, quickShelfId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.ItemQuickShelf.DELETE().WHERE(
		table.ItemQuickShelf.ItemID.EQ(mysql.String(itemId.String())).
			AND(table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(quickShelfId.String()))),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error removing item from quick shelf")
	}

	return nil
}

func (qsr *ItemQuickShelfRepository) ClearQuickShelf(tx *sql.Tx, quickShelfId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.ItemQuickShelf.DELETE().WHERE(table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(quickShelfId.String())))

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error clearing quick shelf")
	}

	return nil
}

func (qsr *ItemQuickShelfRepository) GetItemsFromUserInQuickShelf(userId *uuid.UUID) (*[]model.ItemQuickShelf, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.ItemQuickShelf.AllColumns,
	).FROM(
		table.ItemQuickShelf,
	).WHERE(
		table.ItemQuickShelf.UserID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	var item []model.ItemQuickShelf
	err := stmt.Query(qsr.GetDatabaseConnection(), &item)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading items in quick shelf")
	}

	return &item, nil
}

func (qsr *ItemQuickShelfRepository) GetQuantityOfItemInQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) (*int32, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.ItemQuickShelf.Quantity,
	).FROM(
		table.ItemQuickShelf,
	).WHERE(
		table.ItemQuickShelf.ItemID.EQ(mysql.String(itemId.String())).
			AND(table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(quickShelfId.String()))),
	)

	// Execute the query TODO
	var quantity models.GetQuantity
	err := stmt.Query(qsr.GetDatabaseConnection(), &quantity)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading quantity of item in quick shelf")
	}

	return quantity.Quantity, nil
}

func (qsr *ItemQuickShelfRepository) CheckIfItemAlreadyInQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) (*bool, *models.INVError) {
	var varTrue bool = true
	var varFalse bool = false

	count, err := utils.CountStatement(table.ItemQuickShelf, table.ItemQuickShelf.ItemID.EQ(mysql.String(itemId.String())).AND(table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(quickShelfId.String()))), qsr.GetDatabaseConnection())
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if item is already in quick shelf")
	}
	if count > 0 {
		return &varTrue, nil
	}
	return &varFalse, nil
}

func (qsr *ItemQuickShelfRepository) CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError) {
	count, err := utils.CountStatement(table.ItemQuickShelf, table.ItemQuickShelf.ItemID.EQ(mysql.String(itemId.String())), qsr.GetDatabaseConnection())
	if err != nil {
		return false, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if itemId exists in ItemQuickShelf table")
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
