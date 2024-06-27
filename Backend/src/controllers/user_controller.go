package controllers

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/managers"
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
	GetUsers() (*[]models.Users, *models.INVError)
	IsAdmin(userId *uuid.UUID) (bool, *models.INVError)

	UpdateUser(user models.UserWithoutRoles) *models.INVError

	AcceptUserRegistrationRequest(userId *string) *models.INVError
	DeclineUserRegistrationRequest(userId *string) *models.INVError
	GetRegistrationRequests() (*[]model.RegistrationRequests, *models.INVError)

	ValidateRegistrationCode(code *string) (*uuid.UUID, *models.INVError)
	RegisterUserAndCode(registrationData models.RegistrationRequest) (*models.RegistrationCodeResponse, *models.INVError)
	DeleteRegistrationCode(code *string) *models.INVError

	UpdateUserPassword(userId *uuid.UUID, password string) *models.INVError
	ForgotPassword(username string) *models.INVError

	SendEmailToAdmins(message string) *models.INVError

	UploadUserImage(itemId *uuid.UUID) (*uuid.UUID, *models.INVError)
	GetImageIdFromUser(userId *uuid.UUID) (*uuid.UUID, *models.INVError)
	RemoveImageIdFromUser(userId *uuid.UUID) *models.INVError
}

