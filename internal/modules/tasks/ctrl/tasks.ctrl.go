package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-api/internal/config"
	"simple-api/internal/modules/tasks/dto"
	"simple-api/internal/modules/tasks/ent"
	"simple-api/internal/modules/tasks/svc"
)

type TasksController struct {
	tSvc *svc.TasksService
}

// NewTasksController creates a new TasksController instance
func NewTasksController(tSvc *svc.TasksService) *TasksController {
	return &TasksController{tSvc: tSvc}
}

func (ctrl *TasksController) GetTasks(c *gin.Context) {
	tasks, err := ctrl.tSvc.GetTasks() // Fetch tasks from the service (GORM DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks) // Send tasks as a JSON response
}

//func (ctrl *TasksController) GetTaskById(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
//		return
//	}
//
//	task, err := ctrl.tSvc.GetTaskByID(id)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, task)
//}

func (ctrl *TasksController) CreateTask(c *gin.Context) {
	var taskDTO dto.CreateTask

	// Bind JSON to DTO
	if err := c.ShouldBindJSON(&taskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate DTO using the generic validator
	if err := config.AppValidator.ValidateStruct(&taskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map DTO to Entity
	task := ent.Task{
		Name: taskDTO.Name,
		Done: taskDTO.Done,
	}

	// Call the service to create the task
	createdTask, err := ctrl.tSvc.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

//
//func updateTask(c *gin.Context) {
//	mu.Lock()
//	defer mu.Unlock()
//
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
//		return
//	}
//
//	for i, task := range tasks {
//		if task.ID == id {
//			if err := c.ShouldBindJSON(&tasks[i]); err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
//				return
//			}
//			tasks[i].ID = id // Ensure ID is preserved
//			c.JSON(http.StatusOK, tasks[i])
//			return
//		}
//	}
//
//	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
//}
//
//func deleteTask(c *gin.Context) {
//	mu.Lock()
//	defer mu.Unlock()
//
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
//		return
//	}
//
//	for i, task := range tasks {
//		if task.ID == id {
//			tasks = append(tasks[:i], tasks[i+1:]...)
//			c.JSON(http.StatusNoContent, nil)
//			return
//		}
//	}
//
//	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
//}
