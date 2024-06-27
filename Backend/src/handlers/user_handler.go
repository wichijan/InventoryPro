package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		amountOfUsersReceived := hub.HandleMessage(websocket.Message{
			Type:         utils.MESSAGE_TYPE_TO_ADMINS,
			SentToUserId: "",
			Sender:       "server",
			Content:      "New user registration request: " + registrationData.Username,
		})
		if amountOfUsersReceived == nil || *amountOfUsersReceived == 0 {
			// send emails to admins
			inv_err = userCtrl.SendEmailToAdmins(registrationData.Username)
			if inv_err != nil {
				utils.HandleErrorAndAbort(c, inv_err)
				return
			}
		}

		c.JSON(http.StatusCreated, nil)
	}
}

// @Summary Register user with code
// @Description Register user with code
// @Tags Users
// @Accept  json
// @Produce  json
// @Param code path string true "Registration Code"
// @Param user body models.RegistrationCode true "Registration code"
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
		userId, inv_err := userCtrl.ValidateRegistrationCode(&code)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		// Register user
		var newPassword models.RegistrationCode
		err := c.ShouldBind(&newPassword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}

		// user is logged in after registration
		inv_err = userCtrl.UpdateUserPassword(userId, newPassword.Password)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		// Delete code
		inv_err = userCtrl.DeleteRegistrationCode(&code)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Generate User Registration Code
// @Description Generate User Registration Code - User gets email with code and link to website
// @Tags Users
// @Accept  json
// @Produce  json
// @Param code path string true "Registration Code"
// @Param user body models.RegistrationRequest true "User data"
// @Success 201 {object} string
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/generate-code [post]
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
// @Success 200 {object} models.UserWithTypeName
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

// @Summary Get Users
// @Description Get Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Users
// @Failure 400 {object} models.INVErrorMessage
// @Router /users [GET]
func GetUsersHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, inv_err := userCtrl.GetUsers()
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// @Summary Get User By Id
// @Description Get User By Id
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Users
// @Failure 400 {object} models.INVErrorMessage
// @Router /users/{id} [GET]
func GetUserByIdHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid user ID"))
			return
		}

		user, inv_err := userCtrl.GetUserById(&userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// @Summary Update User
// @Description Update User
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.UserWithoutRolesODT true "User data"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /users [PUT]
func UpdateUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		var newUser models.UserWithoutRolesODT
		err_convert := c.ShouldBind(&newUser)
		if err_convert != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}

		updateUser := models.UserWithoutRoles{
			ID:           userId,
			Email:        newUser.Email,
			FirstName:    newUser.FirstName,
			LastName:     newUser.LastName,
			JobTitle:     newUser.JobTitle,
			PhoneNumber:  newUser.PhoneNumber,
			UserTypeName: newUser.UserTypeName,
		}

		inv_err := userCtrl.UpdateUser(updateUser)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Update User as Admin
// @Description Update User as Admin
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.UserWithoutRoles true "User data"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /users/admin [PUT]
func UpdateUserAsAdminHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.UserWithoutRoles
		err_convert := c.ShouldBind(&newUser)
		if err_convert != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}

		inv_err := userCtrl.UpdateUser(newUser)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		c.JSON(http.StatusOK, nil)
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
// @Router /auth/accept-registration/:userId [POST]
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

// @Summary Decline User Registration Request
// @Description Decline User Registration Request
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID from registration request"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/decline-registration/:userId [Delete]
func DeclineUserRegistrationRequestHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		if userId == "" {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("User ID is empty"))
			return
		}

		inv_err := userCtrl.DeclineUserRegistrationRequest(&userId)
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

// @Summary Is Admin Check
// @Description Is Admin
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /auth/is-admin [get]
func IsAdmin(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		isAdmin, inv_err := userCtrl.IsAdmin(userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, isAdmin)
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
// @Router /auth/reset-password [POST]
func ResetPasswordHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
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
			newPassword.Password,
		) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid username or password"))
			return
		}

		// user is logged in after registration
		inv_err := userCtrl.UpdateUserPassword(userId, newPassword.Password)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Forget Password
// @Description Forget Password => Reset | send email to user with link for reset password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.Username true "Username"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /email-forget-password [POST]
func EmailForgetPasswordHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Register user
		var username models.Username
		err := c.ShouldBind(&username)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(
			username.Username,
		) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid username or password"))
			return
		}

		inv_err := userCtrl.ForgotPassword(username.Username)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Request Password Reset
// @Description Request Password Reset in Database. UserId should be in URL of Frontend
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.PasswordResetEmail true "User data"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Router /request-forgot-password [POST]
func RequestForgetPasswordHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Register user
		var newPassword models.PasswordResetEmail
		err := c.ShouldBind(&newPassword)
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid request body"))
			return
		}
		if utils.ContainsEmptyString(
			newPassword.Password,
		) {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Invalid username or password"))
			return
		}

		// user is logged in after registration
		inv_err := userCtrl.UpdateUserPassword(newPassword.UserId, newPassword.Password)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Upload Img for user
// @Description Upload Img for user. Form with enctype="multipart/form-data" <input type="file" name="file" /> & <input type="hidden" name="id" /> for item id
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /users-picture [post]
func UploadImageForUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		// single file
		form, err := c.MultipartForm()
		if err != nil {
			utils.HandleErrorAndAbort(c, inv_errors.INV_BAD_REQUEST.WithDetails("Unable to read the form data"))
			return
		}
		file := form.File["file"][0]

		imageId, inv_err := userCtrl.UploadUserImage(userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		imageName := "./../uploads/" + imageId.String() + ".jpeg"
		c.SaveUploadedFile(file, imageName)

		c.JSON(http.StatusOK, imageName)
	}
}

// @Summary Get ImagePath For User Profile Picture
// @Description Get ImagePath For User Profile Picture
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Success 200 {object} models.PicturePath
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /users-picture [get]
func GetImagePathForUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// single file
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		imageId, inv_err := userCtrl.GetImageIdFromUser(userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}
		imageName := "./../uploads/" + imageId.String() + ".jpeg"
		log.Print("Reading image: ", imageName)

		// Open the file
		fileData, err := os.Open(imageName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
			return
		}
		defer fileData.Close()

		// Read the first 512 bytes of the file to determine its content type
		fileHeader := make([]byte, 512)
		_, err = fileData.Read(fileHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
			return
		}
		fileContentType := http.DetectContentType(fileHeader)

		// Get the file info
		fileInfo, err := fileData.Stat()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get image info"})
			return
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", imageId.String()))
		c.Header("Content-Type", fileContentType)
		c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
		c.File(imageName)
	}
}

// @Summary Delete Img for User
// @Description Delete Picture from User and replace with ""
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Success 200
// @Failure 400 {object} models.INVErrorMessage
// @Failure 500 {object} models.INVErrorMessage
// @Router /users-picture [delete]
func RemoveImageForUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		// single file
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, inv_errors.INV_UNAUTHORIZED)
			return
		}

		inv_err := userCtrl.RemoveImageIdFromUser(userId)
		if inv_err != nil {
			utils.HandleErrorAndAbort(c, inv_err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
