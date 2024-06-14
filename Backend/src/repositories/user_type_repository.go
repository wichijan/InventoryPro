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
)

type UserTypeRepositoryI interface {
	GetUserTypes() (*[]model.UserTypes, *models.INVError)
	GetUserTypeById(id *string) (*string, *models.INVError)
	GetUserTypeByName(name *string) (*string, *models.INVError)
	CreateUserType(tx *sql.Tx, type_name *string) (*uuid.UUID, *models.INVError)
	UpdateUserType(tx *sql.Tx, userType *model.UserTypes) *models.INVError
	DeleteUserType(tx *sql.Tx, userTypeId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type UserTypeRepository struct {
	managers.DatabaseManagerI
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
	err := stmt.Query(utr.GetDatabaseConnection(), &userTypes)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading user types")
	}

	return &userTypes, nil
}

func (utr *UserTypeRepository) GetUserTypeById(id *string) (*string, *models.INVError) {
	var userTypes model.UserTypes

	// Create the query
	stmt := mysql.SELECT(
		table.UserTypes.AllColumns,
	).FROM(
		table.UserTypes,
	).WHERE(
		table.UserTypes.ID.EQ(mysql.String(*id)),
	)

	// Execute the query
	err := stmt.Query(utr.GetDatabaseConnection(), &userTypes)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading user type")
	}

	if userTypes.TypeName == nil {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("User type not found")
	}

	return userTypes.TypeName, nil
}

func (utr *UserTypeRepository) GetUserTypeByName(name *string) (*string, *models.INVError) {
	var userTypes model.UserTypes

	// Create the query
	stmt := mysql.SELECT(
		table.UserTypes.AllColumns,
	).FROM(
		table.UserTypes,
	).WHERE(
		table.UserTypes.TypeName.EQ(mysql.String(*name)),
	)

	// Execute the query
	err := stmt.Query(utr.GetDatabaseConnection(), &userTypes)
	if err != nil {
		return nil, inv_errors.INV_CONFLICT.WithDetails("User type not found")
	}

	if userTypes.TypeName == nil {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("User type not found")
	}

	return &userTypes.ID, nil
}

func (utr *UserTypeRepository) CreateUserType(tx *sql.Tx, type_name *string) (*uuid.UUID, *models.INVError) {
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
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating user type")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating user type")
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("User type already exists")
	}

	return &uuid, nil
}

func (utr *UserTypeRepository) UpdateUserType(tx *sql.Tx, userType *model.UserTypes) *models.INVError {
	// Create the update statement
	updateQuery := table.UserTypes.UPDATE(
		table.UserTypes.TypeName,
	).SET(
		userType.TypeName,
	).WHERE(table.UserTypes.ID.EQ(mysql.String(userType.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating user type")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating user type")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("User type not found")
	}

	return nil
}

func (utr *UserTypeRepository) DeleteUserType(tx *sql.Tx, userTypeId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.UserTypes.DELETE().WHERE(
		table.UserTypes.ID.EQ(mysql.String(userTypeId.String())),
	)

	// Execute the query
	rows, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting user type")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting user type")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("User type not found")
	}

	return nil
}
