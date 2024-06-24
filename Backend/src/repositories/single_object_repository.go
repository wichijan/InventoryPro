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

type SingleObjectRepositoryI interface {
	GetSingleObjectById(singleObjectId *uuid.UUID) (*model.SingleObject, *models.INVError)
	CreateSingleObject(tx *sql.Tx, singleObject *model.SingleObject) *models.INVError
	// UpdateSingleObject(tx *sql.Tx, singleObject *model.SingleObject) *models.INVError // Not needed because there are not attributes
	DeleteSingleObject(tx *sql.Tx, itemId *uuid.UUID) *models.INVError
	CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError)

	managers.DatabaseManagerI
}

type SingleObjectRepository struct {
	managers.DatabaseManagerI
}

func (sor *SingleObjectRepository) GetSingleObjectById(singleObjectId *uuid.UUID) (*model.SingleObject, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.SingleObject.AllColumns,
	).FROM(
		table.SingleObject,
	).WHERE(
		table.SingleObject.ItemID.EQ(mysql.String(singleObjectId.String())),
	)

	// Execute the query
	var singleObject model.SingleObject
	err := stmt.Query(sor.GetDatabaseConnection(), &singleObject)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading single object")
	}

	return &singleObject, nil
}

func (sor *SingleObjectRepository) CreateSingleObject(tx *sql.Tx, singleObject *model.SingleObject) *models.INVError {
	// Create the query
	stmt := table.SingleObject.INSERT(
		table.SingleObject.ItemID,
	).VALUES(
		singleObject.ItemID,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating single object")
	}

	return nil
}

func (sor *SingleObjectRepository) DeleteSingleObject(tx *sql.Tx, itemId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.SingleObject.DELETE().WHERE(
		table.SingleObject.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting single object")
	}

	return nil
}

func (sor *SingleObjectRepository) CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError) {
	count, err := utils.CountStatement(table.SingleObject, table.SingleObject.ItemID.EQ(mysql.String(itemId.String())), sor.GetDatabaseConnection())
	if err != nil {
		return false, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if itemId exists in Single_Object table")
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}


