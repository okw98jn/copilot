package controller

import (
	"copilot/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TodoController handles HTTP requests for todos
type TodoController struct {
	todoUseCase usecase.TodoUseCase
}

// NewTodoController creates a new todo controller
func NewTodoController(todoUseCase usecase.TodoUseCase) *TodoController {
	return &TodoController{
		todoUseCase: todoUseCase,
	}
}

// GetAll handles GET requests to retrieve all todos
func (c *TodoController) GetAll(ctx *gin.Context) {
	todos, err := c.todoUseCase.GetAllTodos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": todos})
}

// RegisterRoutes registers the routes for the TodoController
func (c *TodoController) RegisterRoutes(router *gin.Engine) {
	todoGroup := router.Group("/api/todos")
	{
		todoGroup.GET("", c.GetAll)
	}
}