package router

import (
	"copilot/internal/interface/controller"

	"github.com/gin-gonic/gin"
)

// Router handles the routing setup
type Router struct {
	engine        *gin.Engine
	todoController *controller.TodoController
}

// NewRouter creates a new Router
func NewRouter(todoController *controller.TodoController) *Router {
	return &Router{
		engine:        gin.Default(),
		todoController: todoController,
	}
}

// SetupRoutes sets up all the routes
func (r *Router) SetupRoutes() {
	// Register todo routes
	r.todoController.RegisterRoutes(r.engine)
}

// Run starts the router
func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}