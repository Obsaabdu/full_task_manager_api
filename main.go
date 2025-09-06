// main.go
// Entry point for the Task Manager API application. Sets up the router and starts the server.
package main

import (
	"context"
	"log"
	"os"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// main initializes the router and starts the HTTP server.
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("Task_db")

	userService := &data.UserService{Collection: db.Collection("users")}
	taskService := &data.TaskService{Collection: db.Collection("tasks")}
	ctrl := &controllers.Controller{UserService: userService, TaskService: taskService}
	
	r := router.SetupRouter(ctrl)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on :"+port)
	if err := r.Run(":"+port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
