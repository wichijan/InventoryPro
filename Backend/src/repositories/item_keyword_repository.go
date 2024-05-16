package repositories

import (
	"log"

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
	CreateKeywordForItem(keyword *models.ItemWithKeyword) (*uuid.UUID, *models.INVError)
	DeleteKeywordForItem(keyword *models.ItemWithKeyword) *models.INVError
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

	log.Print("GetKeywords Repo")

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

func (kfir *ItemKeywordRepository) CreateKeywordForItem(keyword *models.ItemWithKeyword) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.KeywordsForItems.INSERT(
		table.KeywordsForItems.ID,
		table.KeywordsForItems.KeywordID,
		table.KeywordsForItems.ItemID,
	).VALUES(
		uuid.String(),
		keyword.KeywordID,
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

func (kfir *ItemKeywordRepository) DeleteKeywordForItem(keyword *models.ItemWithKeyword) *models.INVError {
	deleteQuery := table.KeywordsForItems.DELETE().WHERE(
		table.KeywordsForItems.ID.EQ(mysql.String(keyword.KeywordID)).
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
