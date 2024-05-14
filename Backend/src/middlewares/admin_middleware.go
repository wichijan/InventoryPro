package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
		}

		adminId, err := uuid.Parse("dddddddd-dddd-dddd-dddd-dddddddddddd")
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_INTERNAL_ERROR)
			return
		}

		if *userId != adminId {
			utils.HandleErrorAndAbort(c, inv_errors.INV_FORBIDDEN)
			return
		}
		c.Next()
	}
}
