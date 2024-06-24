package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

// @Summary Get rooms
// @Description Get rooms
// @Tags Rooms
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Rooms
// @Failure 500 {object} models.INVErrorMessage
// @Router /rooms [get]
func GetRoomsHandler(roomCtrl controllers.RoomControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		rooms, inv_err := roomCtrl.GetRooms()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, rooms)
	}
}

// @Summary Get room by id
// @Description Get room by id
// @Tags Rooms
// @Accept  json
// @Produce  json
// @Param id path string true "Room id"
// @Success 200 {object} model.Rooms
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /rooms/{id} [get]
func GetRoomsByIdHandle(roomCtrl controllers.RoomControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid room id"))
			return
		}

		room, inv_err := roomCtrl.GetRoomsById(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, room)
	}
}

// @Summary Get rooms with shelves
// @Description Get rooms with shelves
// @Tags Rooms
// @Accept  json
// @Produce  json
// @Param id path string true "room id"
// @Success 200 {object} models.RoomWithShelves
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /rooms-with-shelves [get]
func GetRoomsWithShelvesHandle(roomCtrl controllers.RoomControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		rooms, inv_err := roomCtrl.GetRoomsWithShelves()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, rooms)
	}
}

// @Summary Get room by id with shelves
// @Description Get room by id with shelves
// @Tags Rooms
// @Accept  json
// @Produce  json
// @Param id path string true "Room id"
// @Success 200 {object} models.RoomWithShelves
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /rooms-with-shelves/{id} [get]
func GetRoomsByIdWithShelvesHandle(roomCtrl controllers.RoomControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid room id"))
			return
		}

		room, inv_err := roomCtrl.GetRoomsByIdWithShelves(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, room)
	}
}

// @Summary Create Room
// @Description Create Room
// @Tags Rooms
// @Accept  json
// @Produce  json
// @Param RoomsODT body models.RoomsODT true "RoomsODT model"
// @Success 201 {object} model.Rooms
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /rooms [post]
func CreateRoomHandle(roomCtrl controllers.RoomControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var room models.RoomsODT
		err := c.ShouldBindJSON(&room)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(*room.Name) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid room name"))
			return
		}

		roomId, inv_err := roomCtrl.CreateRoom(&room)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, roomId)
	}
}

// @Summary Update room
// @Description Update room
// @Tags Rooms
// @Accept  json
// @Produce  json
// @Param room body model.Rooms true "Room model"
// @Success 200 {object} model.Rooms
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /rooms [put]
func UpdateRoomHandle(roomCtrl controllers.RoomControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var room model.Rooms
		err := c.ShouldBindJSON(&room)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(room.Name) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid room id"))
			return
		}

		inv_err := roomCtrl.UpdateRoom(&room)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, room)
	}
}

// @Summary Delete room
// @Description Delete room
// @Tags Rooms
// @Accept  json
// @Produce  json
// @Param id path string true "Room ID"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /rooms/{id} [delete]
func DeleteRoomHandle(roomCtrl controllers.RoomControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		genreId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := roomCtrl.DeleteRoom(&genreId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
