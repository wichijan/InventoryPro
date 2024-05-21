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

type ItemInShelveRepositoryI interface {
	GetItemsInShelve(shelveId *string) (*[]model.ItemsInShelve, *models.INVError)
	CreateItemInShelve(itemInShelve *model.ItemsInShelve) *models.INVError
	UpdateItemInShelve(itemInShelve *model.ItemsInShelve) *models.INVError
	DeleteItemInShelve(itemIdInShelve *uuid.UUID) *models.INVError
}

type ItemInShelveRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (iisr *ItemInShelveRepository) GetItemsInShelve(shelveId *string) (*[]model.ItemsInShelve, *models.INVError) {
	var itemsInShelve []model.ItemsInShelve

	// Create the query
	stmt := mysql.SELECT(
		table.ItemsInShelve.AllColumns,
	).FROM(
		table.ItemsInShelve,
	).WHERE(table.ItemsInShelve.ShelveID.EQ(mysql.String(*shelveId)))

	// Execute the query
	err := stmt.Query(iisr.DatabaseManager.GetDatabaseConnection(), &itemsInShelve)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(itemsInShelve) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &itemsInShelve, nil
}

func (iisr *ItemInShelveRepository) CreateItemInShelve(itemInShelve *model.ItemsInShelve) *models.INVError {

	// Create the insert statement
	insertQuery := table.ItemsInShelve.INSERT(
		table.ItemsInShelve.ItemID,
		table.ItemsInShelve.ShelveID,
	).VALUES(
		itemInShelve.ItemID,
		itemInShelve.ShelveID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(iisr.DatabaseManager.GetDatabaseConnection())
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

func (iisr *ItemInShelveRepository) UpdateItemInShelve(itemInShelve *model.ItemsInShelve) *models.INVError {
	// Create the update statement
	updateQuery := table.ItemsInShelve.UPDATE(
		table.ItemsInShelve.ItemID,
		table.ItemsInShelve.ShelveID,
	).SET(
		itemInShelve.ItemID,
		itemInShelve.ShelveID,
	).WHERE(table.ItemsInShelve.ItemID.EQ(mysql.String(itemInShelve.ItemID)).
		AND(table.ItemsInShelve.ShelveID.EQ(mysql.String(itemInShelve.ShelveID))))

	// Execute the query
	rows, err := updateQuery.Exec(iisr.DatabaseManager.GetDatabaseConnection())
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

func (iisr *ItemInShelveRepository) DeleteItemInShelve(itemIdInShelve *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
