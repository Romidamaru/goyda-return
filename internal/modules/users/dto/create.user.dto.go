package dto

// CreateUserDTO represents the data required to create a new user
type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"` // Simple validation: password must have at least 8 characters
}
