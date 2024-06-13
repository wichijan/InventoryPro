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
	GetUserByNameClean(username *string) (*model.Users, *models.INVError)
	GetUserByUsername(username string) (*models.UserWithTypeName, *models.INVError)
	CreateUser(tx *sql.Tx, user model.Users) *models.INVError
	UpdateUser(tx *sql.Tx, user *model.Users) *models.INVError
	DeleteUser(tx *sql.Tx, userId *uuid.UUID) *models.INVError
	CheckIfUsernameExists(username string) *models.INVError
	CheckIfEmailExists(email string) *models.INVError

	StoreUserPicture(tx *sql.Tx, userId *uuid.UUID) (*uuid.UUID, *models.INVError)
	GetPictureIdFromUser(userId *uuid.UUID) (*uuid.UUID, *models.INVError)
	RemovePictureIdFromUser(tx *sql.Tx, userId *uuid.UUID) *models.INVError

	AcceptUserRegistrationRequest(tx *sql.Tx, userId *string) *models.INVError

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
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_USER_NOT_FOUND.WithDetails("User not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading user")
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByNameClean(username *string) (*model.Users, *models.INVError) {
	var user model.Users
	stmt := mysql.SELECT(
		table.Users.ID,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.JobTitle,
		table.Users.PhoneNumber,
		table.Users.ProfilePicture,
		table.Users.RegistrationTime,
		table.Users.RegistrationAccepted,
		table.Users.IsActive,
	).FROM(
		table.Users,
	).WHERE(
		table.Users.Username.EQ(utils.MySqlString(*username)),
	)
	err := stmt.Query(ur.GetDatabaseConnection(), &user)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_USER_NOT_FOUND.WithDetails("User not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading user")
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
			return nil, inv_errors.INV_USER_NOT_FOUND.WithDetails("User not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading user")
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating user")
	}

	return nil
}

func (ur *UserRepository) UpdateUser(tx *sql.Tx, user *model.Users) *models.INVError {
	stmt := table.Users.UPDATE(
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
	).SET(
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
	).WHERE(
		table.Users.ID.EQ(mysql.String(user.ID)),
	)

	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating user")
	}

	return nil
}

func (ur *UserRepository) DeleteUser(tx *sql.Tx, userId *uuid.UUID) *models.INVError {
	stmt := table.Users.DELETE().WHERE(
		table.Users.ID.EQ(mysql.String(userId.String())),
	)

	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting user")
	}

	return nil
}

func (ur *UserRepository) CheckIfUsernameExists(username string) *models.INVError {
	count, err := utils.CountStatement(table.Users, table.Users.Username.EQ(mysql.String(username)), ur.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if username exists")
	}
	if count > 0 {
		return inv_errors.INV_USERNAME_EXISTS.WithDetails("Username already exists")
	}
	return nil
}

func (ur *UserRepository) CheckIfEmailExists(email string) *models.INVError {
	count, err := utils.CountStatement(table.Users, table.Users.Email.EQ(mysql.String(email)), ur.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if email exists")
	}
	if count > 0 {
		return inv_errors.INV_EMAIL_EXISTS.WithDetails("Email already exists")
	}
	return nil
}

func (ur *UserRepository) AcceptUserRegistrationRequest(tx *sql.Tx, userId *string) *models.INVError {
	stmt := table.Users.UPDATE(
		table.Users.RegistrationAccepted,
		table.Users.IsActive,
	).SET(
		mysql.Bool(true),
		mysql.Bool(true),
	).WHERE(
		table.Users.ID.EQ(mysql.String(*userId)),
	)
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error accepting user registration request")
	}
	return nil
}

func (ur *UserRepository) StoreUserPicture(tx *sql.Tx, userId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	updatePictureQuery := table.Users.UPDATE(
		table.Users.ProfilePicture,
	).SET(
		uuid.String(),
	).WHERE(table.Users.ID.EQ(mysql.String(userId.String())))

	// Execute the query
	rows, err := updatePictureQuery.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error storing picture for user")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error storing picture for user")
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("User not found")
	}

	return &uuid, nil
}

func (ur *UserRepository) GetPictureIdFromUser(userId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	var picture models.UserPicture

	// Create the query
	stmt := mysql.SELECT(
		table.Users.ProfilePicture,
	).FROM(
		table.Users,
	).WHERE(
		table.Users.ID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	err := stmt.Query(ur.GetDatabaseConnection(), &picture)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading picture for user")
	}

	if picture.PictureId == "" {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("User not found")
	}

	pictureId, err := uuid.Parse(picture.PictureId)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error parsing picture id")
	}

	return &pictureId, nil
}

func (ur *UserRepository) RemovePictureIdFromUser(tx *sql.Tx, userId *uuid.UUID) *models.INVError {
	// Create the update statement
	updatePictureQuery := table.Users.UPDATE(
		table.Users.ProfilePicture,
	).SET(
		"",
	).WHERE(table.Users.ID.EQ(mysql.String(userId.String())))

	// Execute the query
	rows, err := updatePictureQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error removing picture for User")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error removing picture for User")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("User not found")
	}

	return nil
}