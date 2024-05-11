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

type UserRoleRepositoryI interface {
	GetUserRoles() (*[]model.UserRoles, *models.INVError)
	CreateUserRole(user_role *model.UserRoles) (*uuid.UUID, *models.INVError)
	UpdateUserRole(userRole *model.UserRoles) *models.INVError
	DeleteUserRole(userRoleId *uuid.UUID) *models.INVError
}

type UserRoleRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (urr *UserRoleRepository) GetUserRoles() (*[]model.UserRoles, *models.INVError) {
	var userRoles []model.UserRoles

	// Create the query
	stmt := mysql.SELECT(
		table.UserRoles.AllColumns,
	).FROM(
		table.UserRoles,
	)

	// Execute the query
	err := stmt.Query(urr.DatabaseManager.GetDatabaseConnection(), &userRoles)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(userRoles) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &userRoles, nil
}

func (urr *UserRoleRepository) CreateUserRole(user_role *model.UserRoles) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.UserRoles.INSERT(
		table.UserRoles.ID,
		table.UserRoles.RoleID,
		table.UserRoles.UserID,
	).VALUES(
		uuid.String(),
		user_role.RoleID,
		user_role.UserID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(urr.DatabaseManager.GetDatabaseConnection())
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

func (urr *UserRoleRepository) UpdateUserRole(userRole *model.UserRoles) *models.INVError {
	// Create the update statement
	updateQuery := table.UserRoles.UPDATE(
		table.UserRoles.RoleID,
		table.UserRoles.UserID,
	).SET(
		userRole.RoleID,
		userRole.UserID,
	).WHERE(table.UserRoles.ID.EQ(mysql.String(userRole.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(urr.DatabaseManager.GetDatabaseConnection())
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

func (urr *UserRoleRepository) DeleteUserRole(userRoleId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
