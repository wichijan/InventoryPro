package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
	"github.com/wichijan/InventoryPro/src/websocket"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
)

// @Summary Websocket Handler - js -> new WebSocket("ws://localhost:8080/ws/:roomId")
// @Description RoomId can be empty for public notifications. RoomId is required for chat functions if ever implemented. IMPORTANT: WebSocket has to be called / created after Login.
// @Tags Websocket
// @Success 200
// @Router /ws/:roomId [get]
func WebsocketHandler(databaseManager managers.DatabaseManagerI, hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomId := c.Param("roomId")
		if roomId == "" {
			roomId = utils.WEBSOCKET_DEFAULT_ROOM
		}

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
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
		if err != nil && err.Error() != "qrm: no rows in result set" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_FORBIDDEN)
			return
		}

		var isAdmin bool = false
		for _, a := range userRoles.UserNames {
			if a.RoleName == "Admin" {
				isAdmin = true
				continue
			}
		}
		websocket.ServeWS(c, roomId, userId.String(), isAdmin, hub)
	}
}
