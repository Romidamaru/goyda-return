package svc

import (
	"simple-api/internal/modules/core/db"
	"simple-api/internal/modules/users/ent"
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

// CreateUser creates a new user in the database
func (s *UsersService) CreateUser(user ent.User) (ent.User, error) {
	if err := s.db.GetDB().Create(&user).Error; err != nil {
		return ent.User{}, nil
	}

	return user, nil
}
