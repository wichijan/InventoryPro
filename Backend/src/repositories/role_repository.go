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

type RoleRepositoryI interface {
	GetRoles() (*[]model.Roles, *models.INVError)
	CreateRole(roleName *string) (*uuid.UUID, *models.INVError)
	UpdateRole(role *model.Roles) *models.INVError
	DeleteRole(roleId *uuid.UUID) *models.INVError
}

type RoleRepository struct {
	DatabaseManager managers.DatabaseManagerI
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
	err := stmt.Query(rr.DatabaseManager.GetDatabaseConnection(), &roles)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(roles) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &roles, nil
}

func (rr *RoleRepository) CreateRole(roleName *string) (*uuid.UUID, *models.INVError) {
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
	rows, err := insertQuery.Exec(rr.DatabaseManager.GetDatabaseConnection())
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

func (rr *RoleRepository) UpdateRole(role *model.Roles) *models.INVError {
	// Create the update statement
	updateQuery := table.Roles.UPDATE(
		table.Roles.RoleName,
	).SET(
		role.RoleName,
	).WHERE(table.Roles.ID.EQ(mysql.String(role.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(rr.DatabaseManager.GetDatabaseConnection())
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

func (rr *RoleRepository) DeleteRole(roleId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
