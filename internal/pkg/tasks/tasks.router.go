package tasks

import (
	"github.com/gin-gonic/gin"
	"simple-api/internal/core/db"
	"simple-api/internal/pkg/tasks/ctrl"
	"simple-api/internal/pkg/tasks/svc"
)

type TaskRouter struct {
	taskCtrl *ctrl.TasksController
	router   *gin.Engine
}

func InitTaskRouter(router *gin.Engine, db db.Database) *TaskRouter {
	// Initialize the service with the database connection
	taskService := svc.NewTasksSVC(db)
	// Initialize the controller with the service
	taskController := ctrl.NewTasksController(taskService)

	// Define the routes
	router.GET("/tasks", taskController.GetTasks)
	// You can add more routes here, such as POST, PUT, DELETE
	router.POST("/tasks", taskController.CreateTask)
	router.GET("/tasks/:id", taskController.GetTaskById)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	return &TaskRouter{
		taskCtrl: taskController,
	}
}
