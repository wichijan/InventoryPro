package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wichijan/InventoryPro/src/controllers"
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
