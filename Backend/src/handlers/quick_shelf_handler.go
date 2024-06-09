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
func GetQuickShelvesHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelves, inv_err := itemQuickShelfCtrl.GetQuickShelves()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, quickShelves)
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
func CreateQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var quickShelf models.QuickShelfCreate
		err := c.ShouldBindJSON(&quickShelf)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}

		quickShelfID, inv_err := itemQuickShelfCtrl.CreateQuickShelf(&quickShelf)
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
func UpdateQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
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

		inv_err := itemQuickShelfCtrl.UpdateQuickShelf(&quickShelf)
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
// @Router /quick-shelves/:id [delete]
func DeleteQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelfId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quick shelf id"))
			return
		}

		inv_err := itemQuickShelfCtrl.DeleteQuickShelf(&quickShelfId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

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
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
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
		var itemQuick models.ItemQuickShelfRemoveSingleItem
		err := c.ShouldBindJSON(&itemQuick)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}

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
// @Router /clear-quick-shelf/:id [post]
func ClearQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelfId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quick shelf id"))
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
// @Success 200 {array} models.GetQuickShelf
// @Failure 400 {object} models.INVErrorMessage
// @Router /quick-shelf/:id [get]
func GetItemsInQuickShelfHandler(itemQuickShelfCtrl controllers.ItemQuickShelfControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		quickShelfId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid quick shelf id"))
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
