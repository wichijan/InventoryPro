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

type ReservationRepositoryI interface {
	GetReservationByUserId(userId *uuid.UUID) (*[]model.Reservations, *models.INVError)
	GetReservationByItemId(itemId *uuid.UUID) (*[]model.Reservations, *models.INVError)
	GetReservationById(reservationId *uuid.UUID) (*model.Reservations, *models.INVError)
	GetReservationByItemIdAndUserId(itemId *uuid.UUID, userId *uuid.UUID) (*model.Reservations, *models.INVError)
	CreateReservation(tx *sql.Tx, reservation *models.ReservationCreate) (*uuid.UUID, *models.INVError)
	DeleteReservation(tx *sql.Tx, userId *uuid.UUID, reservationID *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type ReservationRepository struct {
	managers.DatabaseManagerI
}

func (rr *ReservationRepository) GetReservationByUserId(userId *uuid.UUID) (*[]model.Reservations, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.Reservations.AllColumns,
	).FROM(
		table.Reservations,
	).WHERE(
		table.Reservations.UserID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	var reservations []model.Reservations
	err := stmt.Query(rr.GetDatabaseConnection(), &reservations)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &reservations, nil
}


func (rr *ReservationRepository) GetReservationById(reservationId *uuid.UUID) (*model.Reservations, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.Reservations.AllColumns,
	).FROM(
		table.Reservations,
	).WHERE(
		table.Reservations.ReservationID.EQ(mysql.String(reservationId.String())),
	)

	// Execute the query
	var reservation model.Reservations
	err := stmt.Query(rr.GetDatabaseConnection(), &reservation)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &reservation, nil
}

func (rr *ReservationRepository) GetReservationByItemId(itemId *uuid.UUID) (*[]model.Reservations, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.Reservations.AllColumns,
	).FROM(
		table.Reservations,
	).WHERE(
		table.Reservations.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	var reservations []model.Reservations
	err := stmt.Query(rr.GetDatabaseConnection(), &reservations)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &reservations, nil
}

func (rr *ReservationRepository) GetReservationByItemIdAndUserId(itemId *uuid.UUID, userId *uuid.UUID) (*model.Reservations, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.Reservations.AllColumns,
	).FROM(
		table.Reservations,
	).WHERE(
		table.Reservations.ItemID.EQ(mysql.String(itemId.String())).
			AND(table.Reservations.UserID.EQ(mysql.String(userId.String()))),
	)

	// Execute the query
	var reservation model.Reservations
	err := stmt.Query(rr.GetDatabaseConnection(), &reservation)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &reservation, nil
}

func (rr *ReservationRepository) CreateReservation(tx *sql.Tx, reservation *models.ReservationCreate) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()
	
	// Create the query
	stmt := table.Reservations.INSERT(
		table.Reservations.ReservationID,
		table.Reservations.UserID,
		table.Reservations.ItemID,
		table.Reservations.Quantity,
		table.Reservations.TimeFrom,
		table.Reservations.TimeTo,
	).VALUES(
		uuid.String(),
		reservation.UserID,
		reservation.ItemID,
		reservation.Quantity,
		reservation.TimeFrom,
		reservation.TimeTo,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &uuid, nil
}

func (rr *ReservationRepository) DeleteReservation(tx *sql.Tx, userId *uuid.UUID, reservationID *uuid.UUID) *models.INVError {
	// Create the delete statement
	deleteQuery := table.Reservations.DELETE().WHERE(
		table.Reservations.ReservationID.EQ(mysql.String(reservationID.String())).
			AND(table.Reservations.UserID.EQ(mysql.String(userId.String()))),
	)

	// Execute the query
	_, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}