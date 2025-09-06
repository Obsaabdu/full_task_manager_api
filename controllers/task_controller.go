// task_controller.go
// Handles HTTP requests for task CRUD operations: get all, get one, create, update, and delete tasks.
package controllers

import (
	"net/http"
	"time"

	"task_manager/data"
	"task_manager/middleware"
	"task_manager/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserService *data.UserService
	TaskService *data.TaskService

}

// --- Auth ---
func (ctrl *Controller) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err:= ctrl.UserService.Register(user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"messaage": "User registered successfully"})
}

func (ctrl *Controller) Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := ctrl.UserService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour*24).Unix(),
	})

	tokenString, _ := token.SignedString(middleware.JWtSecret)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (ctrl *Controller) GetUsers(c *gin.Context) {
	users, err := ctrl.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users: "+ err.Error()})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load users."})
	}
	c.JSON(http.StatusOK, users)
}

// CreateTask handles POST /tasks. Binds JSON to new Task, loads tasks, adds new task, and saves. Returns error if any step fails.
func (ctrl *Controller) CreateTask(c *gin.Context) {
	var task models.Task
	
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := ctrl.TaskService.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task: "+ err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"Task added"})
}

// GetTasks handles GET /tasks. Loads all tasks and returns them as JSON. If loading fails, returns error.
func (ctrl *Controller) GetTasks(c *gin.Context) {
	tasks, err := ctrl.TaskService.GetTasks()
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
func (ctrl *Controller) GetTask(c *gin.Context) {
	id:= c.Param("id")

	task, err := ctrl.TaskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found."})
		return
	}
	c.JSON(http.StatusOK, task)
}


// UpdateTask handles PUT /tasks/:id. Loads all tasks, finds the one with matching ID, updates its fields, and saves. Returns error if not found or loading fails.
func (ctrl *Controller) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.TaskService.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to update task."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// DeleteTask handles DELETE /tasks/:id. Loads all tasks, finds the one with matching ID, removes it, and saves. Returns error if not found or loading fails.
func (ctrl *Controller) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.TaskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to delete task."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}
