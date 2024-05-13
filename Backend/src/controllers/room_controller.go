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

func (mc *RoomController) GetRooms() (*[]model.Rooms, *models.INVError) {
	rooms, inv_errors := mc.RoomRepo.GetRooms()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return rooms, nil
}

func (mc *RoomController) CreateRoom(room *model.Rooms) (*uuid.UUID, *models.INVError) {
	if room == nil {
		return nil, inv_errors.INV_BAD_REQUEST
	}

	roomId, inv_errors := mc.RoomRepo.CreateRoom(room)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return roomId, nil
}

func (mc *RoomController) UpdateRoom(room *model.Rooms) *models.INVError {
	inv_errors := mc.RoomRepo.UpdateRoom(room)
	if inv_errors != nil {
		return inv_errors
	}
	return nil
}

func (mc *RoomController) DeleteRoom(roomId *uuid.UUID) *models.INVError {
	// TODO Needs to be implemented
	return nil
}

func (mc *RoomController) GetRoomsById(id *uuid.UUID) (*model.Rooms, *models.INVError) {
	room, inv_errors := mc.RoomRepo.GetRoomsById(id)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return room, nil
}

func (mc *RoomController) GetRoomsWithShelves() (*[]models.RoomWithShelves, *models.INVError) {
	rooms, inv_errors := mc.RoomRepo.GetRoomsWithShelves()
	if inv_errors != nil {
		return nil, inv_errors
	}
	return rooms, nil
}

func (mc *RoomController) GetRoomsByIdWithShelves(id *uuid.UUID) (*models.RoomWithShelves, *models.INVError) {
	room, inv_errors := mc.RoomRepo.GetRoomsByIdWithShelves(id)
	if inv_errors != nil {
		return nil, inv_errors
	}
	return room, nil
}
