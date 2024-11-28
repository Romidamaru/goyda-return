package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-api/internal/config"
	"simple-api/internal/modules/users/dto"
	"simple-api/internal/modules/users/ent"
	"simple-api/internal/modules/users/svc"
)

// UsersController defines the methods for handling user-related HTTP requests
type UsersController struct {
	uSvc *svc.UsersService
}

// NewUsersController creates a new UsersController instance
func NewUsersController(uSvc *svc.UsersService) *UsersController {
	return &UsersController{
		uSvc: uSvc,
	}
}

// CreateUser handles the HTTP request to create a new user
func (ctrl *UsersController) CreateUser(c *gin.Context) {
	var userDTO dto.CreateUserDTO

	// Bind JSON to DTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate DTO using the generic validator
	if err := config.AppValidator.ValidateStruct(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map DTO to Entity
	user := ent.User{
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Password: userDTO.Password, // Password should be hashed before saving
	}

	// Call the service to create the user
	createdUser, err := ctrl.uSvc.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the created user as the response
	c.JSON(http.StatusCreated, createdUser)
}
