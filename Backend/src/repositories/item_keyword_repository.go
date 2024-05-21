package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

type ItemKeywordRepositoryI interface {
	GetKeywordsForItems() (*[]model.KeywordsForItems, *models.INVError)
	CreateKeywordForItem(keyword *models.ItemWithKeyword) *models.INVError
	DeleteKeywordForItem(keyword *models.ItemWithKeyword) *models.INVError
	CheckIfKeywordAndItemExists(keywordAndItem models.ItemWithKeyword) *models.INVError
}

type ItemKeywordRepository struct {
	DatabaseManager managers.DatabaseManagerI
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
	err := stmt.Query(kfir.DatabaseManager.GetDatabaseConnection(), &keywords)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(keywords) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &keywords, nil
}

func (kfir *ItemKeywordRepository) CreateKeywordForItem(keyword *models.ItemWithKeyword) *models.INVError {

	// Create the insert statement
	insertQuery := table.KeywordsForItems.INSERT(
		table.KeywordsForItems.KeywordID,
		table.KeywordsForItems.ItemID,
	).VALUES(
		keyword.KeywordID,
		keyword.ItemID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(kfir.DatabaseManager.GetDatabaseConnection())
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

func (kfir *ItemKeywordRepository) DeleteKeywordForItem(keyword *models.ItemWithKeyword) *models.INVError {
	deleteQuery := table.KeywordsForItems.DELETE().WHERE(
		table.KeywordsForItems.KeywordID.EQ(mysql.String(keyword.KeywordID)).
			AND(table.KeywordsForItems.ItemID.EQ(mysql.String(keyword.ItemID))),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(kfir.DatabaseManager.GetDatabaseConnection())
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

func (kfir *ItemKeywordRepository) CheckIfKeywordAndItemExists(keywordAndItem models.ItemWithKeyword) *models.INVError {
	count, err := utils.CountStatement(table.KeywordsForItems, table.KeywordsForItems.KeywordID.EQ(mysql.String(keywordAndItem.KeywordID)).AND(table.KeywordsForItems.ItemID.EQ(mysql.String(keywordAndItem.ItemID))), kfir.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	if count > 0 {
		return inv_errors.INV_KEYWORDS_ITEM_COMBI_EXISTS
	}
	return nil
}
