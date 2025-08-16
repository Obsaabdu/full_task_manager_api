// router.go
// Sets up the Gin router and defines API routes for the Task Manager.
package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter creates and configures a new Gin router with all task routes.
func SetupRouter() *gin.Engine {
	// Create a new Gin router
	r := gin.Default()

	// Define routes for task CRUD operations
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTask)
	r.POST("/tasks", controllers.CreateTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	// Return the configured router
	return r
}
