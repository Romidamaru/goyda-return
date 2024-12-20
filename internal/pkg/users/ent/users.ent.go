package ent

import (
	"gorm.io/gorm"
	"simple-api/internal/pkg/tasks/ent"
)

// User represents a user model for GORM
type User struct {
	gorm.Model
	Username string     `gorm:"unique;not null" json:"username"` // Unique username
	Email    string     `gorm:"unique;not null" json:"email"`    // Unique email address
	Password string     `gorm:"not null" json:"password"`        // Encrypted password
	Tasks    []ent.Task `gorm:"foreignKey:UserID"`
}
