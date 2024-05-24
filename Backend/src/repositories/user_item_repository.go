package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type UserItemRepositoryI interface {
	ReserveItem(itemReserve *models.ItemReserve) *models.INVError
	DeleteReserveItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError

	GetQuantityFromReservedItem(itemId *uuid.UUID) (*int32, *models.INVError)
}

type UserItemRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (uir *UserItemRepository) ReserveItem(itemReserve *models.ItemReserve) *models.INVError {
	// Create the insert statement
	insertQuery := table.UserItems.INSERT(
		table.UserItems.UserID,
		table.UserItems.ItemID,
		table.UserItems.ReservedDate,
		table.UserItems.Quantity,
		table.UserItems.StatusID,
	).VALUES(
		itemReserve.UserID,
		itemReserve.ItemID,
		itemReserve.ReserveDate,
		itemReserve.Quantity,
		itemReserve.StatusID,
	)

	// Execute the query
	_, err := insertQuery.Exec(uir.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (uir *UserItemRepository) DeleteReserveItem(userId *uuid.UUID, itemId *uuid.UUID) *models.INVError {
	// Create the delete statement
	deleteQuery := table.UserItems.DELETE().WHERE(
		table.UserItems.UserID.EQ(mysql.String(userId.String())).
			AND(table.UserItems.ItemID.EQ(mysql.String(itemId.String()))),
	)

	// Execute the query
	_, err := deleteQuery.Exec(uir.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
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
	err := stmt.Query(uir.DatabaseManager.GetDatabaseConnection(), &quantity)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &quantity.Quantity, nil
}
