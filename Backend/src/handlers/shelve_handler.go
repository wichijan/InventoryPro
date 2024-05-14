package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/utils"
)

// @Summary Get shelves
// @Description Get shelves
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Success 200 {array} models.OwnShelve
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelves [get]
func GetShelvesHandler(shelveCtrl controllers.ShelveControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		shelves, inv_err := shelveCtrl.GetShelves()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, shelves)
	}
}

// @Summary Get shelve by id
// @Description Get shelve by id
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Param id path string true "Shelve id"
// @Success 200 {object} models.OwnShelve
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelves/{id} [get]
func GetShelveByIdHandler(shelveCtrl controllers.ShelveControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		shelve, inv_err := shelveCtrl.GetShelveById(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, shelve)
	}
}

// @Summary Get shelves with items
// @Description Get shelves with items
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ShelveWithItems
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelveswithitems [get]
func GetShelvesWithItemsHandler(shelveCtrl controllers.ShelveControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		shelvesWithItems, inv_err := shelveCtrl.GetShelvesWithItems()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, shelvesWithItems)
	}
}

// @Summary Get shelve by id with items
// @Description Get shelve by id with items
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Param id path string true "shelve id"
// @Success 200 {object} models.ShelveWithItems
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelveswithitems/{id} [get]
func GetShelveByIdWithItemsHandler(shelveCtrl controllers.ShelveControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		shelveWithItems, inv_err := shelveCtrl.GetShelveByIdWithItems(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, shelveWithItems)
	}
}

/*
// @Summary Create Shelve
// @Description Create Shelve
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Param room body model.Rooms true "Shelve model"
// @Success 201 {object} model.Shelves
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelves [post]
func CreateShelveHandler(shelveCtrl controllers.ShelveControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shelve models.OwnShelve
		err := c.ShouldBindJSON(&shelve)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		shelveId, inv_err := shelveCtrl.CreateShelve(&shelve)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, shelveId)
	}
}
*/

/*
// @Summary Update shelve
// @Description Update shelve
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Param shelve body model.Rooms true "Shelve model"
// @Success 200 {object} model.Shelves
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelves [put]
func UpdateShelveHandler(shelveCtrl controllers.ShelveControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shelve models.OwnShelve
		err := c.ShouldBindJSON(&shelve)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := shelveCtrl.UpdateShelve(&shelve)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, shelve)
	}
}
*/

/*
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
func DeleteGenre(shelveCtrl controllers.ShelveControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		genreId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := shelveCtrl.DeleteGenre(&genreId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
*/
