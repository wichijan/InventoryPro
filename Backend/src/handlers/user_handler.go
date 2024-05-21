package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/controllers"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

// @Summary Register user
// @Description Register user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.RegistrationRequest true "User data"
// @Success 201 {object} models.LoginResponse
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/register [post]
func RegisterUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationData models.RegistrationRequest
		err := c.ShouldBind(&registrationData)
		if err != nil ||
			utils.ContainsEmptyString(
				registrationData.Username, registrationData.Email, registrationData.Password,
				registrationData.FirstName, registrationData.LastName,
			) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}
		// user is logged in after registration
		loginResponse, inv_err := userCtrl.RegisterUser(registrationData)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		utils.SetJWTCookies(c, loginResponse.Token, loginResponse.RefreshToken, false)
		c.JSON(http.StatusCreated, loginResponse.User)
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

func LogoutUserHandler(c *gin.Context) {
	utils.SetJWTCookies(c, "", "", true)
	c.Status(http.StatusOK)
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
		if err != nil ||
			utils.ContainsEmptyString(
				loginData.Username, loginData.Password,
			) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}
		loginResponse, kts_err := userCtrl.LoginUser(loginData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
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

		c.Status(http.StatusOK)
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
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST)
			return
		}

		err := userCtrl.CheckUsername(requestData.Username)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.Status(http.StatusOK)
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

func IsAdminHandler(c *gin.Context) {
	userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
	if !ok {
		utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
		return
	}
	adminId := uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")
	c.JSON(http.StatusOK, *userId == adminId)
}



// @Summary Get Users and their items
// @Description Get Users and their items for list view
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ItemWithUser
// @Failure 500 {object} models.INVErrorMessage
// @Router /users-items [get]
func GetUsersItemshandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		usersItems, inv_err := userCtrl.GetUsersItems()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, usersItems)
	}
}