package dto

type CreateTask struct {
	Name string `json:"name" validate:"required,min=3"` // At least 3 characters
	Done bool   `json:"done"`                           // Optional, default false
}
