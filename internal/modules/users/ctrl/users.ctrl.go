package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-api/internal/modules/auth/svc"
	"simple-api/internal/modules/users/dto"
	userService "simple-api/internal/modules/users/svc"
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
