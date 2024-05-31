package repositories

import (
	"database/sql"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type UserItemRepositoryI interface {
	BorrowItem(tx *sql.Tx, itemBorrow *models.ItemBorrow) *models.INVError
	ReturnItem(tx *sql.Tx, userId *uuid.UUID, itemId *uuid.UUID) *models.INVError

	GetQuantityFromReservedItem(itemId *uuid.UUID) (*int32, *models.INVError)

	managers.DatabaseManagerI
}

type UserItemRepository struct {
	managers.DatabaseManagerI
}


func (uir *UserItemRepository) GetQuantityFromReservedItem(itemId *uuid.UUID) (*int32, *models.INVError) {
	var quantity models.GetQuantityReserved

	// Create the query
	stmt := mysql.SELECT(
		table.UserItems.Quantity,
	).FROM(
		table.UserItems,
	).WHERE(
		table.UserItems.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	err := stmt.Query(uir.GetDatabaseConnection(), &quantity)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &quantity.Quantity, nil
}

func (uir *UserItemRepository) BorrowItem(tx *sql.Tx, itemBorrow *models.ItemBorrow) *models.INVError {
	// Create the insert statement
	insertQuery := table.UserItems.INSERT(
		table.UserItems.UserID,
		table.UserItems.ItemID,
		table.UserItems.Quantity,
	).VALUES(
		itemBorrow.UserID,
		itemBorrow.ItemID,
		itemBorrow.TransactionDate,
		itemBorrow.Quantity,
	)

	// Execute the query
	_, err := insertQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (uir *UserItemRepository) ReturnItem(tx *sql.Tx, userId *uuid.UUID, itemId *uuid.UUID) *models.INVError {
	// Create the delete statement
	deleteQuery := table.UserItems.DELETE().WHERE(
		table.UserItems.UserID.EQ(mysql.String(userId.String())).
			AND(table.UserItems.ItemID.EQ(mysql.String(itemId.String()))),
	)

	// Execute the query
	_, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}
