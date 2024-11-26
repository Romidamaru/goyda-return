package dto

type TaskType string

const (
	TaskTypePersonal TaskType = "personal"
	TaskTypeWork     TaskType = "work"
	TaskTypeOther    TaskType = "other"
)

// ToInt converts TaskType string to its corresponding integer value
func (t TaskType) ToInt() int {
	switch t {
	case TaskTypePersonal:
		return 0
	case TaskTypeWork:
		return 1
	case TaskTypeOther:
		return 2
	default:
		return -1 // Invalid case
	}
}

// FromInt converts an integer value to its corresponding TaskType string
func FromInt(value int) TaskType {
	switch value {
	case 0:
		return TaskTypePersonal
	case 1:
		return TaskTypeWork
	case 2:
		return TaskTypeOther
	default:
		return "unknown" // Invalid case
	}
}

type CreateTask struct {
	Name string   `json:"name" validate:"required,min=3"`
	Type TaskType `json:"type" validate:"required"`
	Done bool     `json:"done"`
}
