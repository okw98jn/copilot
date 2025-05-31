//go:build wireinject
// +build wireinject

package di

import (
	"copilot/internal/adapter/controller"
	"copilot/internal/application/usecase"
	"copilot/internal/infrastructure/config"
	"copilot/internal/infrastructure/db"
	"copilot/internal/infrastructure/repository"
	"copilot/internal/infrastructure/router"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Init は依存性注入を使用してgin.Engineを初期化します
func Init() (*gin.Engine, error) {
	wire.Build(
		config.NewConfig,
		db.NewDB,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		controller.NewUserController,
		router.NewRouter,
	)
	return &gin.Engine{}, nil
}
