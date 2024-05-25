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

// @Summary Get shelve-type
// @Description Get shelve-type
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ShelveTypes
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelve-types [get]
func GetShelveTypesHandler(shelveTypeCtrl controllers.ShelveTypeControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		shelveTypes, inv_err := shelveTypeCtrl.GetShelveTypes()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, shelveTypes)
	}
}

// @Summary Create shelve-type
// @Description Create shelve-type
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Param shelveTypeName body string true "shelve-type name"
// @Success 200 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelve-types [post]
func CreateShelveTypeHandler(shelveTypeCtrl controllers.ShelveTypeControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shelveTypeName models.ShelveTypeODT
		if err := c.ShouldBindJSON(&shelveTypeName); err != nil || shelveTypeName.ShelveTypeName == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		shelveTypeId, inv_err := shelveTypeCtrl.CreateShelveType(&shelveTypeName.ShelveTypeName)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, shelveTypeId)
	}
}

// @Summary Update shelve-type
// @Description Update shelve-type
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Param shelveType body model.ShelveTypes true "shelve-type model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelve-types [put]
func UpdateShelveTypeHandler(shelveTypeCtrl controllers.ShelveTypeControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shelveType model.ShelveTypes
		if err := c.ShouldBindJSON(&shelveType); err != nil || shelveType.TypeName == nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := shelveTypeCtrl.UpdateShelveType(&shelveType)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Delete shelve-type by id
// @Description Delete shelve-type by id
// @Tags Shelves
// @Accept  json
// @Produce  json
// @Param id path string true "shelve-type id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /shelve-types/{id} [delete]
func DeleteShelveTypeHandler(shelveTypeCtrl controllers.ShelveTypeControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := shelveTypeCtrl.DeleteShelveType(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
