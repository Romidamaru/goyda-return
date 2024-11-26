package modules

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"simple-api/internal/modules/tasks"
)

type Router struct {
	tasks *tasks.TaskRouter
}

func NewRouter(r *gin.Engine, db *gorm.DB) *Router {
	return &Router{
		tasks: tasks.InitTaskRouter(r, db),
	}
}
