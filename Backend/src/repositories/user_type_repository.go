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

type UserTypeRepositoryI interface {
	GetUserTypes() (*[]model.UserTypes, *models.INVError)
	CreateUserType(type_name *string) (*uuid.UUID, *models.INVError)
	UpdateUserType(userType *model.UserTypes) *models.INVError
	DeleteUserType(userTypeId *uuid.UUID) *models.INVError
}

type UserTypeRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (utr *UserTypeRepository) GetUserTypes() (*[]model.UserTypes, *models.INVError) {
	var userTypes []model.UserTypes

	// Create the query
	stmt := mysql.SELECT(
		table.UserTypes.AllColumns,
	).FROM(
		table.UserTypes,
	)

	// Execute the query
	err := stmt.Query(utr.DatabaseManager.GetDatabaseConnection(), &userTypes)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(userTypes) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &userTypes, nil
}

func (utr *UserTypeRepository) CreateUserType(type_name *string) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.UserTypes.INSERT(
		table.UserTypes.ID,
		table.UserTypes.TypeName,
	).VALUES(
		uuid.String(),
		type_name,
	)

	// Execute the query
	rows, err := insertQuery.Exec(utr.DatabaseManager.GetDatabaseConnection())
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

func (utr *UserTypeRepository) UpdateUserType(userType *model.UserTypes) *models.INVError {
	// Create the update statement
	updateQuery := table.UserTypes.UPDATE(
		table.UserTypes.TypeName,
	).SET(
		userType.TypeName,
	).WHERE(table.UserTypes.ID.EQ(mysql.String(userType.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(utr.DatabaseManager.GetDatabaseConnection())
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

func (utr *UserTypeRepository) DeleteUserType(userTypeId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
