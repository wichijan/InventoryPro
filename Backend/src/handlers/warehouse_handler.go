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

// @Summary Get warehouses
// @Description Get warehouses
// @Tags Warehouses
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Warehouses
// @Failure 500 {object} models.INVErrorMessage
// @Router /warehouses [get]
func GetWarehousesHandler(warehouseCtrl controllers.WarehouseControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouses, inv_err := warehouseCtrl.GetWarehouses()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, warehouses)
	}
}

// @Summary Get warehouse by id
// @Description Get warehouse by id
// @Tags Warehouses
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse id"
// @Success 200 {object} model.Warehouses
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /warehouses/{id} [get]
func GetWarehouseByIdHandler(warehouseCtrl controllers.WarehouseControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		warehouse, inv_err := warehouseCtrl.GetWarehouseById(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, warehouse)
	}
}

// @Summary Get warehouse by id
// @Description Get warehouse by id
// @Tags Warehouses
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse id"
// @Success 200 {object} models.WarehouseWithRooms
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /warehouses/{id} [get]
func GetWarehousesWithRoomsHandler(warehouseCtrl controllers.WarehouseControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouses, inv_err := warehouseCtrl.GetWarehousesWithRooms()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, warehouses)
	}
}

// @Summary Get warehouse by id
// @Description Get warehouse by id
// @Tags Warehouses
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse id"
// @Success 200 {object} models.WarehouseWithRooms
// @Failure 400 {object} models.INVErrorMessage
// @Failure 404 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /warehouses/{id} [get]
func GetWarehouseByIdWithRoomsHandler(warehouseCtrl controllers.WarehouseControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		warehouses, inv_err := warehouseCtrl.GetWarehouseByIdWithRooms(&id)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, warehouses)
	}
}

// @Summary Create warehouse
// @Description Create warehouse
// @Tags Warehouses
// @Accept  json
// @Produce  json
// @Param WarehousesODT body models.WarehousesODT true "WarehousesODT model"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /warehouses [post]
func CreateWarehouseHandler(warehouseCtrl controllers.WarehouseControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var warehouse models.WarehousesODT
		err := c.ShouldBindJSON(&warehouse)
		if err != nil || utils.ContainsEmptyString(*warehouse.Name) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		warehouseId, inv_err := warehouseCtrl.CreateWarehouse(&warehouse)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, warehouseId)
	}
}

// @Summary Update warehouse
// @Description Update warehouse
// @Tags Warehouses
// @Accept  json
// @Produce  json
// @Param genre body model.Warehouses true "Warehouses model"
// @Success 200 {object} model.Warehouses
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /warehouses [put]
func UpdateWarehouseHandler(warehouseCtrl controllers.WarehouseControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var warehouse model.Warehouses
		err := c.ShouldBindJSON(&warehouse)
		if err != nil || utils.ContainsEmptyString(*warehouse.Name) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := warehouseCtrl.UpdateWarehouse(&warehouse)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, warehouse)
	}
}

/*
// @Summary Delete warehouse
// @Description Delete warehouse
// @Tags Warehouses
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse ID"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /warehouses/{id} [delete]
func DeleteGenre(warehouseCtrl controllers.WarehouseControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		genreId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		inv_err := warehouseCtrl.DeleteGenre(&genreId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
*/
