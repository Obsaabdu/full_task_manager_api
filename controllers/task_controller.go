// task_controller.go
// Handles HTTP requests for task CRUD operations: get all, get one, create, update, and delete tasks.
package controllers

import (
	"net/http"
	"strconv"

	"task_manager/data"
	"task_manager/dto"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks, err := data.ReadTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load tasks."})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTasks handles GET /tasks. Loads all tasks and returns them as JSON. If loading fails, returns error.
func GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found."})
		return
	}
	c.JSON(http.StatusOK, task)
}

// GetTask handles GET /tasks/:id. Loads all tasks, finds the one with matching ID, and returns it. If not found, returns error.
// CreateTask handles POST /tasks. Binds JSON to new Task, loads tasks, adds new task, and saves. Returns error if any step fails.
func CreateTask(c *gin.Context) {
	var input dto.CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	newTask := models.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Status:      models.StatusPending,
	}

	if err := data.AddTask(newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save task."})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// UpdateTask handles PUT /tasks/:id. Loads all tasks, finds the one with matching ID, updates its fields, and saves. Returns error if not found or loading fails.
func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input dto.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update format"})
		return
	}

	existing, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found."})
		return
	}
	
	updated := *existing

	if input.Title != "" {
		updated.Title = input.Title
	}

	if input.Description != "" {
		updated.Description = input.Description
	}

	if !input.DueDate.IsZero() {
		updated.DueDate = input.DueDate
	}

	if input.Status != "" {
		updated.Status = input.Status
	}


	if err := data.UpdateTask(id, updated); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// DeleteTask handles DELETE /tasks/:id. Loads all tasks, finds the one with matching ID, removes it, and saves. Returns error if not found or loading fails.
func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}
