package svc

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func (s *TasksService) CreateTask() (ent.Task, error) {
	ent.Mu.Lock()
	defer ent.Mu.Unlock()

	var task ent.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	task.ID = ent.IdCount
	ent.IdCount++
	ent.Tasks = append(ent.Tasks, ent.Task{})

	c.JSON(http.StatusCreated, task)
}
