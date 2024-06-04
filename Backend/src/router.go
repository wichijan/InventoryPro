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
	WarehouseController   controllers.WarehouseControllerI
	RoomController        controllers.RoomControllerI
	ShelveController      controllers.ShelveControllerI
	ItemController        controllers.ItemControllerI
	UserController        controllers.UserControllerI
	KeywordController     controllers.KeywordControllerI
	UserRoleController    controllers.UserRoleControllerI
	RoleController        controllers.RoleControllerI
	SubjectController     controllers.SubjectControllerI
	UserTypeController    controllers.UserTypeControllerI
	ReservationController controllers.ReservationControllerI
	QuickShelfController  controllers.ItemQuickShelfControllerI
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
		DatabaseManagerI: databaseManager,
	}

	roomRepo := &repositories.RoomRepository{
		DatabaseManagerI: databaseManager,
	}

	shelveRepo := &repositories.ShelveRepository{
		DatabaseManagerI: databaseManager,
	}

	itemRepo := &repositories.ItemRepository{
		DatabaseManagerI: databaseManager,
	}

	userRepo := &repositories.UserRepository{
		DatabaseManagerI: databaseManager,
	}

	userTypeRepo := &repositories.UserTypeRepository{
		DatabaseManagerI: databaseManager,
	}

	itemKeywordRepo := &repositories.ItemKeywordRepository{
		DatabaseManagerI: databaseManager,
	}

	keywordRepo := &repositories.KeywordRepository{
		DatabaseManagerI: databaseManager,
	}

	subjectRepo := &repositories.SubjectRepository{
		DatabaseManagerI: databaseManager,
	}

	itemSubjectRepo := &repositories.ItemSubjectRepository{
		DatabaseManagerI: databaseManager,
	}

	itemInShelveRepo := &repositories.ItemInShelveRepository{
		DatabaseManagerI: databaseManager,
	}

	userItemRepo := &repositories.UserItemRepository{
		DatabaseManagerI: databaseManager,
	}

	userRoleRepo := &repositories.UserRoleRepository{
		DatabaseManagerI: databaseManager,
	}

	roleRepo := &repositories.RoleRepository{
		DatabaseManagerI: databaseManager,
	}

	reservationRepo := &repositories.ReservationRepository{
		DatabaseManagerI: databaseManager,
	}

	quickShelfRepo := &repositories.ItemQuickShelfRepository{
		DatabaseManagerI: databaseManager,
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
		SubjectController: &controllers.SubjectController{
			SubjectRepo: subjectRepo,
		},
		KeywordController: &controllers.KeywordController{
			KeywordRepo: keywordRepo,
		},
		UserTypeController: &controllers.UserTypeController{
			UserTypeRepo: userTypeRepo,
		},
		ReservationController: &controllers.ReservationController{
			ReservationRepo: reservationRepo,
		},
		QuickShelfController : &controllers.ItemQuickShelfController{
			ItemQuickShelfRepo: quickShelfRepo,
			UserItemRepo: userItemRepo,
		},
	}

	// user routes
	// TODO Add route to request registration
	// TODO Add route to accept registration request 
	// TODO Add route to reset password
	// TODO Add route to upload Picture
	publicRoutes.Handle(http.MethodPost, "/auth/register", handlers.RegisterUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/login", handlers.LoginUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/logout", handlers.LogoutUserHandler)
	publicRoutes.Handle(http.MethodPost, "/auth/check-email", handlers.CheckEmailHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-username", handlers.CheckUsernameHandler(controller.UserController))
	publicRoutes.Handle(http.MethodGet, "/auth/logged-in", handlers.LoggedInHandler)

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
	publicRoutes.Handle(http.MethodGet, "/rooms-with-shelves", handlers.GetRoomsWithShelvesHandle(controller.RoomController))
	publicRoutes.Handle(http.MethodGet, "/rooms-with-shelves/:id", handlers.GetRoomsByIdWithShelvesHandle(controller.RoomController))
	publicRoutes.Handle(http.MethodGet, "/rooms/:id", handlers.GetRoomsByIdHandle(controller.RoomController))
	adminRoutes.Handle(http.MethodPost, "/rooms", handlers.CreateRoomHandle(controller.RoomController))
	adminRoutes.Handle(http.MethodPut, "/rooms", handlers.UpdateRoomHandle(controller.RoomController))

	// Shelve routes
	publicRoutes.Handle(http.MethodGet, "/shelves", handlers.GetShelvesHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelves-with-items", handlers.GetShelvesWithItemsHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelves-with-items/:id", handlers.GetShelveByIdWithItemsHandler(controller.ShelveController))
	publicRoutes.Handle(http.MethodGet, "/shelves/:id", handlers.GetShelveByIdHandler(controller.ShelveController))
	 adminRoutes.Handle(http.MethodPost, "/shelves", handlers.CreateShelveHandler(controller.ShelveController))
	 adminRoutes.Handle(http.MethodPut, "/shelves", handlers.UpdateShelveHandler(controller.ShelveController))

	// Items routes
	publicRoutes.Handle(http.MethodGet, "/items", handlers.GetItemsHandler(controller.ItemController))
	publicRoutes.Handle(http.MethodGet, "/items/:id", handlers.GetItemByIdHandler(controller.ItemController))
	 adminRoutes.Handle(http.MethodPost, "/items", handlers.CreateItemHandler(controller.ItemController))
	 adminRoutes.Handle(http.MethodPut, "/items", handlers.UpdateItemHandler(controller.ItemController))
	// Picture for item
	 adminRoutes.Handle(http.MethodPost, "/items-picture", handlers.UploadImageForItemHandler(controller.ItemController))
	 adminRoutes.Handle(http.MethodGet, "/items-picture/:id", handlers.GetImagePathForItemHandler(controller.ItemController))
	 adminRoutes.Handle(http.MethodDelete, "/items-picture/:id", handlers.RemoveImageForItemHandler(controller.ItemController))
	// Keyword for item
	 adminRoutes.Handle(http.MethodPost, "/items/add-keyword", handlers.AddKeywordToItemHandler(controller.ItemController))
	 adminRoutes.Handle(http.MethodPost, "/items/remove-keyword", handlers.RemoveKeywordFromItemHandler(controller.ItemController))
	// Subject for item
	 adminRoutes.Handle(http.MethodPost, "/items/add-subject", handlers.AddSubjectToItemHandler(controller.ItemController))
	 adminRoutes.Handle(http.MethodPost, "/items/remove-subject", handlers.RemoveSubjectFromItemHandler(controller.ItemController))
	// Item reserve
	securedRoutes.Handle(http.MethodPost, "/items/reserve", handlers.ReserveItemHandler(controller.ReservationController))
	securedRoutes.Handle(http.MethodDelete, "/items/reserve-cancel/:id", handlers.CancelReserveItemHandler(controller.ReservationController))
	// Item move
	securedRoutes.Handle(http.MethodPost, "/items/borrow", handlers.BorrowItemHandler(controller.ItemController))
	securedRoutes.Handle(http.MethodPost, "/items/return/:id", handlers.ReturnReserveItemHandler(controller.ItemController))
	// Item quick shelf
	securedRoutes.Handle(http.MethodPost, "/items/add-item-to-quick-shelf", handlers.AddToQuickShelfHandler(controller.QuickShelfController))
	securedRoutes.Handle(http.MethodPost, "/items/remove-item-to-quick-shelf", handlers.RemoveItemFromQuickShelfHandler(controller.QuickShelfController))
	securedRoutes.Handle(http.MethodDelete, "/items/clear-quick-shelf/:id", handlers.ClearQuickShelfHandler(controller.QuickShelfController))
	securedRoutes.Handle(http.MethodGet, "/items/quick-shelf/:id", handlers.GetItemsInQuickShelfHandler(controller.QuickShelfController))

	// Subject Routes
	publicRoutes.Handle(http.MethodGet, "/subjects", handlers.GetSubjectsHandler(controller.SubjectController))
	 adminRoutes.Handle(http.MethodPost, "/subjects", handlers.CreateSubjectHandler(controller.SubjectController))
	 adminRoutes.Handle(http.MethodPut, "/subjects", handlers.UpdateSubjectHandler(controller.SubjectController))
	 adminRoutes.Handle(http.MethodDelete, "/subjects/:id", handlers.DeleteSubjectHandler(controller.SubjectController))

	// Keyword routes
	publicRoutes.Handle(http.MethodGet, "/keywords", handlers.GetKeywordsHandler(controller.KeywordController))
	 adminRoutes.Handle(http.MethodPost, "/keywords", handlers.CreateKeywordHandler(controller.KeywordController))
	 adminRoutes.Handle(http.MethodPut, "/keywords", handlers.UpdateKeywordHandler(controller.KeywordController))
	 adminRoutes.Handle(http.MethodDelete, "/keywords/:id", handlers.DeleteKeywordHandler(controller.KeywordController))

	// Roles routes
	adminRoutes.Handle(http.MethodGet, "/roles", handlers.GetRolesHandler(controller.RoleController))
	adminRoutes.Handle(http.MethodPost, "/roles", handlers.CreateRoleHandler(controller.RoleController))
	adminRoutes.Handle(http.MethodPut, "/roles", handlers.UpdateRoleHandler(controller.RoleController))

	// User roles routes
	adminRoutes.Handle(http.MethodPost, "/user-roles/add-role", handlers.AddRoleToUserHandler(controller.UserRoleController))
	adminRoutes.Handle(http.MethodDelete, "/user-roles/remove-role", handlers.RemoveRoleFromUserHandler(controller.UserRoleController))

	// User type routes
	publicRoutes.Handle(http.MethodGet, "/user-types", handlers.GetUserTypesHandler(controller.UserTypeController))
	 adminRoutes.Handle(http.MethodPost, "/user-types", handlers.CreateUserTypeHandler(controller.UserTypeController))
	 adminRoutes.Handle(http.MethodPut, "/user-types", handlers.UpdateUserTypeHandler(controller.UserTypeController))
	 adminRoutes.Handle(http.MethodDelete, "/user-types/:id", handlers.DeleteUserTypeHandler(controller.UserTypeController))

	// swagger
	docs.SwaggerInfo.Title = "InventoryPro API"
	docs.SwaggerInfo.Description = "This is the API for the InventoryPro project"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
