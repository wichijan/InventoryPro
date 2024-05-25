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

// @Summary Get UserTypes
// @Description Get UserTypes
// @Tags UserTypes
// @Accept  json
// @Produce  json
// @Success 200 {array} model.UserTypes
// @Failure 500 {object} models.INVErrorMessage
// @Router /user-types [get]
func GetUserTypesHandler(userTypeCtrl controllers.UserTypeControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userTypes, inv_err := userTypeCtrl.GetUserTypes()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, userTypes)
	}
}

// @Summary Create UserType
// @Description Create UserType
// @Tags UserTypes
// @Accept  json
// @Produce  json
// @Param userType body string true "UserTypeODT model"
// @Success 200 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /user-types [post]
func CreateUserTypeHandler(userTypeCtrl controllers.UserTypeControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userTypeName models.UserTypeODT
		if err := c.ShouldBindJSON(&userTypeName); err != nil || userTypeName.Name == nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		userTypeId, inv_err := userTypeCtrl.CreateUserType(userTypeName.Name)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, userTypeId)
	}
}

// @Summary Update UserType
// @Description Update UserType
// @Tags UserTypes
// @Accept  json
// @Produce  json
// @Param userType body model.UserTypes true "UserType model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /user-types [put]
func UpdateUserTypeHandler(userTypeCtrl controllers.UserTypeControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userType model.UserTypes
		err := c.ShouldBindJSON(&userType)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := userTypeCtrl.UpdateUserType(&userType)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Delete UserType
// @Description Delete UserType
// @Tags UserTypes
// @Accept  json
// @Produce  json
// @Param id path string true "UserType id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /user-types [delete]
func DeleteUserTypeHandler(userTypeCtrl controllers.UserTypeControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := userTypeCtrl.DeleteUserType(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
