package svc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"simple-api/internal/modules/tasks/ent"
)

type TasksService struct {
	DB *gorm.DB
}

func NewTasksSVC(db *gorm.DB) *TasksService {
	return &TasksService{DB: db}
}

func (s *TasksService) GetTasks() ([]ent.Task, error) {
	var tasks []ent.Task
	if err := s.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TasksService) GetTaskByID(id int) (ent.Task, error) {
	var task ent.Task
	if err := s.DB.First(&task, id).Error; err != nil {
		return ent.Task{}, err
	}

	return task, nil
}

func (s *TasksService) CreateTask(task ent.Task) (ent.Task, error) {
	// Use GORM to create the task in the database
	if err := s.DB.Create(&task).Error; err != nil {
		return ent.Task{}, err
	}
	// Return the created task
	return task, nil
}

func (s *TasksService) UpdateTask(task ent.Task) (ent.Task, error) {
	if err := s.DB.Save(&task).Error; err != nil {
		return ent.Task{}, err
	}
	return task, nil
}

func (s *TasksService) DeleteTask(id int) (gin.H, error) {
	var task ent.Task
	if err := s.DB.Delete(&task, id).Error; err != nil {
		return nil, err
	}
	return gin.H{"message": fmt.Sprintf("Task with ID %d has been deleted", id)}, nil
}
