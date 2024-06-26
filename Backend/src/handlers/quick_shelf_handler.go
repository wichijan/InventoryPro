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

// @Summary Get all Quick Shelves
// @Description Get all Quick Shelves
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Success 200 {array} models.QuickShelfWithItems
// @Failure 400 {object} models.INVErrorMessage
// @Router /quick-shelves [get]
func GetQuickShelvesHandler(quickShelfCtrl controllers.QuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelves, inv_err := quickShelfCtrl.GetQuickShelves()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, quickShelves)
	}
}

// @Summary Get Quick Shelf by Id
// @Description Get Quick Shelf by Id
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Success 200 {object} models.QuickShelfWithItems
// @Failure 400 {object} models.INVErrorMessage
// @Router /quick-shelves/{id} [get]
func GetQuickShelfByIdHandler(quickShelfCtrl controllers.QuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quick shelf id"))
			return
		}

		quickShelf, inv_err := quickShelfCtrl.GetQuickShelfById(&quickShelId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, quickShelf)
	}
}

// @Summary Create Quick Shelf
// @Description Create Quick Shelf
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Param QuickShelfCreate body models.QuickShelfCreate true "QuickShelfCreate model"
// @Success 200 {string} string
// @Failure 400 {object} models.INVErrorMessage
// @Router /quick-shelves [post]
func CreateQuickShelfHandler(quickShelfCtrl controllers.QuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var quickShelf models.QuickShelfCreate
		err := c.ShouldBindJSON(&quickShelf)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}

		quickShelfID, inv_err := quickShelfCtrl.CreateQuickShelf(&quickShelf)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, quickShelfID)
	}
}

// @Summary Update Quick Shelf
// @Description Update Quick Shelf
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Param QuickShelfUpdate body model.QuickShelves true "QuickShelves model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /quick-shelves [put]
func UpdateQuickShelfHandler(quickShelfCtrl controllers.QuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var quickShelf model.QuickShelves
		err := c.ShouldBindJSON(&quickShelf)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if quickShelf.QuickShelfID == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quick shelf id"))
			return
		}
		if quickShelf.RoomID == nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid room id"))
			return
		}

		inv_err := quickShelfCtrl.UpdateQuickShelf(&quickShelf)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Delete Quick Shelf
// @Description Delete Quick Shelf
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Param id path string true "quick shelf id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /quick-shelves/{id} [delete]
func DeleteQuickShelfHandler(quickShelfCtrl controllers.QuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelfId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quick shelf id"))
			return
		}

		inv_err := quickShelfCtrl.DeleteQuickShelf(&quickShelfId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
