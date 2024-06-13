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

type RoleRepositoryI interface {
	GetRoles() (*[]model.Roles, *models.INVError)
	CreateRole(tx *sql.Tx, roleName *string) (*uuid.UUID, *models.INVError)
	UpdateRole(tx *sql.Tx, role *model.Roles) *models.INVError
	DeleteRole(tx *sql.Tx, roleId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type RoleRepository struct {
	managers.DatabaseManagerI
}

func (rr *RoleRepository) GetRoles() (*[]model.Roles, *models.INVError) {
	var roles []model.Roles

	// Create the query
	stmt := mysql.SELECT(
		table.Roles.AllColumns,
	).FROM(
		table.Roles,
	)

	// Execute the query
	err := stmt.Query(rr.GetDatabaseConnection(), &roles)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading roles")
	}

	return &roles, nil
}

func (rr *RoleRepository) CreateRole(tx *sql.Tx, roleName *string) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Roles.INSERT(
		table.Roles.ID,
		table.Roles.RoleName,
	).VALUES(
		uuid.String(),
		roleName,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating role")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating role")
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Role already exists")
	}

	return &uuid, nil
}

func (rr *RoleRepository) UpdateRole(tx *sql.Tx, role *model.Roles) *models.INVError {
	// Create the update statement
	updateQuery := table.Roles.UPDATE(
		table.Roles.RoleName,
	).SET(
		role.RoleName,
	).WHERE(table.Roles.ID.EQ(mysql.String(role.ID)))

	// Execute the query
	_, err := updateQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating role")
	}

	return nil
}

func (rr *RoleRepository) DeleteRole(tx *sql.Tx, roleId *uuid.UUID) *models.INVError {
	// Create the delete statement
	deleteQuery := table.Roles.DELETE().WHERE(table.Roles.ID.EQ(mysql.String(roleId.String())))

	// Execute the query
	_, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting role")
	}

	return nil
}
