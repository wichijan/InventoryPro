package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

type ItemStatusRepositoryI interface {
	GetItemStatus() (*[]model.ItemStatus, *models.INVError)
	GetStatusIdByName(statusName *string) (*uuid.UUID, *models.INVError)
	CreateItemStatus(itemStatusName *string) (*uuid.UUID, *models.INVError)
	UpdateItemStatus(itemStatus *model.ItemStatus) *models.INVError
	DeleteItemStatus(itemStatusId *uuid.UUID) *models.INVError
}

type ItemStatusRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (isr *ItemStatusRepository) GetItemStatus() (*[]model.ItemStatus, *models.INVError) {
	var itemStatus []model.ItemStatus

	// Create the query
	stmt := mysql.SELECT(
		table.ItemStatus.AllColumns,
	).FROM(
		table.ItemStatus,
	)

	// Execute the query
	err := stmt.Query(isr.DatabaseManager.GetDatabaseConnection(), &itemStatus)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(itemStatus) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &itemStatus, nil
}

func (isr *ItemStatusRepository) GetStatusIdByName(statusName *string) (*uuid.UUID, *models.INVError) {
	var itemStatus uuid.UUID

	// Create the query
	stmt := mysql.SELECT(
		table.ItemStatus.ID,
	).FROM(
		table.ItemStatus,
	).WHERE(
		table.ItemStatus.StatusName.EQ(utils.MySqlString(*statusName)),
	)

	// Execute the query
	err := stmt.Query(isr.DatabaseManager.GetDatabaseConnection(), &itemStatus)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &itemStatus, nil
}

func (isr *ItemStatusRepository) CreateItemStatus(itemStatusName *string) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.ItemStatus.INSERT(
		table.ItemStatus.ID,
		table.ItemStatus.StatusName,
	).VALUES(
		uuid.String(),
		itemStatusName,
	)

	// Execute the query
	rows, err := insertQuery.Exec(isr.DatabaseManager.GetDatabaseConnection())
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

func (isr *ItemStatusRepository) UpdateItemStatus(itemStatus *model.ItemStatus) *models.INVError {
	// Create the update statement
	updateQuery := table.ItemStatus.UPDATE(
		table.ItemStatus.StatusName,
	).SET(
		itemStatus.StatusName,
	).WHERE(table.ItemStatus.ID.EQ(mysql.String(itemStatus.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(isr.DatabaseManager.GetDatabaseConnection())
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

func (isr *ItemStatusRepository) DeleteItemStatus(itemStatusId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
