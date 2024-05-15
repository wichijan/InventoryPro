package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

// @Summary Get items
// @Description Get items
// @Tags Items
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ItemWithEverything
// @Failure 500 {object} models.INVErrorMessage
// @Router /items [get]
func GetItemsHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, inv_err := itemCtrl.GetItems()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

// @Summary Get item by id
// @Description Get item by id
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200 {array} models.ItemWithEverything
// @Failure 500 {object} models.INVErrorMessage
// @Router /items/:id [get]
func GetItemByIdHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		item, inv_err := itemCtrl.GetItemById(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

// @Summary Create Item
// @Description Create Item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param room body model.Rooms true "ItemWithStatus model"
// @Success 201 {object} models.ItemWithStatus
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /items [post]
func CreateItemHandler(itemCtrl controllers.ItemControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ItemWithStatus
		err := c.ShouldBindJSON(&item)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		itemId, inv_err := itemCtrl.CreateItem(&item)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, itemId)
	}
}
