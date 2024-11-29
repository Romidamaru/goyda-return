package svc

import (
	"simple-api/internal/core/config"
	"simple-api/internal/core/db"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	db     db.Database
	secret string
}

// NewAuthService creates a new AuthService instance
func NewAuthService(db db.Database) *AuthService {
	return &AuthService{
		db:     db,
		secret: config.Inst().SecretKey,
	}
}

// GenerateToken generates a JWT for a user
func (s *AuthService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}
