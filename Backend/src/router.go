package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/wichijan/InventoryPro/src/controllers"
	"github.com/wichijan/InventoryPro/src/docs"
	"github.com/wichijan/InventoryPro/src/handlers"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/middlewares"
	"github.com/wichijan/InventoryPro/src/repositories"
	"github.com/wichijan/InventoryPro/src/websocket"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controllers struct {
	WarehouseController       controllers.WarehouseControllerI
	RoomController            controllers.RoomControllerI
	ShelveController          controllers.ShelveControllerI
	ItemController            controllers.ItemControllerI
	UserController            controllers.UserControllerI
	KeywordController         controllers.KeywordControllerI
	UserRoleController        controllers.UserRoleControllerI
	RoleController            controllers.RoleControllerI
	SubjectController         controllers.SubjectControllerI
	UserTypeController        controllers.UserTypeControllerI
	ReservationController     controllers.ReservationControllerI
	ItemsQuickShelfController controllers.ItemQuickShelfControllerI
	QuickShelfController      controllers.QuickShelfControllerI
}

func createRouter(dbConnection *sql.DB, hub *websocket.Hub, mgInstance *mailgun.MailgunImpl) *gin.Engine {

	mailMgr := &managers.MailManager{
		MailgunInstance: mgInstance,
	}

	// Create managers and repositories
	databaseManager := &managers.DatabaseManager{
		Connection: dbConnection,
	}

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

	itemQuickShelfRepo := &repositories.ItemQuickShelfRepository{
		DatabaseManagerI: databaseManager,
	}

	quickShelfRepo := &repositories.QuickShelfRepository{
		DatabaseManagerI: databaseManager,
	}

	transactionRepo := &repositories.TransactionRepository{
		DatabaseManagerI: databaseManager,
	}

	registrationRequestRepo := &repositories.RegistrationRequestRepository{
		DatabaseManagerI: databaseManager,
	}

	registrationCodeRepo := &repositories.RegistrationCodeRepository{
		DatabaseManagerI: databaseManager,
	}

	transferRequestRepo := &repositories.TransferRequestRepository{
		DatabaseManagerI: databaseManager,
	}

	bookRepo := &repositories.BookRepository{
		DatabaseManagerI: databaseManager,
	}

	singleObjectRepo := &repositories.SingleObjectRepository{
		DatabaseManagerI: databaseManager,
	}

	setOfObjectsRepo := &repositories.SetsOfObjectsRepository{
		DatabaseManagerI: databaseManager,
	}

	// Create controllers
	controller := Controllers{
		WarehouseController: &controllers.WarehouseController{
			WarehouseRepo: warehouseRepo,
			RoomRepo:      roomRepo,
		},
		RoomController: &controllers.RoomController{
			RoomRepo:        roomRepo,
			ShelveRepo:      shelveRepo,
			QuickShelveRepo: quickShelfRepo,
		},
		ShelveController: &controllers.ShelveController{
			ShelveRepo: shelveRepo,
		},
		ItemController: &controllers.ItemController{
			ItemRepo:            itemRepo,
			ItemInShelveRepo:    itemInShelveRepo,
			UserItemRepo:        userItemRepo,
			KeywordRepo:         keywordRepo,
			SubjectRepo:         subjectRepo,
			ItemKeywordRepo:     itemKeywordRepo,
			ItemSubjectRepo:     itemSubjectRepo,
			ShelveRepo:          shelveRepo,
			TransactionRepo:     transactionRepo,
			TransferRequestRepo: transferRequestRepo,

			BookRepo:            bookRepo,
			SingleObjectRepo:    singleObjectRepo,
			SetOfObjectsRepo:    setOfObjectsRepo,
			ReservationRepo:     reservationRepo,
			ItemsQuickShelfRepo: itemQuickShelfRepo,
		},
		UserController: &controllers.UserController{
			UserRepo:                userRepo,
			UserTypeRepo:            userTypeRepo,
			RegistrationRequestRepo: registrationRequestRepo,
			RegistrationCodeRepo:    registrationCodeRepo,
			RoleRepo:                roleRepo,
			MailMgr:                 mailMgr,
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
			TransactionRepo: transactionRepo,
			UserRepo:        userRepo,
		},
		ItemsQuickShelfController: &controllers.ItemQuickShelfController{
			ItemQuickShelfRepo: itemQuickShelfRepo,
			UserItemRepo:       userItemRepo,
			ItemRepo:           itemRepo,
			ItemsInShelfRepo:   itemInShelveRepo,
		},
		QuickShelfController: &controllers.QuickShelfController{
			QuickShelfRepo:     quickShelfRepo,
			ItemQuickShelfRepo: itemQuickShelfRepo,
		},
	}

	router := gin.Default()

	// Attach Middleware
	router.Use(middlewares.CorsMiddleware())

	// Create api groups, with special middleware
	publicRoutes := router.Group("/")
	securedRoutes := router.Group("/", middlewares.JwtAuthMiddleware())
	adminRoutes := router.Group("/", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware(controller.UserController))

	// user routes
	publicRoutes.Handle(http.MethodPost, "/auth/register", handlers.RegisterUserHandler(controller.UserController, hub))
	adminRoutes.Handle(http.MethodPost, "/auth/accept-registration/:userId", handlers.AcceptUserRegistrationRequestHandler(controller.UserController))
	adminRoutes.Handle(http.MethodDelete, "/auth/decline-registration/:userId", handlers.DeclineUserRegistrationRequestHandler(controller.UserController))
	adminRoutes.Handle(http.MethodPost, "/auth/generate-code", handlers.GenerateUserRegistrationCodeHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/register/:code", handlers.RegisterUserWithCodeHandler(controller.UserController))
	
	securedRoutes.Handle(http.MethodPost, "/auth/reset-password", handlers.ResetPasswordHandler(controller.UserController))
	securedRoutes.Handle(http.MethodPost, "/auth/request-forgot-password", handlers.RequestForgetPasswordHandler(controller.UserController))
	securedRoutes.Handle(http.MethodPost, "/auth/forget-password", handlers.ForgetPasswordHandler(controller.UserController))

	publicRoutes.Handle(http.MethodPost, "/auth/login", handlers.LoginUserHandler(controller.UserController))
	securedRoutes.Handle(http.MethodPost, "/auth/logout", handlers.LogoutUserHandler)
	publicRoutes.Handle(http.MethodPost, "/auth/check-email", handlers.CheckEmailHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-username", handlers.CheckUsernameHandler(controller.UserController))
	publicRoutes.Handle(http.MethodGet, "/auth/logged-in", handlers.LoggedInHandler)
	securedRoutes.Handle(http.MethodGet, "/auth/is-Admin", handlers.IsAdmin(controller.UserController))

	adminRoutes.Handle(http.MethodGet, "/registration-requests", handlers.GetRegistrationRequestsHandler(controller.UserController))

	securedRoutes.Handle(http.MethodGet, "/users/get-me", handlers.GetUserHandler(controller.UserController))
	securedRoutes.Handle(http.MethodGet, "/users", handlers.GetUsersHandler(controller.UserController))
	securedRoutes.Handle(http.MethodGet, "/users/:id", handlers.GetUserByIdHandler(controller.UserController))
	// Picture for Users
	securedRoutes.Handle(http.MethodPost, "/users-picture", handlers.UploadImageForUserHandler(controller.UserController))
	securedRoutes.Handle(http.MethodGet, "/users-picture", handlers.GetImagePathForUserHandler(controller.UserController))
	securedRoutes.Handle(http.MethodDelete, "/users-picture", handlers.RemoveImageForUserHandler(controller.UserController))

	// Warehouse routes
	securedRoutes.Handle(http.MethodGet, "/warehouses", handlers.GetWarehousesHandler(controller.WarehouseController))
	securedRoutes.Handle(http.MethodGet, "/warehouses-with-rooms", handlers.GetWarehousesWithRoomsHandler(controller.WarehouseController))
	securedRoutes.Handle(http.MethodGet, "/warehouses-with-rooms/:id", handlers.GetWarehouseByIdWithRoomsHandler(controller.WarehouseController))
	securedRoutes.Handle(http.MethodGet, "/warehouses/:id", handlers.GetWarehouseByIdHandler(controller.WarehouseController))
	adminRoutes.Handle(http.MethodPost, "/warehouses", handlers.CreateWarehouseHandler(controller.WarehouseController))
	adminRoutes.Handle(http.MethodPut, "/warehouses", handlers.UpdateWarehouseHandler(controller.WarehouseController))
	adminRoutes.Handle(http.MethodDelete, "/warehouses/:id", handlers.DeleteWarehouse(controller.WarehouseController))

	// Room routes
	securedRoutes.Handle(http.MethodGet, "/rooms", handlers.GetRoomsHandler(controller.RoomController))
	securedRoutes.Handle(http.MethodGet, "/rooms-with-shelves", handlers.GetRoomsWithShelvesHandle(controller.RoomController))
	securedRoutes.Handle(http.MethodGet, "/rooms-with-shelves/:id", handlers.GetRoomsByIdWithShelvesHandle(controller.RoomController))
	securedRoutes.Handle(http.MethodGet, "/rooms/:id", handlers.GetRoomsByIdHandle(controller.RoomController))
	adminRoutes.Handle(http.MethodPost, "/rooms", handlers.CreateRoomHandle(controller.RoomController))
	adminRoutes.Handle(http.MethodPut, "/rooms", handlers.UpdateRoomHandle(controller.RoomController))
	adminRoutes.Handle(http.MethodDelete, "/rooms/:id", handlers.DeleteRoomHandle(controller.RoomController))

	// Shelve routes
	securedRoutes.Handle(http.MethodGet, "/shelves", handlers.GetShelvesHandler(controller.ShelveController))
	securedRoutes.Handle(http.MethodGet, "/shelves-with-items", handlers.GetShelvesWithItemsHandler(controller.ShelveController))
	securedRoutes.Handle(http.MethodGet, "/shelves-with-items/:id", handlers.GetShelveByIdWithItemsHandler(controller.ShelveController))
	securedRoutes.Handle(http.MethodGet, "/shelves/:id", handlers.GetShelveByIdHandler(controller.ShelveController))
	adminRoutes.Handle(http.MethodPost, "/shelves", handlers.CreateShelveHandler(controller.ShelveController))
	adminRoutes.Handle(http.MethodPut, "/shelves", handlers.UpdateShelveHandler(controller.ShelveController))
	adminRoutes.Handle(http.MethodDelete, "/shelves/:id", handlers.DeleteShelveHandler(controller.ShelveController))

	// Quick Shelf routes
	securedRoutes.Handle(http.MethodGet, "/quick-shelves", handlers.GetQuickShelvesHandler(controller.QuickShelfController))
	securedRoutes.Handle(http.MethodGet, "/quick-shelves/:id", handlers.GetQuickShelfByIdHandler(controller.QuickShelfController))
	adminRoutes.Handle(http.MethodPost, "/quick-shelves", handlers.CreateQuickShelfHandler(controller.QuickShelfController))
	adminRoutes.Handle(http.MethodPut, "/quick-shelves", handlers.UpdateQuickShelfHandler(controller.QuickShelfController))
	adminRoutes.Handle(http.MethodDelete, "/quick-shelves/:id", handlers.DeleteQuickShelfHandler(controller.QuickShelfController))

	// Items routes
	securedRoutes.Handle(http.MethodGet, "/items", handlers.GetItemsHandler(controller.ItemController))
	securedRoutes.Handle(http.MethodGet, "/items/:id", handlers.GetItemByIdHandler(controller.ItemController))
	securedRoutes.Handle(http.MethodGet, "/items/book/:id", handlers.GetBookByIdHandler(controller.ItemController))
	securedRoutes.Handle(http.MethodGet, "/items/single-object/:id", handlers.GetSingleObjectByIdHandler(controller.ItemController))
	securedRoutes.Handle(http.MethodGet, "/items/set-of-objects/:id", handlers.GetSetOfObjectsByIdHandler(controller.ItemController))
	// Creates item - book or single object or set of objects
	adminRoutes.Handle(http.MethodPost, "/items/book", handlers.CreateItemWithBookHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodPost, "/items/single-object", handlers.CreateItemWithSingleObjectHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodPost, "/items/set-of-objects", handlers.CreateItemWithSetOfObjectsHandler(controller.ItemController))
	// Updates item - book or single object or set of objects
	adminRoutes.Handle(http.MethodPut, "/items/book", handlers.UpdateItemWithBookHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodPut, "/items/single-object", handlers.UpdateItemWithSingleObjectHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodPut, "/items/set-of-objects", handlers.UpdateItemWithSetOfObjectsHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodDelete, "/items/:id", handlers.DeleteItemHandler(controller.ItemController))
	// Picture for item
	adminRoutes.Handle(http.MethodPost, "/items-picture", handlers.UploadImageForItemHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodGet, "/items-picture/:id", handlers.GetImagePathForItemHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodDelete, "/items-picture/:id", handlers.RemoveImageForItemHandler(controller.ItemController))
	// Keyword for item
	adminRoutes.Handle(http.MethodPost, "/items/add-keyword", handlers.AddKeywordToItemHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodDelete, "/items/remove-keyword", handlers.RemoveKeywordFromItemHandler(controller.ItemController))
	// Subject for item
	adminRoutes.Handle(http.MethodPost, "/items/add-subject", handlers.AddSubjectToItemHandler(controller.ItemController))
	adminRoutes.Handle(http.MethodDelete, "/items/remove-subject", handlers.RemoveSubjectFromItemHandler(controller.ItemController))
	// Item reserve
	securedRoutes.Handle(http.MethodPost, "/items/reserve", handlers.ReserveItemHandler(controller.ReservationController))
	securedRoutes.Handle(http.MethodDelete, "/items/reserve-cancel/:id", handlers.CancelReserveItemHandler(controller.ReservationController))
	// Item move
	securedRoutes.Handle(http.MethodPost, "/items/borrow", handlers.BorrowItemHandler(controller.ItemController))
	securedRoutes.Handle(http.MethodDelete, "/items/return/:id", handlers.ReturnReserveItemHandler(controller.ItemController))
	// Item quick shelf
	securedRoutes.Handle(http.MethodPost, "/items/add-item-to-quick-shelf", handlers.AddToQuickShelfHandler(controller.ItemsQuickShelfController))
	securedRoutes.Handle(http.MethodDelete, "/items/remove-item-to-quick-shelf", handlers.RemoveItemFromQuickShelfHandler(controller.ItemsQuickShelfController))
	securedRoutes.Handle(http.MethodDelete, "/items/clear-quick-shelf/:id", handlers.ClearQuickShelfHandler(controller.ItemsQuickShelfController))
	securedRoutes.Handle(http.MethodGet, "/items/quick-shelf/:id", handlers.GetItemsInQuickShelfHandler(controller.ItemsQuickShelfController))
	// Item Move - From User A -> User B
	securedRoutes.Handle(http.MethodPost, "/items/transfer-request", handlers.MoveItemRequestHandler(controller.ItemController, hub))
	securedRoutes.Handle(http.MethodPost, "/items/transfer-accept/:id", handlers.MoveItemAcceptedHandler(controller.ItemController, hub))
	securedRoutes.Handle(http.MethodGet, "/items/transfer-requests", handlers.GetTransferRequestByIdHandler(controller.ItemController))

	// Subject Routes
	securedRoutes.Handle(http.MethodGet, "/subjects", handlers.GetSubjectsHandler(controller.SubjectController))
	adminRoutes.Handle(http.MethodPost, "/subjects", handlers.CreateSubjectHandler(controller.SubjectController))
	adminRoutes.Handle(http.MethodPut, "/subjects", handlers.UpdateSubjectHandler(controller.SubjectController))
	adminRoutes.Handle(http.MethodDelete, "/subjects/:id", handlers.DeleteSubjectHandler(controller.SubjectController))

	// Keyword routes
	securedRoutes.Handle(http.MethodGet, "/keywords", handlers.GetKeywordsHandler(controller.KeywordController))
	adminRoutes.Handle(http.MethodPost, "/keywords", handlers.CreateKeywordHandler(controller.KeywordController))
	adminRoutes.Handle(http.MethodPut, "/keywords", handlers.UpdateKeywordHandler(controller.KeywordController))
	adminRoutes.Handle(http.MethodDelete, "/keywords/:id", handlers.DeleteKeywordHandler(controller.KeywordController))

	// Roles routes
	adminRoutes.Handle(http.MethodGet, "/roles", handlers.GetRolesHandler(controller.RoleController))
	adminRoutes.Handle(http.MethodPost, "/roles", handlers.CreateRoleHandler(controller.RoleController))
	adminRoutes.Handle(http.MethodPut, "/roles", handlers.UpdateRoleHandler(controller.RoleController))
	adminRoutes.Handle(http.MethodDelete, "/roles", handlers.DeleteRoleHandler(controller.RoleController))

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

	securedRoutes.Handle(http.MethodGet, "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	securedRoutes.Handle(http.MethodGet, "/ws", handlers.WebsocketHandler(databaseManager, hub))

	return router
}
