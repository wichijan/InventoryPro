package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type ItemKeywordRepositoryI interface {
	GetKeywordsForItems() (*[]model.KeywordsForItems, *models.INVError)
	GetKeywordsForItem(itemId *string) (*[]model.KeywordsForItems, *models.INVError)
	CreateKeywordForItem(keyword *model.KeywordsForItems) (*uuid.UUID, *models.INVError)
	UpdateKeywordForItem(keyword *model.KeywordsForItems) *models.INVError
	DeleteKeywordForItem(keywordId *uuid.UUID) *models.INVError
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

func (kfir *ItemKeywordRepository) GetKeywordsForItem(itemId *string) (*[]model.KeywordsForItems, *models.INVError) {
	var keywords []model.KeywordsForItems

	// Create the query
	stmt := mysql.SELECT(
		table.KeywordsForItems.AllColumns,
	).FROM(
		table.KeywordsForItems,
	).WHERE(
		table.KeywordsForItems.ItemID.EQ(mysql.String(*itemId)),
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

func (kfir *ItemKeywordRepository) CreateKeywordForItem(keyword *model.KeywordsForItems) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.KeywordsForItems.INSERT(
		table.KeywordsForItems.ID,
		table.KeywordsForItems.Keyword,
		table.KeywordsForItems.ItemID,
	).VALUES(
		uuid.String(),
		keyword.Keyword,
		keyword.ItemID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(kfir.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &uuid, nil
}

func (kfir *ItemKeywordRepository) UpdateKeywordForItem(keyword *model.KeywordsForItems) *models.INVError {
	// Create the update statement
	updateQuery := table.KeywordsForItems.UPDATE(
		table.KeywordsForItems.Keyword,
		table.KeywordsForItems.ItemID,
	).SET(
		keyword.Keyword,
		keyword.ItemID,
	).WHERE(table.KeywordsForItems.ID.EQ(mysql.String(keyword.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(kfir.DatabaseManager.GetDatabaseConnection())
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

func (kfir *ItemKeywordRepository) DeleteKeywordForItem(keywordId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
