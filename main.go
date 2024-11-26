package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"simple-api/internal/modules"
	"simple-api/internal/modules/tasks/ent"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "postgres://postgres:yura2rubles@localhost:5432/goyda-return?sslmode=disable"

	// Connect to the database using GORM
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema (this will create the tables in the database)
	err = db.AutoMigrate(&ent.Task{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}
}

func main() {
	router := gin.Default()

	_ = modules.NewRouter(router, db)

	_ = router.Run(":8080")
}
