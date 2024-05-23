package repositories

import (
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type UserItemRepositoryI interface {
	ReserveItem(itemReserve *models.ItemReserve) *models.INVError
	//UpdateUserRole(userRole *model.UserRoles) *models.INVError
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
