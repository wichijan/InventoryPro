package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

// @Summary Get Roles
// @Description Get Roles
// @Tags Roles
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Roles
// @Failure 500 {object} models.INVErrorMessage
// @Router /roles [get]
func GetRolesHandler(roleCtrl controllers.RoleControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, inv_err := roleCtrl.GetRoles()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, roles)
	}
}

// @Summary Create Item
// @Description Create Item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param role body models.RoleODT true "RoleODT model"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /roles [post]
func CreateRoleHandler(roleCtrl controllers.RoleControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role models.RoleODT
		err := c.ShouldBindJSON(&role)
		if err != nil || role.RoleName == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		roleId, inv_err := roleCtrl.CreateRole(&role.RoleName)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, roleId)
	}
}

// @Summary Create Item
// @Description Create Item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param role body model.Roles true "Roles model"
// @Success 200 {object}
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /roles [put]
func UpdateRoleHandler(roleCtrl controllers.RoleControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role model.Roles
		err := c.ShouldBindJSON(&role)
		if err != nil || *role.RoleName == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := roleCtrl.UpdateRole(&role)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, nil)
	}
}
