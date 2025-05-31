package controller

import (
	"net/http"

	"copilot/internal/application/usecase"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUsers(c *gin.Context)
}

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) UserController {
	return &userController{
		userUseCase: userUseCase,
	}
}

func (uc *userController) GetAllUsers(c *gin.Context) {
	users, err := uc.userUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
