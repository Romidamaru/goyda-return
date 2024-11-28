package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-api/internal/config"
	"simple-api/internal/modules/tasks/dto"
	"simple-api/internal/modules/tasks/ent"
	"simple-api/internal/modules/tasks/svc"
	"strconv"
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

	// Create a slice to hold the task response
	var taskResponses []map[string]any

	// Iterate over tasks and create response objects
	for _, task := range tasks {
		response := map[string]any{
			"id":          task.ID,
			"description": task.Description,
			"name":        task.Name,
			"type":        dto.FromInt(int(task.Type)),
			"done":        task.Done,
		}
		taskResponses = append(taskResponses, response)
	}

	c.JSON(http.StatusOK, taskResponses) // Return the array of task responses
}

func (ctrl *TasksController) GetTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := ctrl.tSvc.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	response := map[string]any{
		"id":          task.ID,
		"description": task.Description,
		"name":        task.Name,
		"type":        dto.FromInt(int(task.Type)),
		"done":        task.Done,
	}

	c.JSON(http.StatusOK, response)
}

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
		Name:        taskDTO.Name,
		Description: taskDTO.Description,
		Done:        taskDTO.Done,
		Type:        ent.TaskType(taskDTO.Type.ToInt()),
		UserID:      taskDTO.UserID, // Ensure UserID is correctly assigned
	}

	// Call the service to create the task
	createdTask, err := ctrl.tSvc.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// Return the created task as a response
	c.JSON(http.StatusCreated, createdTask)
}

func (ctrl *TasksController) UpdateTask(c *gin.Context) {
	// Extract task ID from the URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updateDTO dto.UpdateTask

	// Bind JSON to UpdateTask DTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Fetch the existing task from the database
	task, err := ctrl.tSvc.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Update fields conditionally based on provided values
	if updateDTO.Name != nil {
		task.Name = *updateDTO.Name
	}
	// In your UpdateTask controller:
	if updateDTO.Type != nil {
		task.Type = updateDTO.Type.ConvertToEnt()
	}
	if updateDTO.Done != nil {
		task.Done = *updateDTO.Done
	}

	if updateDTO.Description != nil {
		task.Description = updateDTO.Description
	}

	// Save the updated task
	updatedTask, err := ctrl.tSvc.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	response := map[string]any{
		"id":          updatedTask.ID,
		"name":        updatedTask.Name,
		"description": updateDTO.Description,
		"type":        dto.FromInt(int(updatedTask.Type)),
		"done":        updatedTask.Done,
		"user_id":     task.UserID,
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *TasksController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := ctrl.tSvc.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}
