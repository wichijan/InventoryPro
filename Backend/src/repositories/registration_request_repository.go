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
)

type RegistrationRequestRepositoryI interface {
	GetRegistrationRequests() (*[]model.RegistrationRequests, *models.INVError)
	GetRequestByUserId(userId *uuid.UUID) (*model.RegistrationRequests, *models.INVError)
	CreateRequest(tx *sql.Tx, book *model.RegistrationRequests) *models.INVError
	DeleteRequest(tx *sql.Tx, userId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type RegistrationRequestRepository struct {
	managers.DatabaseManagerI
}

func (rr *RegistrationRequestRepository) GetRegistrationRequests() (*[]model.RegistrationRequests, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.RegistrationRequests.AllColumns,
	).FROM(
		table.RegistrationRequests,
	)

	// Execute the query
	var requests []model.RegistrationRequests
	err := stmt.Query(rr.GetDatabaseConnection(), &requests)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &requests, nil
}

func (rr *RegistrationRequestRepository) GetRequestByUserId(userId *uuid.UUID) (*model.RegistrationRequests, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.RegistrationRequests.AllColumns,
	).FROM(
		table.RegistrationRequests,
	).WHERE(
		table.RegistrationRequests.UserID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	var request model.RegistrationRequests
	err := stmt.Query(rr.GetDatabaseConnection(), &request)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &request, nil
}

func (rr *RegistrationRequestRepository) CreateRequest(tx *sql.Tx, request *model.RegistrationRequests) *models.INVError {
	// Create the query
	stmt := table.RegistrationRequests.INSERT(
		table.RegistrationRequests.UserID,
		table.RegistrationRequests.RequestTime,
	).VALUES(
		request.UserID,
		request.RequestTime,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (rr *RegistrationRequestRepository) DeleteRequest(tx *sql.Tx, userId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.RegistrationRequests.DELETE().WHERE(
		table.RegistrationRequests.UserID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}
