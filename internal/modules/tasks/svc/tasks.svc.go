package svc

import (
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

//
//func (s *TasksService) GetTaskByID(id int) (*ent.Task, error) {
//	ent.Mu.Lock()
//	defer ent.Mu.Unlock()
//
//	for _, task := range ent.Tasks {
//		if task.ID == id {
//			return &task, nil
//		}
//	}
//	return nil, fmt.Errorf("task with ID %d not found", id)
//}

func (s *TasksService) CreateTask(task ent.Task) (ent.Task, error) {
	// Use GORM to create the task in the database
	if err := s.DB.Create(&task).Error; err != nil {
		return ent.Task{}, err
	}

	// Return the created task
	return task, nil
}
