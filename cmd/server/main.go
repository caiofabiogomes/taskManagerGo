package main

import (
	"example/tasksManager/internal/handlers"
	"example/tasksManager/internal/repository"
	"example/tasksManager/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	taskContext := repository.NewTaskContext()
	repo := repository.NewTaskRepository(*taskContext)
	service := services.NewTaskService(repo)
	handlers.NewTaskHandler(service).Register(router)
	router.Run("localhost:8080")
}
