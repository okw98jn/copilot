package db

import (
	"context"
	"copilot/internal/domain/entity"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// NewDB creates and returns a new database connection
func NewDB() *bun.DB {
	// Get database connection parameters from environment variables
	dbName := getEnv("POSTGRES_DB", "postgres")
	dbUser := getEnv("POSTGRES_USER", "postgres")
	dbPassword := getEnv("POSTGRES_PASSWORD", "postgres")
	dbHost := getEnv("POSTGRES_HOST", "postgres")
	dbPort := getEnv("POSTGRES_PORT", "5432")

	// Create DSN string
	dsn := fmt.Sprintf("******%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open connection
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create bun.DB instance
	db := bun.NewDB(sqldb, pgdialect.New())

	// Register models
	db.RegisterModel((*entity.Todo)(nil))

	// Check if connection is alive
	if err := db.Ping(); err != nil {
		fmt.Printf("Failed to connect to database: %s\n", err.Error())
	}

	return db
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Close closes the database connection
func Close(db *bun.DB) error {
	return db.Close()
}

// InitSchema creates the necessary schema for the models
func InitSchema(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		(*entity.Todo)(nil),
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
