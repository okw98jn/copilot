package main

import (
	"context"
	"copilot/internal/infrastructure/db"
	"copilot/internal/infrastructure/di"
	"fmt"
	"log"
)

func main() {
	// Initialize DB connection
	database := db.NewDB()
	defer db.Close(database)

	// Initialize schema
	ctx := context.Background()
	if err := db.InitSchema(ctx, database); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}

	// Initialize API using Wire DI
	router, err := di.InitializeAPI()
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	// Setup routes
	router.SetupRoutes()

	// Start server
	fmt.Println("Starting server on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
