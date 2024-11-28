package dto

import "simple-api/internal/modules/tasks/ent"

// ConvertToEnt converts a dto.TaskType (string) to ent.TaskType (int)
func (t *TaskType) ConvertToEnt() ent.TaskType {
	if t == nil {
		return ent.TaskTypePersonal // Default value, or handle as needed
	}

	switch *t {
	case TaskTypePersonal:
		return ent.TaskTypePersonal
	case TaskTypeWork:
		return ent.TaskTypeWork
	case TaskTypeOther:
		return ent.TaskTypeOther
	default:
		return ent.TaskTypePersonal // Default to personal or handle error case
	}
}

// UpdateTask - DTO for updating a task
type UpdateTask struct {
	Name        *string   `json:"name,omitempty"`        // Pointer to allow optional value
	Description *string   `json:"description,omitempty"` // Optional description field
	Type        *TaskType `json:"type,omitempty"`        // Pointer to allow optional value
	Done        *bool     `json:"done,omitempty"`        // Pointer to allow optional value
}
