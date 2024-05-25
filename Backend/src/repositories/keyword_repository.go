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

type KeywordRepositoryI interface {
	GetKeywords() (*[]model.Keywords, *models.INVError)
	GetKeywordByName(keywordName *string) (*model.Keywords, *models.INVError)
	CheckIfKeywordExists(keywordId *string) *models.INVError
	CreateKeyword(tx *sql.Tx, keywordName *string) (*uuid.UUID, *models.INVError)
	UpdateKeyword(tx *sql.Tx, keyword *model.Keywords) *models.INVError
	DeleteKeyword(tx *sql.Tx, keywordId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type KeywordRepository struct {
	managers.DatabaseManagerI
}

func (kr *KeywordRepository) GetKeywords() (*[]model.Keywords, *models.INVError) {
	var keywords []model.Keywords

	// Create the query
	stmt := mysql.SELECT(
		table.Keywords.AllColumns,
	).FROM(
		table.Keywords,
	)

	// Execute the query
	err := stmt.Query(kr.GetDatabaseConnection(), &keywords)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(keywords) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &keywords, nil
}

func (kr *KeywordRepository) GetKeywordByName(keywordName *string) (*model.Keywords, *models.INVError) {
	var keyword model.Keywords

	// Create the query
	stmt := mysql.SELECT(
		table.Keywords.AllColumns,
	).FROM(
		table.Keywords,
	).WHERE(
		table.Keywords.Keyword.EQ(mysql.String(*keywordName)),
	)

	// Execute the query
	err := stmt.Query(kr.GetDatabaseConnection(), &keyword)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if keyword.Keyword == nil {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &keyword, nil
}

func (kr *KeywordRepository) CreateKeyword(tx *sql.Tx, keywordName *string) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the query
	stmt := table.Keywords.INSERT(
		table.Keywords.ID,
		table.Keywords.Keyword,
	).VALUES(
		uuid.String(),
		keywordName,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &uuid, nil
}

func (kr *KeywordRepository) UpdateKeyword(tx *sql.Tx, keyword *model.Keywords) *models.INVError {
	// Create the query
	stmt := table.Keywords.UPDATE(
		table.Keywords.Keyword,
	).SET(
		keyword.Keyword,
	).WHERE(
		table.Keywords.ID.EQ(mysql.String(keyword.ID)),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (kr *KeywordRepository) DeleteKeyword(tx *sql.Tx, keywordId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.Keywords.DELETE().WHERE(
		table.Keywords.ID.EQ(mysql.String(keywordId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (kr *KeywordRepository) CheckIfKeywordExists(keywordId *string) *models.INVError {
	count, err := utils.CountStatement(table.Keywords, table.Keywords.ID.EQ(mysql.String(*keywordId)), kr.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	if count > 0 {
		return inv_errors.INV_USERNAME_EXISTS
	}
	return nil
}
