package handlers

import (
	"example/tasksManager/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{}
var invalidIdMessage string = "This ID doesn't exists."

type TaskHandler struct {
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) getTasks(c *gin.Context) {
	tasksFiltered := getActiveTasks()
	c.IndentedJSON(http.StatusOK, tasksFiltered)
}

func (h *TaskHandler) getTask(c *gin.Context) {
	var idParam = c.Param("id")
	u64, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return
	}

	id := uint(u64)
	var currentTaskIndex = getIndexId(id, getActiveTasks())

	if currentTaskIndex == -1 {
		c.IndentedJSON(http.StatusNotFound, invalidIdMessage)
		return
	}

	var currentTask = tasks[currentTaskIndex]

	c.IndentedJSON(http.StatusOK, currentTask)
}

func (h *TaskHandler) addTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newTask.ID = getTaskById() + 1
	newTask.CreatedAt = time.Now()

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusOK, newTask)
}

func (h *TaskHandler) deleteTask(c *gin.Context) {
	var idParam = c.Param("id")
	u64, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return
	}

	id := uint(u64)
	var currentTaskIndex = getIndexId(id, getActiveTasks())

	if currentTaskIndex == -1 {
		c.IndentedJSON(http.StatusNotFound, invalidIdMessage)
		return
	}

	deletedAt := time.Now()
	tasks[currentTaskIndex].DeletedAt = &deletedAt

	c.IndentedJSON(http.StatusOK, "Deletado com sucesso")
}

func (h *TaskHandler) updateTask(c *gin.Context) {
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}

	var idUpdatedTask = getIndexId(updatedTask.ID, tasks)

	if idUpdatedTask == -1 || tasks[idUpdatedTask].DeletedAt != nil {
		c.IndentedJSON(http.StatusNotFound, invalidIdMessage)
		return
	}

	tasks[idUpdatedTask].Title = updatedTask.Title
	tasks[idUpdatedTask].Description = updatedTask.Description
	tasks[idUpdatedTask].Status = updatedTask.Status
	tasks[idUpdatedTask].UpdatedAt = time.Now()

	if !updatedTask.DueDate.IsZero() {
		tasks[idUpdatedTask].DueDate = updatedTask.DueDate
	}

	c.IndentedJSON(http.StatusOK, tasks[idUpdatedTask])
}

func getTaskById() uint {
	if len(tasks) == 0 {
		return 0
	}

	return tasks[len(tasks)-1].ID
}

func getIndexId(id uint, tasks []models.Task) int {

	if len(tasks) == 0 {
		return -1
	}

	var min = 0
	var max = len(tasks) - 1

	for min <= max {
		var average = (min + max) / 2

		if tasks[average].ID == id {
			return average
		}

		if id > tasks[average].ID {
			min = average + 1
		} else {
			max = average - 1
		}
	}

	return -1
}

func getActiveTasks() []models.Task {
	var activeTasks = []models.Task{}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].DeletedAt == nil {
			activeTasks = append(activeTasks, tasks[i])
		}
	}
	return activeTasks
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
