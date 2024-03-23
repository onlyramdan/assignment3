package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tugas/helper"
	"tugas/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
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
	formatter := user.FormatUser(newUser, "toker")
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
	formatter := user.FormatUser(loginUser, "toker")
	response := helper.APIresponse("Successfuly Loggin", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}
