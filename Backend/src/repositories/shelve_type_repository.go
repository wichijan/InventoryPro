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

type ShelveTypeRepositoryI interface {
	GetShelveTypes() (*[]model.ShelveTypes, *models.INVError)
	CreateShelveType(shelveTypeName *string) (*uuid.UUID, *models.INVError)
	UpdateShelveType(shelveType *model.ShelveTypes) *models.INVError
	DeleteShelveType(shelveTypeId *uuid.UUID) *models.INVError
}

type ShelveTypeRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (str *ShelveTypeRepository) GetShelveTypes() (*[]model.ShelveTypes, *models.INVError) {
	var shelveTypes []model.ShelveTypes

	// Create the query
	stmt := mysql.SELECT(
		table.ShelveTypes.AllColumns,
	).FROM(
		table.ShelveTypes,
	)

	// Execute the query
	err := stmt.Query(str.DatabaseManager.GetDatabaseConnection(), &shelveTypes)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(shelveTypes) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &shelveTypes, nil
}

func (str *ShelveTypeRepository) CreateShelveType(shelveTypeName *string) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.ShelveTypes.INSERT(
		table.ShelveTypes.ID,
		table.ShelveTypes.TypeName,
	).VALUES(
		uuid.String(),
		shelveTypeName,
	)

	// Execute the query
	rows, err := insertQuery.Exec(str.DatabaseManager.GetDatabaseConnection())
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

func (str *ShelveTypeRepository) UpdateShelveType(shelveType *model.ShelveTypes) *models.INVError {
	// Create the update statement
	updateQuery := table.ShelveTypes.UPDATE(
		table.ShelveTypes.TypeName,
	).SET(
		shelveType.TypeName,
	).WHERE(table.ShelveTypes.ID.EQ(mysql.String(shelveType.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(str.DatabaseManager.GetDatabaseConnection())
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

func (str *ShelveTypeRepository) DeleteShelveType(shelveTypeId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
