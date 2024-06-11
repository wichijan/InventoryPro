package repositories

import (
	"database/sql"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type TransferRequestRepositoryI interface {
	GetTransferRequestByUserId(userId uuid.UUID) (*[]models.TransferRequestSelect, *models.INVError)
	GetTransferRequestById(transferId uuid.UUID) (*models.TransferRequestSelect, *models.INVError)
	CreateTransferRequest(tx *sql.Tx, book *models.TransferRequestCreate) (*uuid.UUID, *models.INVError)
	UpdateTransferRequest(tx *sql.Tx, book *models.TransferRequestUpdate) *models.INVError
	DeleteTransferRequest(tx *sql.Tx, trans *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type TransferRequestRepository struct {
	managers.DatabaseManagerI
}

func (trr *TransferRequestRepository) GetTransferRequestByUserId(userId uuid.UUID) (*[]models.TransferRequestSelect, *models.INVError) {
	// Create the query
	stmt := table.TransferRequests.SELECT(
		table.TransferRequests.AllColumns,
	).WHERE(
		table.TransferRequests.UserID.EQ(mysql.String(userId.String())),
	).UNION(
		table.TransferRequests.SELECT(
			table.TransferRequests.AllColumns,
		).WHERE(
			table.TransferRequests.TargetUserID.EQ(mysql.String(userId.String())),
		),
	)

	// Execute the query
	var transferRequest []models.TransferRequestSelect
	err := stmt.Query(trr.GetDatabaseConnection(), &transferRequest)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error getting transfer request")
	}

	return &transferRequest, nil
}

func (trr *TransferRequestRepository) GetTransferRequestById(transferId uuid.UUID) (*models.TransferRequestSelect, *models.INVError) {
	// Create the query
	stmt := table.TransferRequests.SELECT(
		table.TransferRequests.AllColumns,
	).WHERE(
		table.TransferRequests.TransferRequestID.EQ(mysql.String(transferId.String())),
	)

	// Execute the query
	var transferRequest models.TransferRequestSelect
	err := stmt.Query(trr.GetDatabaseConnection(), &transferRequest)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error getting transfer request")
	}

	return &transferRequest, nil
}

func (trr *TransferRequestRepository) CreateTransferRequest(tx *sql.Tx, request *models.TransferRequestCreate) (*uuid.UUID, *models.INVError) {
	newUUID := uuid.New()

	// Create the query
	stmt := table.TransferRequests.INSERT(
		table.TransferRequests.TransferRequestID,
		table.TransferRequests.UserID,
		table.TransferRequests.ItemID,
		table.TransferRequests.TargetUserID,
	).VALUES(
		newUUID.String(),
		request.UserID,
		request.ItemID,
		request.TargetUserID,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transfer request")
	}

	return &newUUID, nil
}

func (trr *TransferRequestRepository) UpdateTransferRequest(tx *sql.Tx, request *models.TransferRequestUpdate) *models.INVError {
	// Create the query
	stmt := table.TransferRequests.UPDATE(
		table.TransferRequests.UserID,
		table.TransferRequests.ItemID,
		table.TransferRequests.TargetUserID,
		table.TransferRequests.IsAccepted,
	).SET(
		request.UserID,
		request.ItemID,
		request.TargetUserID,
		request.IsAccepted,
	).WHERE(
		table.TransferRequests.TransferRequestID.EQ(mysql.String(request.TransferRequestID)),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating transfer request")
	}

	return nil
}

func (trr *TransferRequestRepository) DeleteTransferRequest(tx *sql.Tx, transferId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.TransferRequests.DELETE().WHERE(
		table.TransferRequests.TransferRequestID.EQ(mysql.String(transferId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting transfer request")
	}

	return nil
}