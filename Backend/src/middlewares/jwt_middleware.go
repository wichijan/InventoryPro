package middlewares

import (
	"context"

	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		// check if token is set
		token, err := c.Cookie("token")
		if err != nil {
			// token is not set, check if refresh token is set
			refreshToken, err := c.Cookie("refreshToken")
			if err != nil {
				utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
				return
			}
			token, refreshToken, err = utils.RefreshTokens(refreshToken)
			if err != nil {
				utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
				return
			}
			utils.SetJWTCookies(c, token, refreshToken, false)
		}

		userId, err := utils.ValidateToken(token)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		// add userId to request context
		ctx := context.WithValue(c.Request.Context(), models.ContextKeyUserID, userId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
