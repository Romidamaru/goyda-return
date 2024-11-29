package svc

import (
	"fmt"
	"simple-api/internal/modules/core/db"
	"simple-api/internal/modules/users/ent"
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
