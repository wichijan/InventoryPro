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

// @Summary Return Reserve Item
// @Description Return Reserve Item when logged-in
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "item id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /add-item-to-quick-shelf/:id [post]
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
