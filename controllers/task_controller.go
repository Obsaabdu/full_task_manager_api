// task_controller.go
// Handles HTTP requests for task CRUD operations: get all, get one, create, update, and delete tasks.
package controllers

import (
	"net/http"

	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// CreateTask handles POST /tasks. Binds JSON to new Task, loads tasks, adds new task, and saves. Returns error if any step fails.
func CreateTask(c *gin.Context) {
	var task models.Task
	
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	id, err := data.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task: "+ err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id":id})
}

// GetTasks handles GET /tasks. Loads all tasks and returns them as JSON. If loading fails, returns error.
func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks: "+ err.Error()})
	}
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load tasks."})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask handles GET /tasks/:id. Loads all tasks, finds the one with matching ID, and returns it. If not found, returns error.
func GetTask(c *gin.Context) {
	id:= c.Param("id")

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found."})
		return
	}
	c.JSON(http.StatusOK, task)
}


// UpdateTask handles PUT /tasks/:id. Loads all tasks, finds the one with matching ID, updates its fields, and saves. Returns error if not found or loading fails.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to update task."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// DeleteTask handles DELETE /tasks/:id. Loads all tasks, finds the one with matching ID, removes it, and saves. Returns error if not found or loading fails.
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to delete task."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}
