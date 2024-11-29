package pkg

import (
	"github.com/gin-gonic/gin"
	"simple-api/internal/core/db"
	"simple-api/internal/pkg/tasks"
	"simple-api/internal/pkg/users"
)

type Router struct {
	tasks *tasks.TaskRouter
	users *users.UserRouter
}

func NewRouter(r *gin.Engine, db db.Database) *Router {
	return &Router{
		tasks: tasks.InitTaskRouter(r, db),
		users: users.InitUserRouter(r, db),
	}
}
