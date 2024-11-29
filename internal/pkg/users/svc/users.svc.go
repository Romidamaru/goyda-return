package svc

import (
	"fmt"
	"simple-api/internal/core/db"
	"simple-api/internal/pkg/users/ent"
	"simple-api/internal/utils"
)

// UsersService defines the methods for user-related operations
type UsersService struct {
	db db.Database
}

// NewUsersService creates a new UserService instance
func NewUsersService(db db.Database) *UsersService {
	return &UsersService{
		db: db,
	}
}

func (s *UsersService) IsUsernameTaken(username string) bool {
	var user ent.User
	err := s.db.GetDB().Where("username = ?", username).First(&user).Error
	return err == nil
}

func (s *UsersService) IsEmailTaken(email string) bool {
	var user ent.User
	err := s.db.GetDB().Where("email =?", email).First(&user).Error
	return err == nil
}

func (s *UsersService) AuthenticateUserByUsername(username, password string) (ent.User, error) {
	var user ent.User
	err := s.db.GetDB().Where("username = ?", username).First(&user).Error
	if err != nil {
		return ent.User{}, err
	}

	// Verify the password
	if !utils.VerifyPassword(password, user.Password) {
		return ent.User{}, fmt.Errorf("invalid password")
	}

	return user, nil
}

// CreateUser creates a new user in the database
func (s *UsersService) CreateUser(user ent.User) (ent.User, error) {
	if err := s.db.GetDB().Create(&user).Error; err != nil {
		return ent.User{}, nil
	}

	return user, nil
}

func (s *UsersService) UpdateUsername(userID uint, newUsername string) error {
	var user ent.User
	err := s.db.GetDB().Where("id = ?", userID).First(&user).Error
	if err != nil {
		return err // User not found
	}

	// Update the username
	user.Username = newUsername
	return s.db.GetDB().Save(&user).Error
}

func (s *UsersService) RecoverPassword(userID uint, currentPassword, newPassword string) error {
	var user ent.User

	// Fetch the user from the database by ID
	err := s.db.GetDB().Where("id = ?", userID).First(&user).Error
	if err != nil {
		return fmt.Errorf("user not found: %w", err) // Log the error more specifically
	}

	// Verify if the current password matches the stored hashed password
	if !utils.VerifyPassword(currentPassword, user.Password) {
		return fmt.Errorf("current password is incorrect") // Log password mismatch
	}

	// Hash the new password before saving it
	hashedPassword := utils.HashPassword(newPassword)

	// Update the user's password with the new hashed password
	user.Password = hashedPassword

	// Save the updated user record in the database
	if err := s.db.GetDB().Save(&user).Error; err != nil {
		return fmt.Errorf("failed to update password: %w", err) // Log any database save issues
	}

	return nil
}
