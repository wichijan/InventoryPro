package controllers

import (
	"time"

	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
	"github.com/wichijan/InventoryPro/src/utils"
)

type UserControllerI interface {
	RegisterUser(registrationData models.RegistrationRequest) *models.INVError
	LoginUser(loginData models.LoginRequest) (*models.LoginResponse, *models.INVError)
	CheckEmail(email string) *models.INVError
	CheckUsername(username string) *models.INVError
	GetUserById(userId *uuid.UUID) (*models.UserWithTypeName, *models.INVError)

	AcceptUserRegistrationRequest(userId *string) *models.INVError
	GetRegistrationRequests() (*[]model.RegistrationRequests, *models.INVError)
}

type UserController struct {
	UserRepo                repositories.UserRepositoryI
	UserTypeRepo            repositories.UserTypeRepositoryI
	RegistrationRequestRepo repositories.RegistrationRequestRepositoryI
}

func (uc *UserController) RegisterUser(registrationData models.RegistrationRequest) *models.INVError {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	userId := uuid.New()

	hash, err := utils.HashPassword(registrationData.Password)
	if err != nil {
		return inv_errors.INV_UPSTREAM_ERROR.WithDetails("Invalid password")
	}

	inv_err := uc.UserRepo.CheckIfEmailExists(registrationData.Email)
	if inv_err != nil {
		return inv_err
	}

	var userTypeId *string
	userTypeId = nil
	if registrationData.UserTypeName != "" {
		userTypeId, inv_err = uc.UserTypeRepo.GetUserTypeByName(&registrationData.UserTypeName)
		if inv_err != nil {
			return inv_err
		}
	}

	registrationDate := time.Now()
	isFalse := false

	user := model.Users{
		ID:                   userId.String(),
		Username:             &registrationData.Username,
		Email:                &registrationData.Email,
		Password:             &hash,
		FirstName:            &registrationData.FirstName,
		LastName:             &registrationData.LastName,
		JobTitle:             &registrationData.JobTitle,
		PhoneNumber:          &registrationData.PhoneNumber,
		UserTypeID:           userTypeId,
		RegistrationTime:     &registrationDate,
		RegistrationAccepted: &isFalse,
		IsActive:             &isFalse,
	}

	inv_err = uc.UserRepo.CreateUser(tx, user)
	if inv_err != nil {
		return inv_err
	}

	// Create registration request
	inv_error := uc.RegistrationRequestRepo.CreateRequest(tx, &model.RegistrationRequests{
		UserID:      userId.String(),
		RequestTime: &registrationDate,
	})
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (uc *UserController) LoginUser(loginData models.LoginRequest) (*models.LoginResponse, *models.INVError) {
	// get user from database
	user, inv_err := uc.UserRepo.GetUserByUsername(loginData.Username)
	if inv_err != nil {
		return nil, inv_err
	}

	// check if password is correct
	if ok := utils.ComparePasswordHash(loginData.Password, *user.Password); !ok {
		return nil, inv_errors.INV_CREDENTIALS_INVALID.WithDetails("Invalid username or password")
	}

	// Check if user registration request has been accepted
	if !*user.RegistrationAccepted {
		return nil, inv_errors.INV_USER_NOT_ACCEPTED.WithDetails("User registration request has not been accepted")
	}

	// Convert string to UUID
	userID, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error parsing user ID")
	}

	// generate JWT token
	token, refreshToken, err := utils.GenerateJWT(&userID)
	if err != nil {
		return nil, inv_errors.INV_UPSTREAM_ERROR.WithDetails("Error generating JWT token")
	}

	return &models.LoginResponse{
		User:         *user,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *UserController) CheckEmail(email string) *models.INVError {
	return uc.UserRepo.CheckIfEmailExists(email)
}

func (uc *UserController) CheckUsername(username string) *models.INVError {
	return uc.UserRepo.CheckIfUsernameExists(username)
}

func (uc *UserController) GetUserById(userId *uuid.UUID) (*models.UserWithTypeName, *models.INVError) {
	return uc.UserRepo.GetUserById(userId)
}

func (uc *UserController) AcceptUserRegistrationRequest(userIdString *string) *models.INVError {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	userId, inv_error := uuid.Parse(*userIdString)
	if inv_error != nil {
		return inv_errors.INV_BAD_REQUEST.WithDetails("Invalid user ID")
	}

	// Check if user registration request exists
	_, inv_err := uc.RegistrationRequestRepo.GetRequestByUserId(&userId)
	if inv_err != nil {
		return inv_err
	}

	// Accept user registration request
	inv_err = uc.UserRepo.AcceptUserRegistrationRequest(tx, userIdString)
	if inv_err != nil {
		return inv_err
	}

	// Delete registration request
	inv_err = uc.RegistrationRequestRepo.DeleteRequest(tx, &userId)
	if inv_err != nil {
		return inv_err
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (uc *UserController) GetRegistrationRequests() (*[]model.RegistrationRequests, *models.INVError) {
	return uc.RegistrationRequestRepo.GetRegistrationRequests()
}

