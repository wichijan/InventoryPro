package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wichijan/InventoryPro/src/models"
)

func HandleErrorAndAbort(c *gin.Context, err *models.INVError) {
	log.Printf("Error while handling request: %d %v %v", err.Status, err.ErrorMessage, err.Details)
	c.AbortWithStatusJSON(err.Status, err)
}
