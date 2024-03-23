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
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	response := helper.APIresponse("Account Has been Created", http.StatusOK, "Success", user)

	c.JSON(http.StatusOK, response)

}
