package main

import (
	"example/tasksManager/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handlers.NewTaskHandler().Register(router)
	router.Run("localhost:8080")
}
