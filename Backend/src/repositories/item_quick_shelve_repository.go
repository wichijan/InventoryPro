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
	GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]models.ItemQuickShelfInsert, *models.INVError)
	InsertNewItemInQuickShelf(tx *sql.Tx, itemQuickShelf *model.ItemQuickShelf) *models.INVError
	UpdateQuantityOfItemInQuickShelf(tx *sql.Tx, itemQuickShelf *model.ItemQuickShelf) *models.INVError
	RemoveItemFromQuickShelf(tx *sql.Tx, itemId *uuid.UUID, quickShelfId *uuid.UUID) *models.INVError

	ClearQuickShelf(tx *sql.Tx, quickShelfId *uuid.UUID) *models.INVError

	GetItemsFromUserInQuickShelf(userId *uuid.UUID) (*[]model.ItemQuickShelf, *models.INVError)
	GetQuantityOfItemInQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) (*int32, *models.INVError)

	CheckIfItemAlreadyInQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) (*bool, *models.INVError)

	managers.DatabaseManagerI
}

type ItemQuickShelfRepository struct {
	managers.DatabaseManagerI
}

func (qsr *ItemQuickShelfRepository) GetItemsInQuickShelf(quickShelfId *uuid.UUID) (*[]models.ItemQuickShelfInsert, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.ItemQuickShelf.AllColumns,
	).FROM(
		table.ItemQuickShelf,
	).WHERE(
		table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(quickShelfId.String())),
	)

	// Execute the query
	var items []models.ItemQuickShelfInsert
	err := stmt.Query(qsr.GetDatabaseConnection(), &items)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
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
		return inv_errors.INV_INTERNAL_ERROR
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
		return inv_errors.INV_INTERNAL_ERROR
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
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (qsr *ItemQuickShelfRepository) ClearQuickShelf(tx *sql.Tx, quickShelfId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.ItemQuickShelf.DELETE().WHERE(table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(quickShelfId.String())))

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
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
		return nil, inv_errors.INV_INTERNAL_ERROR
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
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return quantity.Quantity, nil
}

func (qsr *ItemQuickShelfRepository) CheckIfItemAlreadyInQuickShelf(itemId *uuid.UUID, quickShelfId *uuid.UUID) (*bool, *models.INVError) {
	var varTrue bool = true
	var varFalse bool = false

	count, err := utils.CountStatement(table.ItemQuickShelf, table.ItemQuickShelf.ItemID.EQ(mysql.String(itemId.String())).AND(table.ItemQuickShelf.QuickShelfID.EQ(mysql.String(quickShelfId.String()))), qsr.GetDatabaseConnection())
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	if count > 0 {
		return &varTrue, inv_errors.INV_USERNAME_EXISTS
	}
	return &varFalse, nil
}
