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
	CreateUserRole(user_role *model.UserRoles) *models.INVError
	UpdateUserRole(userRole *model.UserRoles) *models.INVError
	DeleteUserRole(userRoleId *uuid.UUID) *models.INVError
}

type UserRoleRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (urr *UserRoleRepository) CreateUserRole(user_role *model.UserRoles) *models.INVError {

	// Create the insert statement
	insertQuery := table.UserRoles.INSERT(
		table.UserRoles.RoleID,
		table.UserRoles.UserID,
	).VALUES(
		user_role.RoleID,
		user_role.UserID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(urr.DatabaseManager.GetDatabaseConnection())
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

func (urr *UserRoleRepository) UpdateUserRole(userRole *model.UserRoles) *models.INVError {
	// Create the update statement
	updateQuery := table.UserRoles.UPDATE(
		table.UserRoles.RoleID,
		table.UserRoles.UserID,
	).SET(
		userRole.RoleID,
		userRole.UserID,
	).WHERE(table.UserRoles.UserID.EQ(mysql.String(userRole.UserID)).AND(table.UserRoles.RoleID.EQ(mysql.String(userRole.RoleID))))

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
