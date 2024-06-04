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

type UserRepositoryI interface {
	GetUserById(id *uuid.UUID) (*models.UserWithTypeName, *models.INVError)
	GetUserByUsername(username string) (*models.UserWithTypeName, *models.INVError)
	CreateUser(tx *sql.Tx, user model.Users) *models.INVError
	CheckIfUsernameExists(username string) *models.INVError
	CheckIfEmailExists(email string) *models.INVError

	// TODO implement upload profile picture

	managers.DatabaseManagerI
}

type UserRepository struct {
	managers.DatabaseManagerI
}

func (ur *UserRepository) GetUserById(id *uuid.UUID) (*models.UserWithTypeName, *models.INVError) {
	var user models.UserWithTypeName
	stmt := mysql.SELECT(
		table.Users.ID,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.JobTitle,
		table.Users.PhoneNumber,
		table.UserTypes.TypeName,

		table.Users.ProfilePicture,
		table.Users.RegistrationTime,
		table.Users.RegistrationAccepted,
		table.Users.IsActive,
	).FROM(
		table.Users.
			LEFT_JOIN(table.UserTypes, table.UserTypes.ID.EQ(table.Users.UserTypeID)),
	).WHERE(
		table.Users.ID.EQ(utils.MySqlString(id.String())),
	)
	err := stmt.Query(ur.GetDatabaseConnection(), &user)
	if err != nil {
		if err.Error() == "jet: sql: no rows in result set" {
			return nil, inv_errors.INV_USER_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.UserWithTypeName, *models.INVError) {
	var user models.UserWithTypeName
	stmt := mysql.SELECT(
		table.Users.ID,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.JobTitle,
		table.Users.PhoneNumber,
		table.UserTypes.TypeName,

		table.Users.ProfilePicture,
		table.Users.RegistrationTime,
		table.Users.RegistrationAccepted,
		table.Users.IsActive,
	).FROM(
		table.Users.
			LEFT_JOIN(table.UserTypes, table.UserTypes.ID.EQ(table.Users.UserTypeID)),
	).WHERE(
		table.Users.Username.EQ(mysql.String(username)),
	)
	err := stmt.Query(ur.GetDatabaseConnection(), &user)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_USER_NOT_FOUND
		}
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &user, nil
}

func (ur *UserRepository) CreateUser(tx *sql.Tx, user model.Users) *models.INVError {
	stmt := table.Users.INSERT(
		table.Users.ID,
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.JobTitle,
		table.Users.PhoneNumber,
		table.Users.UserTypeID,

		table.Users.ProfilePicture,
		table.Users.RegistrationTime,
		table.Users.RegistrationAccepted,
		table.Users.IsActive,
	).VALUES(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
		user.JobTitle,
		user.PhoneNumber,
		user.UserTypeID,
		user.ProfilePicture,
		user.RegistrationTime,
		user.RegistrationAccepted,
		user.IsActive,
	)

	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (ur *UserRepository) CheckIfUsernameExists(username string) *models.INVError {
	count, err := utils.CountStatement(table.Users, table.Users.Username.EQ(mysql.String(username)), ur.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	if count > 0 {
		return inv_errors.INV_USERNAME_EXISTS
	}
	return nil
}

func (ur *UserRepository) CheckIfEmailExists(email string) *models.INVError {
	count, err := utils.CountStatement(table.Users, table.Users.Email.EQ(mysql.String(email)), ur.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	if count > 0 {
		return inv_errors.INV_EMAIL_EXISTS
	}
	return nil
}
