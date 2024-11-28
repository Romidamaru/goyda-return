// users.router.go
package users

import (
	"github.com/gin-gonic/gin"
	"simple-api/internal/modules/core/db"
	"simple-api/internal/modules/users/ctrl"
	"simple-api/internal/modules/users/svc"
)

type UserRouter struct {
	userCtrl *ctrl.UsersController
	router   *gin.Engine
}

// InitUserRouter initializes the user routes with the provided Gin router and database
func InitUserRouter(router *gin.Engine, db db.Database) *UserRouter {
	// Initialize the service with the database connection
	userService := svc.NewUsersService(db)
	// Initialize the controller with the service
	userController := ctrl.NewUsersController(userService)

	// Define the routes for user operations
	//router.GET("/users", userController.GetUsers)         // Endpoint to fetch all users
	router.POST("/users", userController.CreateUser) // Endpoint to create a new user
	//router.GET("/users/:id", userController.GetUserById) // Endpoint to fetch a user by ID
	//router.PUT("/users/:id", userController.UpdateUser)  // Endpoint to update a user by ID
	//router.DELETE("/users/:id", userController.DeleteUser) // Endpoint to delete a user by ID

	return &UserRouter{
		userCtrl: userController,
	}
}
