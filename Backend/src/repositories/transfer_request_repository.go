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
	CreateTransferRequest(tx *sql.Tx, book *models.TransferRequestCreate) (*string, *models.INVError)
	UpdateTransferRequest(tx *sql.Tx, book *models.TransferRequestUpdate) *models.INVError

	managers.DatabaseManagerI
}

type TransferRequestRepository struct {
	managers.DatabaseManagerI
}

func (trr *TransferRequestRepository) CreateTransferRequest(tx *sql.Tx, request *models.TransferRequestCreate) (*string, *models.INVError) {
	newUUID := uuid.New().String()
	
	// Create the query
	stmt := table.TransferRequests.INSERT(
		table.TransferRequests.TransferRequestID,
		table.TransferRequests.UserID,
		table.TransferRequests.ItemID,
		table.TransferRequests.TargetUserID,
	).VALUES(
		newUUID,
		request.UserID,
		request.ItemID,
		request.TargetUserID,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
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
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}