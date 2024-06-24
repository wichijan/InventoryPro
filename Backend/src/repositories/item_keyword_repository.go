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

type ItemKeywordRepositoryI interface {
	GetKeywordsForItems() (*[]model.KeywordsForItems, *models.INVError)
	CreateKeywordForItem(tx *sql.Tx, keyword *models.ItemWithKeyword) *models.INVError
	DeleteKeywordForItem(tx *sql.Tx, keyword *models.ItemWithKeyword) *models.INVError
	DeleteKeywordsForItem(tx *sql.Tx, itemId *uuid.UUID) *models.INVError
	CheckIfKeywordAndItemExists(keywordAndItem models.ItemWithKeyword) *models.INVError

	CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError)

	managers.DatabaseManagerI
}

type ItemKeywordRepository struct {
	managers.DatabaseManagerI
}

func (kfir *ItemKeywordRepository) GetKeywordsForItems() (*[]model.KeywordsForItems, *models.INVError) {
	var keywords []model.KeywordsForItems

	// Create the query
	stmt := mysql.SELECT(
		table.KeywordsForItems.AllColumns,
	).FROM(
		table.KeywordsForItems,
	)

	// Execute the query
	err := stmt.Query(kfir.GetDatabaseConnection(), &keywords)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading keywords for items")
	}

	return &keywords, nil
}

func (kfir *ItemKeywordRepository) CreateKeywordForItem(tx *sql.Tx, keyword *models.ItemWithKeyword) *models.INVError {

	// Create the insert statement
	insertQuery := table.KeywordsForItems.INSERT(
		table.KeywordsForItems.KeywordID,
		table.KeywordsForItems.ItemID,
	).VALUES(
		keyword.KeywordID,
		keyword.ItemID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating keyword for item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating keyword for item")
	}

	if rowsAff == 0 {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Combination already exists")
	}

	return nil
}

func (kfir *ItemKeywordRepository) DeleteKeywordForItem(tx *sql.Tx, keyword *models.ItemWithKeyword) *models.INVError {
	deleteQuery := table.KeywordsForItems.DELETE().WHERE(
		table.KeywordsForItems.KeywordID.EQ(mysql.String(keyword.KeywordID)).
			AND(table.KeywordsForItems.ItemID.EQ(mysql.String(keyword.ItemID))),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting keyword for item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting keyword for item")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Combination not found")
	}

	return nil
}

func (kfir *ItemKeywordRepository) DeleteKeywordsForItem(tx *sql.Tx, itemId *uuid.UUID) *models.INVError {
	deleteQuery := table.KeywordsForItems.DELETE().WHERE(
		table.KeywordsForItems.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting keyword for item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting keyword for item")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Combination not found")
	}

	return nil
}

func (kfir *ItemKeywordRepository) CheckIfKeywordAndItemExists(keywordAndItem models.ItemWithKeyword) *models.INVError {
	count, err := utils.CountStatement(table.KeywordsForItems, table.KeywordsForItems.KeywordID.EQ(mysql.String(keywordAndItem.KeywordID)).AND(table.KeywordsForItems.ItemID.EQ(mysql.String(keywordAndItem.ItemID))), kfir.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if keyword and item exists")
	}
	if count > 0 {
		return inv_errors.INV_CONFLICT.WithDetails("Keyword and item combination already exists")
	}
	return nil
}

func (kfir *ItemKeywordRepository) CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError) {
	count, err := utils.CountStatement(table.KeywordsForItems, table.KeywordsForItems.ItemID.EQ(mysql.String(itemId.String())), kfir.GetDatabaseConnection())
	if err != nil {
		return false, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if itemId exists in KeywordsForItems table")
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
