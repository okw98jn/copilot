package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUsers(c *gin.Context)
}

type userController struct{}

func NewUserController() UserController {
	return &userController{}
}

func (uc *userController) GetAllUsers(c *gin.Context) {
	users := []map[string]string{
		{"id": "1", "name": "John Doe", "email": "john@example.com"},
		{"id": "2", "name": "Jane Smith", "email": "jane@example.com"},
	}

	c.JSON(http.StatusOK, users)
}
