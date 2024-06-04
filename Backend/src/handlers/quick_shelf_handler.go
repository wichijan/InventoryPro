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

// @Summary Add item to Quick Shelf
// @Description Add item from user to Quick Shelf
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Param ItemQuickShelfInsertODT body models.ItemQuickShelfInsertODT true "ItemQuickShelfInsertODT model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /add-item-to-quick-shelf [post]
func AddToQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		var itemQuickShelfODT models.ItemQuickShelfInsertODT
		err := c.ShouldBindJSON(&itemQuickShelfODT)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		var itemQuick models.ItemQuickShelfInsert
		itemQuick.UserID = *userId
		itemQuick.QuickShelfID = itemQuickShelfODT.QuickShelfID
		itemQuick.ItemID = itemQuickShelfODT.ItemID
		itemQuick.Quantity = itemQuickShelfODT.Quantity

		inv_err := itemQuickShelfCtrl.InsertItemInQuickShelf(&itemQuick)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Bring item from Quick shelf to regular shelf
// @Description You can only take all quantity of item from quick shelf to return it to regular shelf
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Param ItemQuickShelfRemoveSingleItem body models.ItemQuickShelfRemoveSingleItem true "ItemQuickShelfRemoveSingleItem model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /remove-item-to-quick-shelf [post]
func RemoveItemFromQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemQuickShelfODT models.ItemQuickShelfRemoveSingleItem
		err := c.ShouldBindJSON(&itemQuickShelfODT)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		var itemQuick models.ItemQuickShelfRemoveSingleItem
		itemQuick.QuickShelfID = itemQuickShelfODT.QuickShelfID
		itemQuick.ItemID = itemQuickShelfODT.ItemID

		inv_err := itemQuickShelfCtrl.RemoveItemFromQuickShelf(&itemQuick.ItemID, &itemQuick.QuickShelfID)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Clear all times from Quick shelf
// @Description Remove all items from quick shelf
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Param id path string true "quick shelf id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /clear-quick-shelf/:id [delete]
func ClearQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelfId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := itemQuickShelfCtrl.ClearQuickShelf(&quickShelfId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Get Items in Quick shelf
// @Description Get all items in quick shelf
// @Tags Quick Shelf
// @Accept  json
// @Produce  json
// @Param id path string true "quick shelf id"
// @Success 200 {array} models.ItemQuickShelfInsert
// @Failure 400 {object} models.INVErrorMessage
// @Router /quick-shelf/:id [get]
func GetItemsInQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelfId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		items, inv_err := itemQuickShelfCtrl.GetItemsInQuickShelf(&quickShelfId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, items)
	}
}
