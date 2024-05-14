package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wichijan/InventoryPro/src/controllers"
	"github.com/wichijan/InventoryPro/src/docs"
	"github.com/wichijan/InventoryPro/src/handlers"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/middlewares"
	"github.com/wichijan/InventoryPro/src/repositories"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controllers struct {
	WarehouseController controllers.WarehouseControllerI
	RoomController      controllers.RoomControllerI
	ShelveController    controllers.ShelveControllerI
	ItemController      controllers.ItemControllerI
}

func createRouter(dbConnection *sql.DB) *gin.Engine {
	router := gin.Default()

	// Attach Middleware
	router.Use(middlewares.CorsMiddleware())
	//securedRoutes := router.Group("/", middlewares.JwtAuthMiddleware())
	//adminRoutes := router.Group("/", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware())

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

	roomRepo := &repositories.RoomRepository{
		DatabaseManager: databaseManager,
	}

	shelveRepo := &repositories.ShelveRepository{
		DatabaseManager: databaseManager,
	}

	itemRepo := &repositories.ItemRepository{
		DatabaseManager: databaseManager,
	}

	// Create controllers
	controller := Controllers{
		WarehouseController: &controllers.WarehouseController{
			WarehouseRepo: warehouseRepo,
		},
		RoomController: &controllers.RoomController{
			RoomRepo: roomRepo,
		},
		ShelveController: &controllers.ShelveController{
			ShelveRepo: shelveRepo,
		},
		ItemController: &controllers.ItemController{
			ItemRepo: itemRepo,
		},
	}

	// Warehouse routes
	publicRoutes.Handle(http.MethodGet, "/warehouses", handlers.GetWarehousesHandler(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehousesWithRooms", handlers.GetWarehousesWithRoomsHandler(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehousesWithRooms/:id", handlers.GetWarehouseByIdWithRoomsHandler(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehouses/:id", handlers.GetWarehouseByIdHandler(controller.WarehouseController))
	publicRoutes.Handle(http.MethodPost, "/warehouses", handlers.CreateWarehouseHandler(controller.WarehouseController))
	publicRoutes.Handle(http.MethodPut, "/warehouses", handlers.UpdateWarehouseHandler(controller.WarehouseController))

	// Room routes
	publicRoutes.Handle(http.MethodGet, "/rooms", handlers.GetRoomsHandler(controller.RoomController))
	publicRoutes.Handle(http.MethodGet, "/roomswithshelves", handlers.GetRoomsWithShelvesHandle(controller.RoomController))
	publicRoutes.Handle(http.MethodGet, "/roomswithshelves/:id", handlers.GetRoomsByIdWithShelvesHandle(controller.RoomController))
	publicRoutes.Handle(http.MethodGet, "/rooms/:id", handlers.GetRoomsByIdHandle(controller.RoomController))
	publicRoutes.Handle(http.MethodPost, "/rooms", handlers.CreateRoomHandle(controller.RoomController))
	publicRoutes.Handle(http.MethodPut, "/rooms", handlers.UpdateRoomHandle(controller.RoomController))

	// Shelve routes
	publicRoutes.Handle(http.MethodGet, "/shelves", handlers.GetShelvesHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelveswithitems", handlers.GetShelvesWithItemsHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelveswithitems/:id", handlers.GetShelveByIdWithItemsHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelves/:id", handlers.GetShelveByIdHandler(controller.ShelveController))
	//publicRoutes.Handle(http.MethodPost, "/shelves", handlers.CreateShelveHandler(controller.ShelveController))
	//publicRoutes.Handle(http.MethodPut, "/shelves", handlers.UpdateShelveHandler(controller.ShelveController))

	// Items routes
	//publicRoutes.Handle(http.MethodGet, "/items", handlers.GetItemsHandler(controller.ItemController))

	// swagger
	docs.SwaggerInfo.Title = "InventoryPro API"
	docs.SwaggerInfo.Description = "This is the API for the InventoryPro project"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
