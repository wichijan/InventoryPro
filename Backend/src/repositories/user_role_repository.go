package repositories

import (
	"database/sql"

	"github.com/go-jet/jet/v2/mysql"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type UserRoleRepositoryI interface {
	CreateUserRole(tx *sql.Tx, user_role *model.UserRoles) *models.INVError
	DeleteUserRole(tx *sql.Tx, userRole *model.UserRoles) *models.INVError

	managers.DatabaseManagerI
}

type UserRoleRepository struct {
	managers.DatabaseManagerI
}

func (urr *UserRoleRepository) CreateUserRole(tx *sql.Tx, user_role *model.UserRoles) *models.INVError {
	// Create the insert statement
	insertQuery := table.UserRoles.INSERT(
		table.UserRoles.RoleID,
		table.UserRoles.UserID,
	).VALUES(
		user_role.RoleID,
		user_role.UserID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating user role")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating user role")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Error creating user role")
	}

	return nil
}

func (urr *UserRoleRepository) DeleteUserRole(tx *sql.Tx, userRole *model.UserRoles) *models.INVError {
	// Create the delete statement
	deleteQuery := table.UserRoles.DELETE().
		WHERE(
			table.UserRoles.UserID.EQ(mysql.String(userRole.UserID)).
				AND(table.UserRoles.RoleID.EQ(mysql.String(userRole.RoleID))),
		)

	// Execute the query
	rows, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting user role")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting user role")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Error deleting user role")
	}

	return nil
}
