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

type ItemInShelveRepositoryI interface {
	GetItemsInShelf(shelfID *string) (*[]model.ItemsInShelf, *models.INVError)
	CreateItemInShelve(tx *sql.Tx, itemInShelve *model.ItemsInShelf) *models.INVError
	UpdateItemInShelve(tx *sql.Tx, itemInShelve *model.ItemsInShelf) *models.INVError
	DeleteItemInShelve(tx *sql.Tx, itemIdInShelve *uuid.UUID) *models.INVError
	GetQuantityInShelve(itemId *uuid.UUID) (*int32, *models.INVError)
	UpdateQuantityInShelve(tx *sql.Tx, itemId *string, quantity *int32) *models.INVError

	managers.DatabaseManagerI
}

type ItemInShelveRepository struct {
	managers.DatabaseManagerI
}

func (iisr *ItemInShelveRepository) GetItemsInShelf(shelfID *string) (*[]model.ItemsInShelf, *models.INVError) {
	var ItemsInShelf []model.ItemsInShelf

	// Create the query
	stmt := mysql.SELECT(
		table.ItemsInShelf.AllColumns,
	).FROM(
		table.ItemsInShelf,
	).WHERE(table.ItemsInShelf.ShelfID.EQ(mysql.String(*shelfID)))

	// Execute the query
	err := stmt.Query(iisr.GetDatabaseConnection(), &ItemsInShelf)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(ItemsInShelf) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &ItemsInShelf, nil
}

func (iisr *ItemInShelveRepository) CreateItemInShelve(tx *sql.Tx, itemInShelve *model.ItemsInShelf) *models.INVError {

	// Create the insert statement
	insertQuery := table.ItemsInShelf.INSERT(
		table.ItemsInShelf.ItemID,
		table.ItemsInShelf.ShelfID,
		table.ItemsInShelf.Quantity,
	).VALUES(
		itemInShelve.ItemID,
		itemInShelve.ShelfID,
		itemInShelve.Quantity,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
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

func (iisr *ItemInShelveRepository) UpdateItemInShelve(tx *sql.Tx, itemInShelve *model.ItemsInShelf) *models.INVError {
	// Create the update statement
	updateQuery := table.ItemsInShelf.UPDATE(
		table.ItemsInShelf.ItemID,
		table.ItemsInShelf.ShelfID,
		table.ItemsInShelf.Quantity,
	).SET(
		itemInShelve.ItemID,
		itemInShelve.ShelfID,
		itemInShelve.Quantity,
	).WHERE(table.ItemsInShelf.ItemID.EQ(mysql.String(itemInShelve.ItemID)).
		AND(table.ItemsInShelf.ShelfID.EQ(mysql.String(itemInShelve.ShelfID))))

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

func (iisr *ItemInShelveRepository) DeleteItemInShelve(tx *sql.Tx, itemIdInShelve *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}

func (iisr *ItemInShelveRepository) GetQuantityInShelve(itemId *uuid.UUID) (*int32, *models.INVError) {

	// Create the query
	stmt := mysql.SELECT(
		table.ItemsInShelf.Quantity,
	).FROM(
		table.ItemsInShelf,
	).WHERE(
		table.ItemsInShelf.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	var quantity models.GetQuantityInShelve
	err := stmt.Query(iisr.GetDatabaseConnection(), &quantity)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &quantity.Quantity, nil
}

func (iisr *ItemInShelveRepository) UpdateQuantityInShelve(tx *sql.Tx, itemId *string, quantity *int32) *models.INVError {
	// Create the update statement
	updateQuery := table.ItemsInShelf.UPDATE(
		table.ItemsInShelf.Quantity,
	).SET(
		quantity,
	).WHERE(table.ItemsInShelf.ItemID.EQ(mysql.String(*itemId)))

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
