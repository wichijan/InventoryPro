package models

import (
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
)

type RegistrationRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	JobTitle     string `json:"jobtitle"`
	PhoneNumber  string `json:"phonenumber"`
	UserTypeName string `json:"usertypename"`
}

type UserWithTypeName struct {
	ID                   string  `alias:"users.id" sql:"primary_key"`
	FirstName            *string `alias:"users.first_name"`
	LastName             *string `alias:"users.last_name"`
	Username             *string `alias:"users.username"`
	Email                *string `alias:"users.email"`
	Password             *string `alias:"users.password"`
	JobTitle             *string `alias:"users.job_title"`
	PhoneNumber          *string `alias:"users.phone_number"`
	UserTypeName         *string `alias:"user_types.type_name"`
	ProfilePicture       *string `alias:"users.profile_picture"`
	RegistrationTime     *string `alias:"users.registration_time"`
	RegistrationAccepted *bool   `alias:"users.registration_accepted"`
	IsActive             *bool   `alias:"users.is_active"`

	Roles []struct {
		model.Roles
	}
}

type Users struct {
	ID             string  `alias:"users.id" sql:"primary_key"`
	Username       string  `alias:"users.username"`
	Email          string  `alias:"users.email"`
	FirstName      string  `alias:"users.first_name"`
	LastName       string  `alias:"users.last_name"`
	JobTitle       string  `alias:"users.job_title"`
	PhoneNumber    string  `alias:"users.phone_number"`
	UserTypeName   *string `alias:"user_types.type_name"`
	ProfilePicture *string `alias:"users.profile_picture"`
	IsActive       bool    `alias:"users.is_active"`

	Roles []struct {
		model.Roles
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User         UserWithTypeName
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

type PasswordReset struct {
	Password string
}

type Username struct {
	Username string
}

type PasswordResetEmail struct {
	UserId   *uuid.UUID
	Password string
}

type RegistrationCode struct {
	Password string `json:"password" binding:"required"`
}

type UserPicture struct {
	PictureId string `alias:"users.profile_picture"`
}
