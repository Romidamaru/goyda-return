package main

import (
	"github.com/gin-gonic/gin"
	"simple-api/internal/modules"
)

func main() {
	router := gin.Default()

	_ = modules.NewRouter(router)

	router.Run(":8080")
}
