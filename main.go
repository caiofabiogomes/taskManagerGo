package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var tasks = []Task{}

var invalidIdMessage string = "This ID doesn't exists."

func getTasks(c *gin.Context) {
	tasksFiltered := getActiveTasks()
	c.IndentedJSON(http.StatusOK, tasksFiltered)
}

func getTask(c *gin.Context) {
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

func addTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newTask.ID = getTaskById() + 1
	newTask.CreatedAt = time.Now()

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusOK, newTask)
}

func deleteTask(c *gin.Context) {
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

func updateTask(c *gin.Context) {
	var updatedTask Task
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}

	var idUpdatedTask = getIndexId(updatedTask.ID, getActiveTasks())

	if idUpdatedTask == -1 {
		c.IndentedJSON(http.StatusNotFound, invalidIdMessage)
		return
	}

	tasks[idUpdatedTask].Title = updatedTask.Title
	tasks[idUpdatedTask].Description = updatedTask.Description
	tasks[idUpdatedTask].DueDate = updatedTask.DueDate
	tasks[idUpdatedTask].Status = updatedTask.Status
	tasks[idUpdatedTask].UpdatedAt = time.Now()

	c.IndentedJSON(http.StatusOK, tasks[idUpdatedTask])
}

func getTaskById() uint {
	if len(tasks) == 0 {
		return 0
	}

	return tasks[len(tasks)-1].ID
}

func getIndexId(id uint, tasks []Task) int {

	if len(tasks) == 0 {
		return -1
	}

	var min = 0
	var max = len(tasks) - 1

	for min <= max {
		var average = max / 2

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

func getActiveTasks() []Task {
	var activeTasks = []Task{}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].DeletedAt == nil {
			activeTasks = append(activeTasks, tasks[i])
		}
	}
	return activeTasks
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/task/:id", getTask)
	router.POST("/task", addTask)
	router.DELETE("/task/:id", deleteTask)
	router.PUT("/task", updateTask)
	router.Run("localhost:8080")
}
