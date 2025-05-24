package router

import (
	"copilot/internal/interface/controller"

	"github.com/gin-gonic/gin"
)

// Router はルーティングのセットアップを処理する
type Router struct {
	engine        *gin.Engine
	todoController *controller.TodoController
}

// NewRouter は新しいRouterを作成する
func NewRouter(todoController *controller.TodoController) *Router {
	return &Router{
		engine:        gin.Default(),
		todoController: todoController,
	}
}

// SetupRoutes は全てのルートをセットアップする
func (r *Router) SetupRoutes() {
	// Todoルートを登録する
	r.todoController.RegisterRoutes(r.engine)
}

// Run はルーターを起動する
func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}