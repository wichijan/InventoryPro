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

// @Summary Add Role to User
// @Description Add Role to User
// @Tags UserRoles
// @Accept  json
// @Produce  json
// @Param RoleIdODT body models.RoleIdODT true "RoleIdODT model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /user-roles/add-role [post]
func AddRoleToUserHandler(userRoleCtrl controllers.UserRoleControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		var roleName models.RoleIdODT
		err := c.ShouldBindJSON(&roleName)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		itemRole := model.UserRoles{
			UserID: userId.String(),
			RoleID: roleName.RoleID,
		}

		inv_err := userRoleCtrl.CreateUserRole(&itemRole)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Add Role to User
// @Description Add Role to User
// @Tags UserRoles
// @Accept  json
// @Produce  json
// @Param RoleIdODT body models.RoleIdODT true "RoleIdODT model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /user-roles/add-role [delete]
func RemoveRoleFromUserHandler(userRoleCtrl controllers.UserRoleControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		var roleName models.RoleIdODT
		err := c.ShouldBindJSON(&roleName)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		itemRole := model.UserRoles{
			UserID: userId.String(),
			RoleID: roleName.RoleID,
		}

		inv_err := userRoleCtrl.DeleteUserRole(&itemRole)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
