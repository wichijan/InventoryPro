package models

type UserRoleWithName struct {
	UserID string `alias:"user_roles.user_id" sql:"primary_key"`

	UserNames []struct {
		RoleName string `alias:"roles.role_name"`
	}
}

type RoleIdODT struct {
	UserID string `binding:"required"`
	RoleID string `binding:"required"`
}
