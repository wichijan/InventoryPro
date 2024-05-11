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

type ItemPictureRepositoryI interface {
	GetItemPictures(itemId *string) (*[]model.ItemPictures, *models.INVError)
	CreateItemPicture(itemPicture *model.ItemPictures) (*uuid.UUID, *models.INVError)
	UpdateItemPicture(itemPicture *model.ItemPictures) *models.INVError
	DeleteItemPicture(itemIdPicture *uuid.UUID) *models.INVError
}

type ItemPictureRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (ipr *ItemPictureRepository) GetItemPictures(itemId *string) (*[]model.ItemPictures, *models.INVError) {
	var itemPictures []model.ItemPictures

	// Create the query
	stmt := mysql.SELECT(
		table.ItemPictures.AllColumns,
	).FROM(
		table.ItemPictures,
	).WHERE(table.ItemPictures.ID.EQ(mysql.String(*itemId)))

	// Execute the query
	err := stmt.Query(ipr.DatabaseManager.GetDatabaseConnection(), &itemPictures)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(itemPictures) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &itemPictures, nil
}

func (ipr *ItemPictureRepository) CreateItemPicture(itemPicture *model.ItemPictures) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.ItemPictures.INSERT(
		table.ItemPictures.ID,
		table.ItemPictures.Picture,
		table.ItemPictures.ItemID,
	).VALUES(
		uuid.String(),
		itemPicture.Picture,
		itemPicture.ItemID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(ipr.DatabaseManager.GetDatabaseConnection())
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

func (ipr *ItemPictureRepository) UpdateItemPicture(itemPicture *model.ItemPictures) *models.INVError {
	// Create the update statement
	updateQuery := table.ItemPictures.UPDATE(
		table.ItemPictures.Picture,
		table.ItemPictures.ItemID,
	).SET(
		itemPicture.Picture,
		itemPicture.ItemID,
	).WHERE(table.ItemPictures.ID.EQ(mysql.String(itemPicture.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(ipr.DatabaseManager.GetDatabaseConnection())
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

func (ipr *ItemPictureRepository) DeleteItemPicture(itemIdPicture *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
