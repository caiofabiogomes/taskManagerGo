package handlers

import (
	"example/tasksManager/internal/dto"
	"example/tasksManager/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var invalidIdMessage string = "This ID doesn't exists."

type TaskHandler struct {
	service services.ITaskService
}

func NewTaskHandler(service services.ITaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) getTasks(c *gin.Context) {
	ctx := c.Request.Context()

	tasks, error := h.service.GetAll(ctx)

	if error != nil {
		c.IndentedJSON(http.StatusInternalServerError, error)
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func (h *TaskHandler) getTask(c *gin.Context) {
	var idParam = c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		// UUID inválido
		c.JSON(400, gin.H{"error": "invalid UUID"})
		return
	}
	ctx := c.Request.Context()
	currentTask, err := h.service.FindById(ctx, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if currentTask == nil {
		c.IndentedJSON(http.StatusNotFound, "Task wasn't found.")
		return
	}

	c.IndentedJSON(http.StatusOK, currentTask)
}

func (h *TaskHandler) addTask(c *gin.Context) {
	var newTask dto.AddTaskRequest
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	err := h.service.Add(newTask)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, newTask)
}

func (h *TaskHandler) deleteTask(c *gin.Context) {
	var id = c.Param("id")
	err := h.service.Delete(dto.DeleteTaskRequest{ID: id})

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "Deletado com sucesso")
}

func (h *TaskHandler) updateTask(c *gin.Context) {
	var updatedTask dto.UpdateTaskRequest
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}

	affectedRows, err := h.service.Update(updatedTask)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if affectedRows == 0 {
		c.IndentedJSON(http.StatusNotFound, "Task wasn't found")
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) Register(router *gin.Engine) {
	routes := router.Group("/tasks")
	{
		routes.GET("", h.getTasks)
		routes.GET("/:id", h.getTask)
		routes.POST("", h.addTask)
		routes.DELETE("/:id", h.deleteTask)
		routes.PUT("", h.updateTask)
	}
}
