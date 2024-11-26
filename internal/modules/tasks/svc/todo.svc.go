package svc

import (
	"fmt"
	"simple-api/internal/modules/tasks/ent"
)

type TasksService struct {
}

func NewTasksSVC() *TasksService {
	return &TasksService{}
}

func (s *TasksService) GetTasks() []ent.Task {
	ent.Mu.Lock()
	defer ent.Mu.Unlock()

	return ent.Tasks
}

func (s *TasksService) GetTaskByID(id int) (*ent.Task, error) {
	ent.Mu.Lock()
	defer ent.Mu.Unlock()

	for _, task := range ent.Tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (s *TasksService) CreateTask(task ent.Task) (ent.Task, error) {
	ent.Mu.Lock()
	defer ent.Mu.Unlock()

	// Assign a new ID to the task
	task.ID = ent.IdCount
	ent.IdCount++

	// Add the task to the list
	ent.Tasks = append(ent.Tasks, task)

	// Return the created task
	return task, nil
}
