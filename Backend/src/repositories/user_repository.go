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

type UserRepositoryI interface {
	GetUsers() (*[]model.Users, *models.INVError)
	CreateUser(user *model.Users) (*uuid.UUID, *models.INVError)
	UpdateUser(user *model.Users) *models.INVError
	DeleteUser(userId *uuid.UUID) *models.INVError
}

type UserRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (ur *UserRepository) GetUsers() (*[]model.Users, *models.INVError) {
	var users []model.Users

	// Create the query
	stmt := mysql.SELECT(
		table.Users.AllColumns,
	).FROM(
		table.Users,
	)

	// Execute the query
	err := stmt.Query(ur.DatabaseManager.GetDatabaseConnection(), &users)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(users) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &users, nil
}

func (ur *UserRepository) CreateUser(user *model.Users) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Users.INSERT(
		table.Users.ID,
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.JobTitle,
		table.Users.PhoneNumber,
		table.Users.UserTypeID,
	).VALUES(
		uuid.String(),
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
		user.JobTitle,
		user.PhoneNumber,
		user.UserTypeID,		
	)

	// Execute the query
	rows, err := insertQuery.Exec(ur.DatabaseManager.GetDatabaseConnection())
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

func (ur *UserRepository) UpdateUser(user *model.Users) *models.INVError {
	// Create the update statement
	updateQuery := table.Users.UPDATE(
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.JobTitle,
		table.Users.PhoneNumber,
		table.Users.UserTypeID,
	).SET(
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
		user.JobTitle,
		user.PhoneNumber,
		user.UserTypeID,		
	).WHERE(table.Users.ID.EQ(mysql.String(user.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(ur.DatabaseManager.GetDatabaseConnection())
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

func (ur *UserRepository) DeleteUser(userId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
