package main

import (
	"database/sql"
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/docs"
	"github.com/gin-gonic/gin"
	"github.com/wichijan/InventoryPro/src/controllers"
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
	UserController      controllers.UserControllerI
	KeywordController   controllers.KeywordControllerI
	UserRoleController  controllers.UserRoleControllerI
	RoleController      controllers.RoleControllerI
}

func createRouter(dbConnection *sql.DB) *gin.Engine {
	router := gin.Default()

	// Attach Middleware
	router.Use(middlewares.CorsMiddleware())

	// Create api groups, with special middleware
	publicRoutes := router.Group("/")
	securedRoutes := router.Group("/", middlewares.JwtAuthMiddleware())

	// Create managers and repositories
	databaseManager := &managers.DatabaseManager{
		Connection: dbConnection,
	}

	adminRoutes := router.Group("/", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware(databaseManager))

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

	userRepo := &repositories.UserRepository{
		DatabaseManager: databaseManager,
	}

	userTypeRepo := &repositories.UserTypeRepository{
		DatabaseManager: databaseManager,
	}

	itemKeywordRepo := &repositories.ItemKeywordRepository{
		DatabaseManager: databaseManager,
	}

	keywordRepo := &repositories.KeywordRepository{
		DatabaseManager: databaseManager,
	}

	subjectRepo := &repositories.SubjectRepository{
		DatabaseManager: databaseManager,
	}

	itemSubjectRepo := &repositories.ItemSubjectRepository{
		DatabaseManager: databaseManager,
	}

	itemInShelveRepo := &repositories.ItemInShelveRepository{
		DatabaseManager: databaseManager,
	}

	itemStatusRepo := &repositories.ItemStatusRepository{
		DatabaseManager: databaseManager,
	}

	userItemRepo := &repositories.UserItemRepository{
		DatabaseManager: databaseManager,
	}

	userRoleRepo := &repositories.UserRoleRepository{
		DatabaseManager: databaseManager,
	}

	roleRepo := &repositories.RoleRepository{
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
			ItemRepo:         itemRepo,
			ItemInShelveRepo: itemInShelveRepo,
			ItemStatusRepo:   itemStatusRepo,
			UserItemRepo:     userItemRepo,
			KeywordRepo:      keywordRepo,
			SubjectRepo:      subjectRepo,
			ItemKeywordRepo:  itemKeywordRepo,
			ItemSubjectRepo:  itemSubjectRepo,
		},
		UserController: &controllers.UserController{
			UserRepo:     userRepo,
			UserTypeRepo: userTypeRepo,
		},
		UserRoleController: &controllers.UserRoleController{
			UserRoleRepo: userRoleRepo,
		},
		RoleController: &controllers.RoleController{
			RoleRepo: roleRepo,
		},
	}

	// user routes
	publicRoutes.Handle(http.MethodPost, "/auth/register", handlers.RegisterUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/login", handlers.LoginUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/logout", handlers.LogoutUserHandler)
	publicRoutes.Handle(http.MethodPost, "/auth/check-email", handlers.CheckEmailHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-username", handlers.CheckUsernameHandler(controller.UserController))
	publicRoutes.Handle(http.MethodGet, "/auth/logged-in", handlers.LoggedInHandler)
	securedRoutes.Handle(http.MethodGet, "/auth/is-admin", handlers.IsAdminHandler)

	securedRoutes.Handle(http.MethodGet, "/users/get-me", handlers.GetUserHandler(controller.UserController))

	// Warehouse routes
	publicRoutes.Handle(http.MethodGet, "/warehouses", handlers.GetWarehousesHandler(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehousesWithRooms", handlers.GetWarehousesWithRoomsHandler(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehousesWithRooms/:id", handlers.GetWarehouseByIdWithRoomsHandler(controller.WarehouseController))
	publicRoutes.Handle(http.MethodGet, "/warehouses/:id", handlers.GetWarehouseByIdHandler(controller.WarehouseController))
	adminRoutes.Handle(http.MethodPost, "/warehouses", handlers.CreateWarehouseHandler(controller.WarehouseController))
	adminRoutes.Handle(http.MethodPut, "/warehouses", handlers.UpdateWarehouseHandler(controller.WarehouseController))

	// Room routes
	publicRoutes.Handle(http.MethodGet, "/rooms", handlers.GetRoomsHandler(controller.RoomController))
	publicRoutes.Handle(http.MethodGet, "/roomswithshelves", handlers.GetRoomsWithShelvesHandle(controller.RoomController))
	publicRoutes.Handle(http.MethodGet, "/roomswithshelves/:id", handlers.GetRoomsByIdWithShelvesHandle(controller.RoomController))
	publicRoutes.Handle(http.MethodGet, "/rooms/:id", handlers.GetRoomsByIdHandle(controller.RoomController))
	adminRoutes.Handle(http.MethodPost, "/rooms", handlers.CreateRoomHandle(controller.RoomController))
	adminRoutes.Handle(http.MethodPut, "/rooms", handlers.UpdateRoomHandle(controller.RoomController))

	// Shelve routes
	publicRoutes.Handle(http.MethodGet, "/shelves", handlers.GetShelvesHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelveswithitems", handlers.GetShelvesWithItemsHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelveswithitems/:id", handlers.GetShelveByIdWithItemsHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelves/:id", handlers.GetShelveByIdHandler(controller.ShelveController))
	//publicRoutes.Handle(http.MethodPost, "/shelves", handlers.CreateShelveHandler(controller.ShelveController))
	//publicRoutes.Handle(http.MethodPut, "/shelves", handlers.UpdateShelveHandler(controller.ShelveController))

	// Items routes
	publicRoutes.Handle(http.MethodGet, "/items", handlers.GetItemsHandler(controller.ItemController))
	publicRoutes.Handle(http.MethodGet, "/items/:id", handlers.GetItemByIdHandler(controller.ItemController))
	publicRoutes.Handle(http.MethodPost, "/items", handlers.CreateItemHandler(controller.ItemController))
	publicRoutes.Handle(http.MethodPut, "/items", handlers.UpdateItemHandler(controller.ItemController))
	publicRoutes.Handle(http.MethodPost, "/items/addkeyword", handlers.AddKeywordToItemHandler(controller.ItemController))
	publicRoutes.Handle(http.MethodPost, "/items/removekeyword", handlers.RemoveKeywordFromItemHandler(controller.ItemController))
	publicRoutes.Handle(http.MethodPost, "/items/addsubject", handlers.AddSubjectToItemHandler(controller.ItemController))
	publicRoutes.Handle(http.MethodPost, "/items/removesubject", handlers.RemoveSubjectFromItemHandler(controller.ItemController))

	// Item reserve
	securedRoutes.Handle(http.MethodPost, "/items/reserve", handlers.ReserveItemHandler(controller.ItemController))
	securedRoutes.Handle(http.MethodDelete, "/items/reserve-cancel/:id", handlers.CancelReserveItemHandler(controller.ItemController))

	// Item move
	securedRoutes.Handle(http.MethodPost, "/items/borrow", nil)
	securedRoutes.Handle(http.MethodPost, "/items/return", nil)

	// Roles routes
	adminRoutes.Handle(http.MethodGet, "/roles", handlers.GetRolesHandler(controller.RoleController))
	adminRoutes.Handle(http.MethodPost, "/roles", handlers.CreateRoleHandler(controller.RoleController))
	adminRoutes.Handle(http.MethodPut, "/roles", handlers.UpdateRoleHandler(controller.RoleController))

	// User roles routes
	adminRoutes.Handle(http.MethodPost, "/user-roles/add-role", handlers.AddRoleToUserHandler(controller.UserRoleController))
	adminRoutes.Handle(http.MethodDelete, "/user-roles/remove-role", handlers.RemoveRoleFromUserHandler(controller.UserRoleController))

	// swagger
	docs.SwaggerInfo.Title = "InventoryPro API"
	docs.SwaggerInfo.Description = "This is the API for the InventoryPro project"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
