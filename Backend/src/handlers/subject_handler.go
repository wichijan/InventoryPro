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

// @Summary Get Subjects
// @Description Get Subjects
// @Tags Subjects
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Subjects
// @Failure 500 {object} models.INVErrorMessage
// @Router /subjects [get]
func GetSubjectsHandler(subjectCtrl controllers.SubjectControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		subjects, inv_err := subjectCtrl.GetSubjects()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, subjects)
	}
}

// @Summary Create Subject
// @Description Create Subject
// @Tags Subjects
// @Accept  json
// @Produce  json
// @Param subject body string true "SubjectODT model"
// @Success 200 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /subjects [post]
func CreateSubjectHandler(subjectCtrl controllers.SubjectControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var subjectName models.SubjectODT
		if err := c.ShouldBindJSON(&subjectName); err != nil || subjectName.Name == nil || subjectName.Description == nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		subjectId, inv_err := subjectCtrl.CreateSubject(&subjectName)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, subjectId)
	}
}

// @Summary Update Subject
// @Description Update Subject
// @Tags Subjects
// @Accept  json
// @Produce  json
// @Param subject body model.Subjects true "Subject model"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /subjects [put]
func UpdateSubjectHandler(subjectCtrl controllers.SubjectControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var subject model.Subjects
		if err := c.ShouldBindJSON(&subject); err != nil || subject.Name == nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := subjectCtrl.UpdateSubject(&subject)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Delete Subject
// @Description Delete Subject
// @Tags Subjects
// @Accept  json
// @Produce  json
// @Param id path string true "Subject id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /subjects/{id} [delete]
func DeleteSubjectHandler(subjectCtrl controllers.SubjectControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := subjectCtrl.DeleteSubject(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
