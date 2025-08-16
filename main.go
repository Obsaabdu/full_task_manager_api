// main.go
// Entry point for the Task Manager API application. Sets up the router and starts the server.
package main

import "task_manager/router"

// main initializes the router and starts the HTTP server.
func main() {
	r := router.SetupRouter()
	r.Run()
}
