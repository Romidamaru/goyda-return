package users

import (
	"github.com/gin-gonic/gin"
	"simple-api/internal/core/db"
	"simple-api/internal/core/middleware"
	"simple-api/internal/pkg/auth/svc"
	"simple-api/internal/pkg/users/ctrl"
	userSvc "simple-api/internal/pkg/users/svc"
)

type UserRouter struct {
	userCtrl *ctrl.UsersController
	router   *gin.Engine
}

// InitUserRouter initializes the user routes with the provided Gin router and database
func InitUserRouter(router *gin.Engine, db db.Database) *UserRouter {
	// Initialize the service with the database connection
	userService := userSvc.NewUsersService(db)
	authService := svc.NewAuthService(db)
	// Initialize the controller with the service
	userController := ctrl.NewUsersController(userService, authService)

	// Define the routes for user operations
	//router.GET("/users", userController.GetUsers)         // Endpoint to fetch all users
	router.POST("/users/register", userController.Register) // Endpoint to create a new user
	router.POST("/users/login", userController.Login)       // Endpoint to create a new user
	//router.GET("/users/:id", userController.GetUserById) // Endpoint to fetch a user by ID
	router.Use(middleware.AuthMiddleware()).PUT("/users/:id", userController.UpdateUsername)
	router.Use(middleware.AuthMiddleware()).PUT("/users/recover/:id", userController.RecoverPassword) // Endpoint to update a user by ID
	//router.DELETE("/users/:id", userController.DeleteUser) // Endpoint to delete a user by ID

	return &UserRouter{
		userCtrl: userController,
	}
}
