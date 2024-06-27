package models

type RegistrationCodeResponse struct {
	RegistrationCode string
}

type RegistrationCodes struct {
	RegistrationCode string `alias:"registration_codes.code"`
	User             struct {
		UserWithoutRoles
	}
}
