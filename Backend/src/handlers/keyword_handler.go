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

// @Summary Get Keywords
// @Description Get Keywords
// @Tags Keywords
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Keywords
// @Failure 500 {object} models.INVErrorMessage
// @Router /keywords [get]
func GetKeywordsHandler(keywordCtrl controllers.KeywordControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		keywords, inv_err := keywordCtrl.GetKeywords()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, keywords)
	}
}

// @Summary Create Keyword
// @Description Create Keyword
// @Tags Keywords
// @Accept  json
// @Produce  json
// @Param keyword body string true "KeywordODT model"
// @Success 200 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /keywords [post]
func CreateKeywordHandler(keywordCtrl controllers.KeywordControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var keywordName models.KeywordODT
		if err := c.ShouldBindJSON(&keywordName); err != nil || keywordName.Name == nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		keywordId, inv_err := keywordCtrl.CreateKeyword(keywordName.Name)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, keywordId)
	}
}

// @Summary Update Keyword
// @Description Update Keyword
// @Tags Keywords
// @Accept  json
// @Produce  json
// @Param keyword body model.Keywords true "Keyword model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /keywords [put]
func UpdateKeywordHandler(keywordCtrl controllers.KeywordControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var keyword model.Keywords
		if err := c.ShouldBindJSON(&keyword); err != nil || keyword.ID == "" || keyword.Keyword == nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := keywordCtrl.UpdateKeyword(&keyword)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Delete Keyword
// @Description Delete Keyword
// @Tags Keywords
// @Accept  json
// @Produce  json
// @Param id path string true "Keyword ID"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /keywords/:id [delete]
func DeleteKeywordHandler(keywordCtrl controllers.KeywordControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := keywordCtrl.DeleteKeyword(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
