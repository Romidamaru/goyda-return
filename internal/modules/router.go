package modules

import (
	"github.com/gin-gonic/gin"
	"simple-api/internal/modules/core/db"
	"simple-api/internal/modules/tasks"
)

type Router struct {
	tasks *tasks.TaskRouter
}

func NewRouter(r *gin.Engine, db db.Database) *Router {
	return &Router{
		tasks: tasks.InitTaskRouter(r, db),
	}
}
