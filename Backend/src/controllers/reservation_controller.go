package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type ReservationControllerI interface {
	GetReservationByUserId(userId *uuid.UUID) (*[]model.Reservations, *models.INVError)
	GetReservationByItemId(itemId *uuid.UUID) (*[]model.Reservations, *models.INVError)
	GetReservationByItemIdAndUserId(itemId *uuid.UUID, userId *uuid.UUID) (*model.Reservations, *models.INVError)
	CreateReservation(reservation *models.ReservationCreate) (*string, *models.INVError)
	DeleteReservation(userId *uuid.UUID, reservationID *uuid.UUID) *models.INVError
}

type ReservationController struct {
	ReservationRepo repositories.ReservationRepositoryI
}

func (rc *ReservationController) GetReservationByUserId(userId *uuid.UUID) (*[]model.Reservations, *models.INVError) {
	reservations, inv_errors := rc.ReservationRepo.GetReservationByUserId(userId)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return reservations, nil
}

func (rc *ReservationController) GetReservationByItemId(itemId *uuid.UUID) (*[]model.Reservations, *models.INVError) {
	reservations, inv_errors := rc.ReservationRepo.GetReservationByItemId(itemId)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return reservations, nil
}

func (rc *ReservationController) GetReservationByItemIdAndUserId(itemId *uuid.UUID, userId *uuid.UUID) (*model.Reservations, *models.INVError) {
	reservation, inv_errors := rc.ReservationRepo.GetReservationByItemIdAndUserId(itemId, userId)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return reservation, nil
}

func (rc *ReservationController) CreateReservation(reservation *models.ReservationCreate) (*string, *models.INVError) {
	tx, err := rc.ReservationRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	if reservation == nil {
		return nil, inv_errors.INV_BAD_REQUEST
	}

	reservationID, inv_error := rc.ReservationRepo.CreateReservation(tx, reservation)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return reservationID, nil
}

func (rc *ReservationController) DeleteReservation(userId *uuid.UUID, reservationID *uuid.UUID) *models.INVError {
	tx, err := rc.ReservationRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	inv_error := rc.ReservationRepo.DeleteReservation(tx, userId, reservationID)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}