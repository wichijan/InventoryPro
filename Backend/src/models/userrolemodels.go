package models

type UserRoleWithName struct {
	UserID string `alias:"user_roles.user_id" sql:"primary_key"`

	UserNames []struct {
		RoleName string `alias:"roles.role_name"`
	}
}

type RoleIdODT struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}
