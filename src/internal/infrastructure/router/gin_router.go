package router

import (
	"copilot/internal/adapter/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController controller.UserController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		userGroup := api.Group("/users")
		{
			userGroup.GET("", userController.GetAllUsers)
		}
	}

	return router
}
