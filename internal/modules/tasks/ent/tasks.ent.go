package ent

import (
	"gorm.io/gorm"
)

// TaskType represents the allowed values for the Type field
type TaskType int

const (
	TaskTypePersonal TaskType = 0
	TaskTypeWork     TaskType = 1
	TaskTypeOther    TaskType = 2
)

func (t TaskType) String() string {
	switch t {
	case TaskTypePersonal:
		return "personal"
	case TaskTypeWork:
		return "work"
	case TaskTypeOther:
		return "other"
	default:
		return "unknown"
	}
}

// Task represents a task model for GORM
type Task struct {
	gorm.Model
	Name        string   `gorm:"not null" json:"name"`           // Name of the task
	Description *string  `gorm:"type:text" json:"description"`   // Optional task description
	Type        TaskType `gorm:"type:int;default:0" json:"type"` // Task type (Personal, Work, Other)
	Done        bool     `json:"done"`                           // Whether the task is completed
	UserID      uint     `gorm:"not null" json:"user_id"`        // Foreign key to User
}
