package controller

import (
	"copilot/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TodoController はTodoのHTTPリクエストを処理する
type TodoController struct {
	todoUseCase usecase.TodoUseCase
}

// NewTodoController は新しいTodoコントローラを作成する
func NewTodoController(todoUseCase usecase.TodoUseCase) *TodoController {
	return &TodoController{
		todoUseCase: todoUseCase,
	}
}

// GetAll は全てのTodoを取得するGETリクエストを処理する
func (c *TodoController) GetAll(ctx *gin.Context) {
	todos, err := c.todoUseCase.GetAllTodos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": todos})
}

// RegisterRoutes はTodoControllerのルートを登録する
func (c *TodoController) RegisterRoutes(router *gin.Engine) {
	todoGroup := router.Group("/api/todos")
	{
		todoGroup.GET("", c.GetAll)
	}
}