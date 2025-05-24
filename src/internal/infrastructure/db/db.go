package db

import (
	"context"
	"copilot/internal/interface/repository"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// NewDB は新しいデータベース接続を作成して返す
func NewDB() *bun.DB {
	// 環境変数からデータベース接続パラメータを取得する
	dbName := getEnv("POSTGRES_DB", "postgres")
	dbUser := getEnv("POSTGRES_USER", "postgres")
	dbPassword := getEnv("POSTGRES_PASSWORD", "postgres")
	dbHost := getEnv("POSTGRES_HOST", "postgres")
	dbPort := getEnv("POSTGRES_PORT", "5432")

	// 接続文字列を作成する
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", 
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// 接続を開く
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))

	// bun.DBインスタンスを作成する
	db := bun.NewDB(sqldb, pgdialect.New())

	// モデルを登録する
	db.RegisterModel((*repository.BunTodo)(nil))

	// 接続が有効かチェックする
	if err := db.Ping(); err != nil {
		fmt.Printf("Failed to connect to database: %s\n", err.Error())
	}

	return db
}

// getEnv は環境変数を取得するか、デフォルト値を返す
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Close はデータベース接続を閉じる
func Close(db *bun.DB) error {
	return db.Close()
}

// InitSchema はモデルに必要なスキーマを作成する
func InitSchema(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		(*repository.BunTodo)(nil),
	}

	for _, model := range models {
		_, err := db.NewCreateTable().
			Model(model).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
