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
	CreateRegistrationCode(tx *sql.Tx, registrationCode *model.RegistrationCodes) *models.INVError
	DeleteRegistrationCode(tx *sql.Tx, userId *uuid.UUID) *models.INVError

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
