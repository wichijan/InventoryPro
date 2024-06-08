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
	RegisterUser(registrationData models.RegistrationRequest) (*models.LoginResponse, *models.INVError)
	LoginUser(loginData models.LoginRequest) (*models.LoginResponse, *models.INVError)
	CheckEmail(email string) *models.INVError
	CheckUsername(username string) *models.INVError
	GetUserById(userId *uuid.UUID) (*models.UserWithTypeName, *models.INVError)
}

type UserController struct {
	UserRepo     repositories.UserRepositoryI
	UserTypeRepo repositories.UserTypeRepositoryI
}

func (uc *UserController) RegisterUser(registrationData models.RegistrationRequest) (*models.LoginResponse, *models.INVError) {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	userId := uuid.New()

	hash, err := utils.HashPassword(registrationData.Password)
	if err != nil {
		return nil, inv_errors.INV_UPSTREAM_ERROR
	}

	inv_err := uc.UserRepo.CheckIfEmailExists(registrationData.Email)
	if inv_err != nil {
		return nil, inv_err
	}

	var userTypeId *string
	userTypeId = nil
	if registrationData.UserTypeName != "" {
		userTypeId, inv_err = uc.UserTypeRepo.GetUserTypeByName(&registrationData.UserTypeName)
		if inv_err != nil {
			return nil, inv_err
		}
	}

	registrationDate := time.Now()
	registrationDateString := registrationDate.String()
	isTrue := true

	userForResponse := models.UserWithTypeName{
		ID:               userId.String(),
		Username:         &registrationData.Username,
		Email:            &registrationData.Email,
		Password:         &hash,
		FirstName:        &registrationData.FirstName,
		LastName:         &registrationData.LastName,
		JobTitle:         &registrationData.JobTitle,
		PhoneNumber:      &registrationData.PhoneNumber,
		UserTypeName:     &registrationData.UserTypeName,
		RegistrationTime: &registrationDateString,
		IsActive:         &isTrue,
	}

	user := model.Users{
		ID:               userId.String(),
		Username:         &registrationData.Username,
		Email:            &registrationData.Email,
		Password:         &hash,
		FirstName:        &registrationData.FirstName,
		LastName:         &registrationData.LastName,
		JobTitle:         &registrationData.JobTitle,
		PhoneNumber:      &registrationData.PhoneNumber,
		UserTypeID:       userTypeId,
		RegistrationTime: &registrationDate,
		IsActive:         &isTrue,
		// TODO RegistrationAccepted
	}

	inv_err = uc.UserRepo.CreateUser(tx, user)
	if inv_err != nil {
		return nil, inv_err
	}

	userUUID, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	token, refreshToken, err := utils.GenerateJWT(&userUUID)
	if err != nil {
		return nil, inv_errors.INV_UPSTREAM_ERROR
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &models.LoginResponse{
		User:         userForResponse,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *UserController) LoginUser(loginData models.LoginRequest) (*models.LoginResponse, *models.INVError) {
	// get user from database
	user, inv_err := uc.UserRepo.GetUserByUsername(loginData.Username)
	if inv_err != nil {
		return nil, inv_err
	}

	// check if password is correct
	if ok := utils.ComparePasswordHash(loginData.Password, *user.Password); !ok {
		return nil, inv_errors.INV_CREDENTIALS_INVALID
	}

	// Convert string to UUID
	userID, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	// generate JWT token
	token, refreshToken, err := utils.GenerateJWT(&userID)
	if err != nil {
		return nil, inv_errors.INV_UPSTREAM_ERROR
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
