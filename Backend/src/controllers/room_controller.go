package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type RoomControllerI interface {
	GetRooms() (*[]model.Rooms, *models.INVError)
	GetRoomsById(id *uuid.UUID) (*model.Rooms, *models.INVError)
	GetRoomsWithShelves() (*[]models.RoomWithShelves, *models.INVError)
	GetRoomsByIdWithShelves(id *uuid.UUID) (*models.RoomWithShelves, *models.INVError)
	CreateRoom(room *model.Rooms) (*uuid.UUID, *models.INVError)
	UpdateRoom(room *model.Rooms) *models.INVError
	DeleteRoom(roomId *uuid.UUID) *models.INVError
}

type RoomController struct {
	RoomRepo repositories.RoomRepositoryI
}

func (rc *RoomController) GetRooms() (*[]model.Rooms, *models.INVError) {
	rooms, inv_errors := rc.RoomRepo.GetRooms()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return rooms, nil
}

func (rc *RoomController) CreateRoom(room *model.Rooms) (*uuid.UUID, *models.INVError) {
	tx, err := rc.RoomRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	if room == nil {
		return nil, inv_errors.INV_BAD_REQUEST
	}

	roomId, inv_error := rc.RoomRepo.CreateRoom(tx, room)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return roomId, nil
}

func (rc *RoomController) UpdateRoom(room *model.Rooms) *models.INVError {
	tx, err := rc.RoomRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	inv_error := rc.RoomRepo.UpdateRoom(tx, room)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (rc *RoomController) DeleteRoom(roomId *uuid.UUID) *models.INVError {
	// TODO Needs to be implemented
	tx, err := rc.RoomRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (rc *RoomController) GetRoomsById(id *uuid.UUID) (*model.Rooms, *models.INVError) {
	room, inv_errors := rc.RoomRepo.GetRoomsById(id)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return room, nil
}

func (rc *RoomController) GetRoomsWithShelves() (*[]models.RoomWithShelves, *models.INVError) {
	rooms, inv_errors := rc.RoomRepo.GetRoomsWithShelves()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return rooms, nil
}

func (rc *RoomController) GetRoomsByIdWithShelves(id *uuid.UUID) (*models.RoomWithShelves, *models.INVError) {
	room, inv_errors := rc.RoomRepo.GetRoomsByIdWithShelves(id)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return room, nil
}
