// router.go
// Sets up the Gin router and defines API routes for the Task Manager.
package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter creates and configures a new Gin router with all task routes.
func SetupRouter(ctrl *controllers.Controller) *gin.Engine {
	// Create a new Gin router
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.POST("/register", ctrl.Register)
	r.POST("/login", ctrl.Login)

	// Define routes for task CRUD operations
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware()) 
	{
		taskRoutes := api.Group("/tasks") 
		{
			taskRoutes.GET("/", ctrl.GetTasks)
			taskRoutes.GET("/:id", ctrl.GetTask)
			taskRoutes.POST("/", ctrl.CreateTask)
			taskRoutes.PUT("/:id", ctrl.UpdateTask)
			taskRoutes.DELETE("/:id", ctrl.DeleteTask)
		}
		api.GET("/users", middleware.AdminMiddleware(),ctrl.GetUsers)
	}

	// Return the configured router
	return r
}
