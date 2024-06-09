package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
	"github.com/wichijan/InventoryPro/src/websocket"
)

// @Summary Register user
// @Description Register user - return "Admin has been informed"
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.RegistrationRequest true "User data"
// @Success 201
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/register [post]
func RegisterUserHandler(userCtrl controllers.UserControllerI, hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationData models.RegistrationRequest
		err := c.ShouldBind(&registrationData)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(
			registrationData.Username, registrationData.Email, registrationData.Password,
			registrationData.FirstName, registrationData.LastName,
		) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid username, email, password, first name or last name"))
			return
		}

		// user is logged in after registration
		inv_err := userCtrl.RegisterUser(registrationData)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		// inform admin
		hub.HandleMessage(websocket.Message{
			Type:     "registrationRequest",
			ForAdmin: true,
			Sender:   "server",
			Content:  "Registration Request for Admins!",
			ID:       utils.WEBSOCKET_DEFAULT_ROOM,
		})

		c.JSON(http.StatusCreated, nil)
	}
}

// @Summary
// @Description
// @Tags Users
// @Accept  json
// @Produce  json
// @Param code path string true "Registration Code"
// @Param user body models.RegistrationRequest true "User data"
// @Success 201 {object} models.LoginResponse
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/register/:code [post]
func RegisterUserWithCodeHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		if code == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Code is empty"))
			return
		}

		// Validate code
		ok, inv_err := userCtrl.ValidateRegistrationCode(&code)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		if !*ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid code"))
			return
		}

		// Register user
		var newPassword models.PasswordReset
		err := c.ShouldBind(&newPassword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(
			newPassword.Password, *newPassword.Username,
		) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid username or password"))
			return
		}

		// user is logged in after registration
		inv_err = userCtrl.UpdateUserPassword(newPassword.Username, newPassword.Password)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary
// @Description
// @Tags Users
// @Accept  json
// @Produce  json
// @Param code path string true "Registration Code"
// @Param user body models.RegistrationRequest true "User data"
// @Success 201 {object} models.LoginResponse
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/register/:code [post]
func GenerateUserRegistrationCodeHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationData models.RegistrationRequest
		err := c.ShouldBind(&registrationData)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(
			registrationData.Username, registrationData.Email,
			registrationData.FirstName, registrationData.LastName,
		) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid username, email, first name or last name"))
			return
		}
		if registrationData.Password != "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Password has to be set by user himself"))
			return
		}

		// user is logged in after registration
		code, inv_err := userCtrl.RegisterUserAndCode(registrationData)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusCreated, code)
	}
}

// @Summary Get User Data
// @Description Get User Data when logged in
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 201 {object} models.UserWithTypeName
// @Failure 400 {object} models.INVErrorMessage
// @Router /users/get-me [GET]
func GetUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		user, inv_err := userCtrl.GetUserById(userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// @Summary Accept User Registration Request
// @Description Accept User Registration Request
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID from registration request"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/accept-registration/:userId [GET]
func AcceptUserRegistrationRequestHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		if userId == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("User ID is empty"))
			return
		}

		inv_err := userCtrl.AcceptUserRegistrationRequest(&userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

func LogoutUserHandler(c *gin.Context) {
	utils.SetJWTCookies(c, "", "", true)
	c.JSON(http.StatusOK, nil)
}

// @Summary Login user
// @Description Login user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.LoginRequest true "User data"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/login [post]
func LoginUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData models.LoginRequest
		err := c.ShouldBind(&loginData)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(loginData.Username, loginData.Password) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid username or password"))
			return
		}

		loginResponse, inv_error := userCtrl.LoginUser(loginData)
		if inv_error != nil {
			utils.HandleErrorAndAbort(c, inv_error)
			return
		}

		utils.SetJWTCookies(c, loginResponse.Token, loginResponse.RefreshToken, false)
		c.JSON(http.StatusOK, loginResponse.User)
	}
}

// @Summary Check email
// @Description Check email
// @Tags Users
// @Accept  json
// @Produce  json
// @Param checkEmailRequest body models.CheckEmailRequest true "Email data"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/check-email [post]
func CheckEmailHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData models.CheckEmailRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Email is empty"))
			return
		}

		err := userCtrl.CheckEmail(requestData.Email)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Check username
// @Description Check username
// @Tags Users
// @Accept  json
// @Produce  json
// @Param checkUsernameRequest body models.CheckUsernameRequest true "Username data"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/check-username [post]
func CheckUsernameHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData models.CheckUsernameRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Username is empty"))
			return
		}

		err := userCtrl.CheckUsername(requestData.Username)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Logged in
// @Description Check if user is logged in
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.LoggedInResponse
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/logged-in [get]
func LoggedInHandler(c *gin.Context) {
	var token string

	// check if token is set
	token, err := c.Cookie("token")
	if err != nil {
		// token is not set, check if refresh token is set
		token, err = c.Cookie("refreshToken")
		if err != nil {
			c.JSON(http.StatusOK, models.LoggedInResponse{
				LoggedIn: false,
			})
			return
		}
	}

	id, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusOK, models.LoggedInResponse{
			LoggedIn: false,
		})
		return
	}

	c.JSON(http.StatusOK, models.LoggedInResponse{
		LoggedIn: true,
		Id:       id,
	})
}

// @Summary Get Registration Requests
// @Description Get Registration Requests
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.RegistrationRequests
// @Failure 400 {object} models.INVErrorMessage
// @Router /registration-requests [GET]
func GetRegistrationRequestsHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		requests, inv_err := userCtrl.GetRegistrationRequests()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, requests)
	}
}

// @Summary Reset Password
// @Description Reset Password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.PasswordReset true "User data"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /reset-password [POST]
func ResetPasswordHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Register user
		var newPassword models.PasswordReset
		err := c.ShouldBind(&newPassword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(
			newPassword.Password, *newPassword.Username,
		) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid username or password"))
			return
		}

		// user is logged in after registration
		inv_err := userCtrl.UpdateUserPassword(newPassword.Username, newPassword.Password)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
