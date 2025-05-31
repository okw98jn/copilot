//go:build wireinject
// +build wireinject

package di

import (
	"copilot/internal/adapter/controller"
	"copilot/internal/infrastructure/router"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Init は依存性注入を使用してgin.Engineを初期化します
func Init() (*gin.Engine, error) {
	wire.Build(
		// config.New,
		// db.New,
		controller.NewUserController,
		router.NewRouter,
	)
	return &gin.Engine{}, nil
}
