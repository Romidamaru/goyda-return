package ent

import (
	"gorm.io/gorm"
)

// Task represents a task model for GORM
type Task struct {
	gorm.Model
	//ID   int    `gorm:"primaryKey" json:"id"` // Primary key
	Name string `gorm:"not null" json:"name"` // Not null constraint
	Done bool   `json:"done"`                 // Boolean field
}
