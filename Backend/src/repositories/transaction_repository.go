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

type TransactionRepositoryI interface {
	GetTransactionByUserId(userId *uuid.UUID) (*[]model.Transactions, *models.INVError)
	GetTransactionByItemId(itemId *uuid.UUID) (*[]model.Transactions, *models.INVError)
	CreateTransaction(tx *sql.Tx, book *model.Transactions) *models.INVError

	managers.DatabaseManagerI
}

type TransactionRepository struct {
	managers.DatabaseManagerI
}

func (trr *TransactionRepository) GetTransactionByUserId(userId *uuid.UUID) (*[]model.Transactions, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.Transactions.AllColumns,
	).FROM(
		table.Transactions,
	).WHERE(
		table.Transactions.UserID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	var transactions []model.Transactions
	err := stmt.Query(trr.GetDatabaseConnection(), &transactions)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading transactions")
	}

	return &transactions, nil
}

func (trr *TransactionRepository) GetTransactionByItemId(itemId *uuid.UUID) (*[]model.Transactions, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.Transactions.AllColumns,
	).FROM(
		table.Transactions,
	).WHERE(
		table.Transactions.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	var transactions []model.Transactions
	err := stmt.Query(trr.GetDatabaseConnection(), &transactions)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading transactions")
	}

	return &transactions, nil
}

func (trr *TransactionRepository) CreateTransaction(tx *sql.Tx, transaction *model.Transactions) *models.INVError {
	newUUID := uuid.New().String()

	// Create the query
	stmt := table.Transactions.INSERT(
		table.Transactions.TransactionID,
		table.Transactions.ItemID,
		table.Transactions.UserID,
		table.Transactions.TransactionType,
		table.Transactions.TargetUserID,
		table.Transactions.OriginUserID,
		table.Transactions.TransactionDate,
		table.Transactions.Note,
	).VALUES(
		newUUID,
		transaction.ItemID,
		transaction.UserID,
		transaction.TransactionType,
		transaction.TargetUserID,
		transaction.OriginUserID,
		utils.MysqlTimeNow(),
		transaction.Note,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}

	return nil
}
