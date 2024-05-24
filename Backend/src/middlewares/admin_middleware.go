package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

func AdminMiddleware(databaseManager managers.DatabaseManagerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
		}

		// Check user role
		var userRoles models.UserRoleWithName

		// Create the query
		stmt := mysql.SELECT(
			table.UserRoles.UserID,
			table.Roles.RoleName,
		).FROM(
			table.UserRoles.
				LEFT_JOIN(table.Roles, table.Roles.ID.EQ(table.UserRoles.RoleID)),
		).WHERE(
			table.UserRoles.UserID.EQ(mysql.String(userId.String())),
		)

		// Execute the query
		err := stmt.Query(databaseManager.GetDatabaseConnection(), &userRoles)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_FORBIDDEN)
		}

		for _, a := range userRoles.UserNames {
			if a.RoleName == "Admin" {
				c.Next()
				return
			}
		}

		utils.HandleErrorAndAbort(c, inv_errors.INV_FORBIDDEN)
	}
}
