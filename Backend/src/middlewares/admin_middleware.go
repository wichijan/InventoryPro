package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

func AdminMiddleware(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
		}

		isAdmin, inv_err := userCtrl.IsAdmin(userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		if isAdmin {
			c.Next()
			return
		}

		utils.HandleErrorAndAbort(c, inv_errors.INV_FORBIDDEN.WithDetails("User has not the correct role"))
	}
}
