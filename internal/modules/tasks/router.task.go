package tasks

import (
	"github.com/gin-gonic/gin"
	"simple-api/internal/modules/tasks/ctrl"
)

type TaskRouter struct {
	taskCtrl *ctrl.TasksController
	router   *gin.Engine
}

func InitTaskRouter(router *gin.Engine) *TaskRouter {
	controller := ctrl.NewTasksController()
	router.GET("/tasks", controller.GetTasks)
	//router.POST("/tasks", createTask)
	//router.GET("/tasks/:id", getTask)
	//router.PUT("/tasks/:id", updateTask)
	//router.DELETE("/tasks/:id", deleteTask)
	return &TaskRouter{
		taskCtrl: controller,
	}
}
