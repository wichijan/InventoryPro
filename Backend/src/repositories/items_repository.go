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

type IteitrepositoryI interface {
	GetItems() (*[]model.Items, *models.INVError)
	CreateItem(item *model.Items) (*uuid.UUID, *models.INVError)
	UpdateItem(item *model.Items) *models.INVError
	DeleteItem(itemId *uuid.UUID) *models.INVError
}

type Iteitrepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (itr *Iteitrepository) GetItems() (*[]model.Items, *models.INVError) {
	var items []model.Items

	// Create the query
	stmt := mysql.SELECT(
		table.Items.AllColumns,
	).FROM(
		table.Items,
	)

	// Execute the query
	err := stmt.Query(itr.DatabaseManager.GetDatabaseConnection(), &items)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(items) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &items, nil
}

func (itr *Iteitrepository) CreateItem(item *model.Items) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Items.INSERT(
		table.Items.ID,
		table.Items.Name,
		table.Items.Description,
		table.Items.ClassOne,
		table.Items.ClassTwo,
		table.Items.ClassThree,
		table.Items.ClassFour,
		table.Items.Damaged,
		table.Items.DamagedDescription,
		table.Items.Quantity,
		table.Items.StatusID,
	).VALUES(
		uuid.String(),
		item.Name,
		item.Description,
		item.ClassOne,
		item.ClassTwo,
		item.ClassThree,
		item.ClassFour,
		item.Damaged,
		item.DamagedDescription,
		item.Quantity,
		item.StatusID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(itr.DatabaseManager.GetDatabaseConnection())
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

func (itr *Iteitrepository) UpdateItem(item *model.Items) *models.INVError {
	// Create the update statement
	updateQuery := table.Items.UPDATE(
		table.Items.Name,
		table.Items.Description,
		table.Items.ClassOne,
		table.Items.ClassTwo,
		table.Items.ClassThree,
		table.Items.ClassFour,
		table.Items.Damaged,
		table.Items.DamagedDescription,
		table.Items.Quantity,
		table.Items.StatusID,
	).SET(
		item.Name,
		item.Description,
		item.ClassOne,
		item.ClassTwo,
		item.ClassThree,
		item.ClassFour,
		item.Damaged,
		item.DamagedDescription,
		item.Quantity,
		item.StatusID,
	).WHERE(table.Items.ID.EQ(mysql.String(item.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(itr.DatabaseManager.GetDatabaseConnection())
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

func (itr *Iteitrepository) DeleteItem(itemId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
