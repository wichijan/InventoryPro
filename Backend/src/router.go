package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wichijan/InventoryPro/src/controllers"
	"github.com/wichijan/InventoryPro/src/handlers"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type Controllers struct {
	WarehouseController controllers.WarehouseControllerI
}

func createRouter(dbConnection *sql.DB) *gin.Engine {
	router := gin.Default()

	// Attach Middleware
	//router.Use(middlewares.CorsMiddleware())

	// Create api groups, with special middleware
	publicRoutes := router.Group("/")
	//securedRoutes := router.Group("/", middlewares.JwtAuthMiddleware())
	//adminRoutes := router.Group("/", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware())

	// Create managers and repositories
	databaseManager := &managers.DatabaseManager{
		Connection: dbConnection,
	}

	// Create repositories
	warehouseRepo := &repositories.WarehouseRepository{
		DatabaseManager: databaseManager,
	}

	// Create controllers
	controller := Controllers{
		WarehouseController: &controllers.WarehouseController{
			WarehouseRepo: warehouseRepo,
		},
	}

	publicRoutes.Handle(http.MethodGet, "/warehouses", handlers.GetWarehouses(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehousesWithRooms", handlers.GetWarehousesWithRooms(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehousesWithRooms/:id", handlers.GetWarehouseByIdWithRooms(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehouses/:id", handlers.GetWarehouseById(controller.WarehouseController))
	publicRoutes.Handle(http.MethodPost, "/warehouses", handlers.CreateWarehouse(controller.WarehouseController))
	publicRoutes.Handle(http.MethodPut, "/warehouses", handlers.UpdateWarehouse(controller.WarehouseController))

	return router
}
