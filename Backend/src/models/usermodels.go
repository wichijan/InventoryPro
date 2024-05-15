package models

import (
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
)

type RegistrationRequest struct {
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	FirstName   string  `json:"firstname"`
	LastName    string  `json:"lastname"`
	JobTitle    string  `json:"jobtitle"`
	PhoneNumber string  `json:"phonenumber"`
	UserTypeID  *string `json:"usertypeid"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User         model.Users
	Token        string
	RefreshToken string
}

type CheckEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}

type LoggedInResponse struct {
	LoggedIn bool       `json:"loggedIn"`
	Id       *uuid.UUID `json:"id"`
}
