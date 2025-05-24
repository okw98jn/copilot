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

// InitializeAPI はAPIの依存関係をセットアップする
func InitializeAPI() (*router.Router, error) {
	wire.Build(
		// インフラストラクチャ
		db.NewDB,
		
		// リポジトリ
		repository.NewTodoRepository,
		
		// ユースケース
		usecase.NewTodoUseCase,
		
		// コントローラ
		controller.NewTodoController,
		
		// ルーター
		router.NewRouter,
	)
	
	return &router.Router{}, nil
}