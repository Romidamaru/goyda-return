package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain text password and returns the hashed version.
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password: " + err.Error()) // Handle this gracefully in production
	}
	return string(hashedPassword)
}

// VerifyPassword compares a plain text password with a hashed password.
func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // If nil, the password matches
}
