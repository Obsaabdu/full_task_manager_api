// main.go
// Entry point for the Task Manager API application. Sets up the router and starts the server.
package main

import (
	"log"
	"os"
	"task_manager/config"
	"task_manager/router"

	"github.com/joho/godotenv"
)

// main initializes the router and starts the HTTP server.
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	config.ConnectDB()
	
	r := router.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":"+port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
