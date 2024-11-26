package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-api/internal/modules/tasks/svc"
)

type TasksController struct {
	tSvc *svc.TasksService
}

func NewTasksController() *TasksController {
	return &TasksController{
		tSvc: svc.NewTasksSVC(),
	}
}

func (ctrl *TasksController) GetTasks(c *gin.Context) {
	tasks := ctrl.tSvc.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func (ctrl *TasksController) CreateTask(c *gin.Context) {
	tasks := ctrl.tSvc.CreateTask()
	c.JSON(http.StatusOK, tasks)
}

//}
//
//func getTask(c *gin.Context) {
//	mu.Lock()
//	defer mu.Unlock()
//
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
//		return
//	}
//
//	for _, task := range tasks {
//		if task.ID == id {
//			c.JSON(http.StatusOK, task)
//			return
//		}
//	}
//
//	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
//}
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
