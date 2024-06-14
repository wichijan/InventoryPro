package repositories

import (
	"database/sql"
	"time"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

type UserItemRepositoryI interface {
	GetUserItemByUserIdAndItemId(userId *uuid.UUID, itemId *uuid.UUID) (*model.UserItems, *models.INVError)
	MoveItemToNewUser(tx *sql.Tx, userId *uuid.UUID, itemId *uuid.UUID, newUserId *uuid.UUID) *models.INVError

	InsertUserItem(tx *sql.Tx, itemBorrow *models.ItemBorrow) *models.INVError
	DeleteItemUser(tx *sql.Tx, userId *uuid.UUID, itemId *uuid.UUID) *models.INVError
	DeleteItemUsers(tx *sql.Tx, itemId *uuid.UUID) *models.INVError
	ReduceQuantityOfUserItem(tx *sql.Tx, userId *uuid.UUID, itemId *uuid.UUID, newQuantity *int32) *models.INVError

	GetQuantityFromUserItem(itemId *uuid.UUID) (*int32, *models.INVError)

	CheckIfItemIdExists(itemId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type UserItemRepository struct {
	managers.DatabaseManagerI
}

func (uir *UserItemRepository) GetUserItemByUserIdAndItemId(userId *uuid.UUID, itemId *uuid.UUID) (*model.UserItems, *models.INVError) {
	var userItem model.UserItems

	// Create the query
	stmt := mysql.SELECT(
		table.UserItems.AllColumns,
	).FROM(
		table.UserItems,
	).WHERE(
		table.UserItems.UserID.EQ(mysql.String(userId.String())).
			AND(table.UserItems.ItemID.EQ(mysql.String(itemId.String()))),
	)

	// Execute the query
	err := stmt.Query(uir.GetDatabaseConnection(), &userItem)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND.WithDetails("User item not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading user item")
	}

	return &userItem, nil
}

func (uir *UserItemRepository) GetQuantityFromUserItem(itemId *uuid.UUID) (*int32, *models.INVError) {
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
			return nil, inv_errors.INV_NOT_FOUND.WithDetails("Item not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading quantity of item")
	}

	return &quantity.Quantity, nil
}

func (uir *UserItemRepository) InsertUserItem(tx *sql.Tx, itemBorrow *models.ItemBorrow) *models.INVError {
	// Create the insert statement
	insertQuery := table.UserItems.INSERT(
		table.UserItems.UserID,
		table.UserItems.ItemID,
		table.UserItems.TransactionDate,
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error inserting user item")
	}

	return nil
}

func (uir *UserItemRepository) DeleteItemUser(tx *sql.Tx, userId *uuid.UUID, itemId *uuid.UUID) *models.INVError {
	// Create the delete statement
	deleteQuery := table.UserItems.DELETE().WHERE(
		table.UserItems.UserID.EQ(mysql.String(userId.String())).
			AND(table.UserItems.ItemID.EQ(mysql.String(itemId.String()))),
	)

	// Execute the query
	_, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting user item")
	}

	return nil
}

func (uir *UserItemRepository) DeleteItemUsers(tx *sql.Tx, itemId *uuid.UUID) *models.INVError {
	// Create the delete statement
	deleteQuery := table.UserItems.DELETE().WHERE(
		table.UserItems.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	_, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting user item")
	}

	return nil
}

func (uir *UserItemRepository) MoveItemToNewUser(tx *sql.Tx, userId *uuid.UUID, itemId *uuid.UUID, newUserId *uuid.UUID) *models.INVError {
	// Create the update statement
	updateQuery := table.UserItems.UPDATE(
		table.UserItems.UserID,
	).SET(
		newUserId,
	).WHERE(
		table.UserItems.UserID.EQ(mysql.String(userId.String())).
			AND(table.UserItems.ItemID.EQ(mysql.String(itemId.String()))),
	)

	// Execute the query
	_, err := updateQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error moving item to new user")
	}

	return nil
}

func (uir *UserItemRepository) ReduceQuantityOfUserItem(tx *sql.Tx, userId *uuid.UUID, itemId *uuid.UUID, newQuantity *int32) *models.INVError {
	// Create the update statement
	updateQuery := table.UserItems.UPDATE(
		table.UserItems.Quantity,
		table.UserItems.TransactionDate,
	).SET(
		newQuantity,
		time.Now(),
	).WHERE(
		table.UserItems.UserID.EQ(mysql.String(userId.String())).
			AND(table.UserItems.ItemID.EQ(mysql.String(itemId.String()))),
	)

	// Execute the query
	_, err := updateQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reducing quantity of user item")
	}

	return nil
}

func (uir *UserItemRepository) CheckIfItemIdExists(itemId *uuid.UUID) *models.INVError {
	count, err := utils.CountStatement(table.UserItems, table.UserItems.ItemID.EQ(mysql.String(itemId.String())), uir.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if itemId exists in UserItems table")
	}
	if count <= 0 {
		return inv_errors.INV_CONFLICT.WithDetails("UserItems still has items in it")
	}
	return nil
}
