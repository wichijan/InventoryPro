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

type RegistrationCodeRepositoryI interface {
	GetRegistrationCodeById(userId *uuid.UUID) (*model.RegistrationCodes, *models.INVError)
	GetUserIdByCode(code *string) (*model.RegistrationCodes, *models.INVError)
	CreateRegistrationCode(tx *sql.Tx, registrationCode *model.RegistrationCodes) *models.INVError
	DeleteRegistrationCode(tx *sql.Tx, code *string) *models.INVError

	CheckIfUserWithCodeExists(code *string) (*bool, *models.INVError)

	managers.DatabaseManagerI
}

type RegistrationCodeRepository struct {
	managers.DatabaseManagerI
}

func (rcr *RegistrationCodeRepository) CheckIfUserWithCodeExists(code *string) (*bool, *models.INVError) {
	var isExists bool = false

	count, err := utils.CountStatement(table.RegistrationCodes, table.RegistrationCodes.Code.EQ(mysql.String(*code)), rcr.GetDatabaseConnection())
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if registration code exists")
	}
	if count > 0 {
		isExists = true
	}
	return &isExists, nil
}

func (rcr *RegistrationCodeRepository) CreateRegistrationCode(tx *sql.Tx, registrationCode *model.RegistrationCodes) *models.INVError {
	stmt := table.RegistrationCodes.INSERT(
		table.RegistrationCodes.UserID,
		table.RegistrationCodes.Code,
	).VALUES(
		registrationCode.UserID,
		registrationCode.Code,
	)
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating registration code")
	}
	return nil
}

func (rcr *RegistrationCodeRepository) DeleteRegistrationCode(tx *sql.Tx, code *string) *models.INVError {
	stmt := table.RegistrationCodes.DELETE().WHERE(table.RegistrationCodes.Code.EQ(utils.MySqlString(*code)))
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting registration code")
	}
	return nil
}

func (rcr *RegistrationCodeRepository) GetRegistrationCodeById(userId *uuid.UUID) (*model.RegistrationCodes, *models.INVError) {
	stmt := table.RegistrationCodes.SELECT(
		table.RegistrationCodes.UserID,
		table.RegistrationCodes.Code,
	).WHERE(
		table.RegistrationCodes.UserID.EQ(utils.MySqlString(userId.String())),
	)
	var registrationCode model.RegistrationCodes
	err := stmt.Query(rcr.GetDatabaseConnection(), &registrationCode)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_USER_NOT_FOUND.WithDetails("User not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading user")
	}

	return &registrationCode, nil
}


func (rcr *RegistrationCodeRepository) GetUserIdByCode(code *string) (*model.RegistrationCodes, *models.INVError) {
	stmt := table.RegistrationCodes.SELECT(
		table.RegistrationCodes.UserID,
		table.RegistrationCodes.Code,
	).WHERE(
		table.RegistrationCodes.Code.EQ(utils.MySqlString(*code)),
	)
	var registrationCode model.RegistrationCodes
	err := stmt.Query(rcr.GetDatabaseConnection(), &registrationCode)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_USER_NOT_FOUND.WithDetails("Code not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading registration code")
	}

	return &registrationCode, nil
}
