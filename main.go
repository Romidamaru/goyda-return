package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"simple-api/internal/modules"
	"simple-api/internal/modules/core/db"
)

func main() {
	// Database connection details
	dsn := "postgres://postgres:yura2rubles@localhost:5432/goyda-return?sslmode=disable"

	// Initialize database
	database, err := db.New(dsn, true).Create(db.PostgresDB)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	// Initialize Gin router and modules
	router := gin.Default()
	_ = modules.NewRouter(router, database)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
