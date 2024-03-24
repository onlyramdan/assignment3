package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tugas/auth"
	"tugas/helper"
	"tugas/user"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		msgError := gin.H{"errors": errors}
		response := helper.APIresponse("Register Account Failed", http.StatusUnprocessableEntity, "Failed", msgError)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIresponse("Register Account Failed", http.StatusBadRequest, "Failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helper.APIresponse("Register Account Failed", http.StatusBadRequest, "Failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.APIresponse("Account Has been Created", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		msgError := gin.H{"errors": errors}
		response := helper.APIresponse("Login Failed", http.StatusUnprocessableEntity, "Failed", msgError)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userService.LoginUser(input)

	if err != nil {
		response := helper.APIresponse("Login Failed", http.StatusBadRequest, "Failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID)

	if err != nil {
		response := helper.APIresponse("Register Account Failed", http.StatusBadRequest, "Failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loginUser, token)
	response := helper.APIresponse("Successfuly Loggin", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}
func (h *userHandler) CekEmail(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		msgError := gin.H{"errors": errors}
		response := helper.APIresponse("Email Checking Failed", http.StatusUnprocessableEntity, "Failed", msgError)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		msgError := gin.H{"error": "Server Error"}
		response := helper.APIresponse("Email Checking Failed", http.StatusUnprocessableEntity, "Failed", msgError)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}
	msg := "Email has been Registered"

	if isEmailAvailable {
		msg = "Email Available"
	}
	response := helper.APIresponse(msg, http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
}
func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Upload Failed", http.StatusBadRequest, "Failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	ID_User := currentUser.ID
	path := fmt.Sprintf("images/profile/%d-%s", ID_User, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Uploaded Failed", http.StatusBadRequest, "Failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(ID_User, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Uploaded Failed", http.StatusBadRequest, "Failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse("Success Uploaded Avatar", http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
	return
}