type UserController struct {
	UserRepo                repositories.UserRepositoryI
	UserTypeRepo            repositories.UserTypeRepositoryI
	RegistrationRequestRepo repositories.RegistrationRequestRepositoryI
	RegistrationCodeRepo    repositories.RegistrationCodeRepositoryI
	RoleRepo                repositories.RoleRepositoryI
	MailMgr                 managers.MailMgr
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

	if !*user.IsActive {
		return nil, inv_errors.INV_CREDENTIALS_INVALID.WithDetails("User not active")
	}

	// check if password is correct
	if ok := utils.ComparePasswordHash(loginData.Password, *user.Password); !ok {
		return nil, inv_errors.INV_CREDENTIALS_INVALID.WithDetails("Invalid username or password")
	}

	// Check if user registration request has been accepted
	if !*user.RegistrationAccepted {
		return nil, inv_errors.INV_CONFLICT.WithDetails("User registration request has not been accepted")
	}
	if !*user.IsActive && *user.Password == "" {
		return nil, inv_errors.INV_CONFLICT.WithDetails("User has not being accepted by using the registration code")
	}
	if !*user.IsActive {
		return nil, inv_errors.INV_CONFLICT.WithDetails("User is not (yet) active")
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

func (uc *UserController) GetUsers() (*[]models.Users, *models.INVError) {
	return uc.UserRepo.GetUsers()
}

func (uc *UserController) CheckUsername(username string) *models.INVError {
	return uc.UserRepo.CheckIfUsernameExists(username)
}

func (uc *UserController) GetUserById(userId *uuid.UUID) (*models.UserWithTypeName, *models.INVError) {
	return uc.UserRepo.GetUserById(userId)
}

func (uc *UserController) UpdateUser(user models.UserWithoutRoles) *models.INVError {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	userTypeId, inv_err := uc.UserTypeRepo.GetUserTypeByName(user.UserTypeName)
	if inv_err != nil {
		return inv_err
	}

	userPure, inv_err := uc.UserRepo.GetUserPureById(user.ID)
	if inv_err != nil {
		return inv_err
	}

	userPure.Email = &user.Email
	userPure.FirstName = &user.FirstName
	userPure.LastName = &user.LastName
	userPure.JobTitle = &user.JobTitle
	userPure.PhoneNumber = &user.PhoneNumber
	userPure.UserTypeID = userTypeId

	inv_error := uc.UserRepo.UpdateUser(tx, userPure)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
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

	user, inv_err := uc.UserRepo.GetUserPureById(&userId)
	if inv_err != nil {
		return inv_err
	}

	// Check if user registration request exists
	_, inv_err = uc.RegistrationRequestRepo.GetRequestByUserId(&userId)
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

	uc.MailMgr.SendWelcomeMail(*user.Email, *user.Username)

	return nil
}

func (uc *UserController) DeclineUserRegistrationRequest(userIdString *string) *models.INVError {
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

	// Delete registration request
	inv_err = uc.RegistrationRequestRepo.DeleteRequest(tx, &userId)
	if inv_err != nil {
		return inv_err
	}

	// Decline user registration request => delete user
	inv_err = uc.UserRepo.DeleteUser(tx, &userId)
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

func (uc *UserController) ValidateRegistrationCode(code *string) (*uuid.UUID, *models.INVError) {
	ok, inv_err := uc.RegistrationCodeRepo.CheckIfUserWithCodeExists(code)
	if inv_err != nil {
		return nil, inv_err
	}

	if !*ok {
		return nil, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid registration code")
	}

	user, inv_err := uc.RegistrationCodeRepo.GetUserIdByCode(code)
	if inv_err != nil {
		return nil, inv_err
	}

	userId, err := uuid.Parse(user.UserID)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error parsing user ID")
	}

	return &userId, nil
}

func (uc *UserController) RegisterUserAndCode(registrationData models.RegistrationRequest) (*models.RegistrationCodeResponse, *models.INVError) {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	userId := uuid.New()

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
	isTrue := true
	isFalse := false

	user := model.Users{
		ID:                   userId.String(),
		Username:             &registrationData.Username,
		Email:                &registrationData.Email,
		FirstName:            &registrationData.FirstName,
		LastName:             &registrationData.LastName,
		JobTitle:             &registrationData.JobTitle,
		PhoneNumber:          &registrationData.PhoneNumber,
		UserTypeID:           userTypeId,
		RegistrationTime:     &registrationDate,
		RegistrationAccepted: &isTrue,
		IsActive:             &isFalse,
	}

	inv_err = uc.UserRepo.CreateUser(tx, user)
	if inv_err != nil {
		return nil, inv_err
	}

	code := utils.GenerateRandomString(20)

	// Insert Code into table RegistrationCodes
	inv_err = uc.RegistrationCodeRepo.CreateRegistrationCode(tx, &model.RegistrationCodes{
		UserID: userId.String(),
		Code:   &code,
	})
	if inv_err != nil {
		return nil, inv_err
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	inv_err = uc.MailMgr.SendRegistrationCodeMail(registrationData.Email, code)
	if inv_err != nil {
		return nil, inv_err
	}

	return &models.RegistrationCodeResponse{
		RegistrationCode: code,
	}, nil
}

func (uc *UserController) UpdateUserPassword(userId *uuid.UUID, password string) *models.INVError {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	hash, err := utils.HashPassword(password)
	if err != nil {
		return inv_errors.INV_UPSTREAM_ERROR.WithDetails("Invalid password")
	}

	user, inv_err := uc.UserRepo.GetUserPureById(userId)
	if inv_err != nil {
		return inv_err
	}

	if !*user.IsActive && user.Password == nil {
		var isTrue bool = true
		user.IsActive = &isTrue
	}
	user.Password = &hash

	inv_error := uc.UserRepo.UpdateUser(tx, user)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return nil
}

func (uc *UserController) DeleteRegistrationCode(code *string) *models.INVError {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	inv_err := uc.RegistrationCodeRepo.DeleteRegistrationCode(tx, code)
	if inv_err != nil {
		return inv_err
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return nil
}

func (uc *UserController) UploadUserImage(userId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	pictureId, inv_error := uc.UserRepo.StoreUserPicture(tx, userId)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return pictureId, nil
}

func (uc *UserController) GetImageIdFromUser(userId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	pictureId, inv_error := uc.UserRepo.GetPictureIdFromUser(userId)
	if inv_error != nil {
		return nil, inv_error
	}

	return pictureId, nil
}

func (uc *UserController) RemoveImageIdFromUser(userId *uuid.UUID) *models.INVError {
	tx, err := uc.UserRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	pictureId, inv_error := uc.UserRepo.GetPictureIdFromUser(userId)
	if inv_error != nil {
		return inv_error
	}

	inv_error = uc.UserRepo.RemovePictureIdFromUser(tx, userId)
	if inv_error != nil {
		return inv_error
	}

	imageName := "./../uploads/" + pictureId.String() + ".jpeg"
	inv_err := os.Remove(imageName)
	if inv_err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error removing image")
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (uc *UserController) IsAdmin(userId *uuid.UUID) (bool, *models.INVError) {
	userRoles, inv_err := uc.RoleRepo.GetRolesForUser(userId)
	if inv_err != nil {
		return false, inv_err
	}

	for _, a := range userRoles.UserNames {
		if a.RoleName == "Admin" {
			return true, nil
		}
	}

	return false, nil
}

func (uc *UserController) ForgotPassword(username string) *models.INVError {
	user, inv_err := uc.UserRepo.GetUserByUsername(username)
	if inv_err != nil {
		return inv_err
	}

	if !*user.IsActive {
		return inv_errors.INV_CREDENTIALS_INVALID.WithDetails("User not active")
	}

	// send new password via email
	inv_err = uc.MailMgr.SendLinkForNewPasswordMail(*user.Email, &user.ID)
	if inv_err != nil {
		return inv_err
	}

	return nil
}

func (uc *UserController) SendEmailToAdmins(userRegistUsername string) *models.INVError {
	admins, inv_err := uc.UserRepo.GetAdmins()
	if inv_err != nil {
		return inv_err
	}
	log.Print("Admins: ", admins)

	for _, admin := range *admins {
		inv_err = uc.MailMgr.SendEmailToAdmin(*admin.Email, userRegistUsername)
		if inv_err != nil {
			return inv_err
		}
	}

	return nil
}
