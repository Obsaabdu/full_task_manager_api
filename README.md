# Task Manager API

A RESTful API for managing tasks, built with Go, Gin, and MongoDB.

## Features
- Create, read, update, and delete tasks
- Task status enum: Pending, In-progress, Completed
- MongoDB ObjectID for unique task IDs
- DTOs for request/response validation

## Technologies
- Go
- Gin (web framework)
- MongoDB (database)
- MongoDB Go Driver

## Getting Started

### Prerequisites
- Go 1.18+
- MongoDB running locally or remotely

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Obsaabdu/full_task_manager_api.git
   cd full_task_manager_api
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set MongoDB connection string in your environment or config.

### Running the API
```bash
# Start the server
 go run main.go
```

The API will be available at `http://localhost:8080`.

## API Endpoints
See [docs/api_documentation.md](docs/api_documentation.md) for full details.

## Example Usage
```bash
# Get all tasks
curl http://localhost:8080/tasks

# Create a task
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Test","description":"Test task","due_date":"2025-09-02T12:00:00Z"}'
```

## Task Model
```go
import "go.mongodb.org/mongo-driver/bson/primitive"

// Status enum
const (
    StatusPending     Status = "Pending"
    StatusInProgress  Status = "In-progress"
    StatusCompleted   Status = "Completed"
)

type Task struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title       string             `json:"title"`
    Description string             `json:"description"`
    DueDate     time.Time          `json:"due_date"`
    Status      Status             `json:"status"`
}
```

## License
MIT
