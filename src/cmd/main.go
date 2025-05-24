package main

import (
	"context"
	"copilot/internal/infrastructure/db"
	"copilot/internal/infrastructure/di"
	"fmt"
	"log"
)

func main() {
	// DB接続を初期化する
	database := db.NewDB()
	defer db.Close(database)

	// スキーマを初期化する
	ctx := context.Background()
	if err := db.InitSchema(ctx, database); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}

	// Wire DIを使用してAPIを初期化する
	router, err := di.InitializeAPI()
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	// ルートをセットアップする
	router.SetupRoutes()

	// サーバーを起動する
	fmt.Println("Starting server on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
