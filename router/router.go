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
	r.SetTrustedProxies(nil)

	// Define routes for task CRUD operations
	taskRoutes := r.Group("/tasks") 
	{
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}

	// Return the configured router
	return r
}
