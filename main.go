package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"simple-api/internal/modules"
	"simple-api/internal/modules/core/db"
)

// TODO:
// task
// GET - get tasks list (list of MY tasks by MY userID (from token) related to tasks) (filter showDeleted)
// POST - create task (create new task for current user by userID(from token))
// PUT - update task /:taskID (update task's name and description and type ONLY. validate by userId from current token)
// PATCH - done task /:taskID (mark specific task as DONE by taskID + userID(from token))
// DELETE - delete task /:taskID (deletes task by id) (paranoid)
//
// user
// POST - auth (uses JWT to attach token for users ) token struct = {userId} (access/refresh tokens)
// POST - registration (creates new user and returns user jwt token reusing simple method for token from auth) - email should be unique, - username - also should be unique
// PUT - update user's username - also should be unique
// PUT - recover - recovers users password with verification old password
// DELETE - delete user (cascade) - deletes user and its tasks (user id from token) (NOT paranoid)

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
