// +build wireinject

package di

import (
	"copilot/internal/infrastructure/db"
	"copilot/internal/infrastructure/router"
	"copilot/internal/interface/controller"
	"copilot/internal/interface/repository"
	"copilot/internal/usecase"
	
	"github.com/google/wire"
)

// InitializeAPI sets up the API dependencies
func InitializeAPI() (*router.Router, error) {
	wire.Build(
		// Infrastructure
		db.NewDB,
		
		// Repository
		repository.NewTodoRepository,
		
		// UseCase
		usecase.NewTodoUseCase,
		
		// Controller
		controller.NewTodoController,
		
		// Router
		router.NewRouter,
	)
	
	return &router.Router{}, nil
}