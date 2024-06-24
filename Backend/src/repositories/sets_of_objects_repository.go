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

type SetsOfObjectsRepositoryI interface {
	GetSetsOfObjectsById(setsOfObjectsId *uuid.UUID) (*model.SetsOfObjects, *models.INVError)
	CreateSetsOfObjects(tx *sql.Tx, setsOfObjects *model.SetsOfObjects) *models.INVError
	UpdateSetsOfObjects(tx *sql.Tx, setsOfObjects *model.SetsOfObjects) *models.INVError
	DeleteSetsOfObjects(tx *sql.Tx, setsOfObjectsId *uuid.UUID) *models.INVError

	CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError)

	managers.DatabaseManagerI
}

type SetsOfObjectsRepository struct {
	managers.DatabaseManagerI
}

func (sor *SetsOfObjectsRepository) GetSetsOfObjectsById(setsOfObjectsId *uuid.UUID) (*model.SetsOfObjects, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.SetsOfObjects.AllColumns,
	).FROM(
		table.SetsOfObjects,
	).WHERE(
		table.SetsOfObjects.ItemID.EQ(mysql.String(setsOfObjectsId.String())),
	)

	// Execute the query
	var setsOfObjects model.SetsOfObjects
	err := stmt.Query(sor.GetDatabaseConnection(), &setsOfObjects)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading sets of objects")
	}

	return &setsOfObjects, nil
}

func (sor *SetsOfObjectsRepository) CreateSetsOfObjects(tx *sql.Tx, setsOfObjects *model.SetsOfObjects) *models.INVError {
	// Create the query
	stmt := table.SetsOfObjects.INSERT(
		table.SetsOfObjects.ItemID,
		table.SetsOfObjects.TotalObjects,
		table.SetsOfObjects.UsefulObjects,
		table.SetsOfObjects.BrokenObjects,
		table.SetsOfObjects.LostObjects,
	).VALUES(
		setsOfObjects.ItemID,
		setsOfObjects.TotalObjects,
		setsOfObjects.UsefulObjects,
		setsOfObjects.BrokenObjects,
		setsOfObjects.LostObjects,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating sets of objects")
	}

	return nil
}

func (sor *SetsOfObjectsRepository) UpdateSetsOfObjects(tx *sql.Tx, setsOfObjects *model.SetsOfObjects) *models.INVError {
	// Create the query
	stmt := table.SetsOfObjects.UPDATE(
		table.SetsOfObjects.TotalObjects,
		table.SetsOfObjects.UsefulObjects,
		table.SetsOfObjects.BrokenObjects,
		table.SetsOfObjects.LostObjects,
	).SET(
		setsOfObjects.TotalObjects,
		setsOfObjects.UsefulObjects,
		setsOfObjects.BrokenObjects,
		setsOfObjects.LostObjects,
	).WHERE(
		table.SetsOfObjects.ItemID.EQ(mysql.String(setsOfObjects.ItemID)),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating sets of objects")
	}

	return nil
}

func (sor *SetsOfObjectsRepository) DeleteSetsOfObjects(tx *sql.Tx, setsOfObjectsId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.SetsOfObjects.DELETE().WHERE(
		table.SetsOfObjects.ItemID.EQ(mysql.String(setsOfObjectsId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting sets of objects")
	}

	return nil
}

func (sor *SetsOfObjectsRepository) CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError) {
	count, err := utils.CountStatement(table.SetsOfObjects, table.SetsOfObjects.ItemID.EQ(mysql.String(itemId.String())), sor.GetDatabaseConnection())
	if err != nil {
		return false, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if itemId exists in Sets_Of_Objects table")
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
