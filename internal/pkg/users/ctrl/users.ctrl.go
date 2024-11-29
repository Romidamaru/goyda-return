package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-api/internal/pkg/auth/svc"
	"simple-api/internal/pkg/users/dto"
	userService "simple-api/internal/pkg/users/svc"
)

// UsersController defines the methods for handling user-related HTTP requests
type UsersController struct {
	uSvc *userService.UsersService
	aSvc *svc.AuthService
}

func NewUsersController(uSvc *userService.UsersService, aSvc *svc.AuthService) *UsersController {
	return &UsersController{
		uSvc: uSvc,
		aSvc: aSvc,
	}
}

// Register handles user registration and returns a JWT
func (ctrl *UsersController) Register(c *gin.Context) {
	var userDTO dto.CreateUser
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if the username is already taken
	if ctrl.uSvc.IsUsernameTaken(userDTO.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "Username is already taken"})
		return
	}

	if ctrl.uSvc.IsEmailTaken(userDTO.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email is already taken"})
		return
	}

	// Create the user
	user, err := ctrl.uSvc.CreateUser(userDTO.ToEntity())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT
	token, err := ctrl.aSvc.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   token,
	})
}

func (ctrl *UsersController) Login(c *gin.Context) {
	var loginDTO dto.LoginUser
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Authenticate the user by username
	user, err := ctrl.uSvc.AuthenticateUserByUsername(loginDTO.Username, loginDTO.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	token, err := ctrl.aSvc.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func (ctrl *UsersController) UpdateUsername(c *gin.Context) {
	var updateDTO dto.UpdateUser

	// Bind the request JSON body
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Retrieve userID from context (set by AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Check if the username already exists
	if ctrl.uSvc.IsUsernameTaken(updateDTO.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "Username is already taken"})
		return
	}

	// Update the user's username
	err := ctrl.uSvc.UpdateUsername(userID.(uint), updateDTO.Username) // Define the err variable
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update username"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Username updated successfully"})
}
