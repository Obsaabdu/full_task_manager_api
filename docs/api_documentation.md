# Task Manager API Documentation

## Overview
The Task Manager API is a RESTful service built with Go and Gin for managing tasks. It allows clients to create, read, update, and delete tasks, each with properties such as title, description, due date, and status. Task status is an enum: Pending, In-progress, or Completed. IDs are MongoDB ObjectIDs and request/response formats use DTOs. Data is stored in MongoDB.

## Base URL
```
http://localhost:8080
```

## Endpoints

### Get All Tasks
- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Returns a list of all tasks.
- **Response:**
  - `200 OK`: Array of task objects

### Get Task by ID
- **URL:** `/tasks/:id`
- **Method:** `GET`
- **Description:** Returns a single task by its MongoDB ObjectID.
- **Response:**
  - `200 OK`: Task object
  - `404 Not Found`: Task not found

### Create Task
- **URL:** `/tasks`
- **Method:** `POST`
- **Description:** Creates a new task. The ID and status are set automatically.
- **Request Body:**
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "YYYY-MM-DDTHH:MM:SSZ"
  }
  ```
- **Response:**
  - `201 Created`: Created task object (with ObjectID)
  - `400 Bad Request`: Invalid input

### Update Task
- **URL:** `/tasks/:id`
- **Method:** `PUT`
- **Description:** Updates an existing task by ObjectID. You may update title, description, due date, and status.
- **Request Body:**
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "YYYY-MM-DDTHH:MM:SSZ",
    "status": "Pending" | "In-progress" | "Completed"
  }
  ```
- **Response:**
  - `200 OK`: Success message
  - `404 Not Found`: Task not found

### Delete Task
- **URL:** `/tasks/:id`
- **Method:** `DELETE`
- **Description:** Deletes a task by ObjectID.
- **Response:**
  - `200 OK`: Success message
  - `404 Not Found`: Task not found

## Task Model
```go
type Status string

const (
  StatusPending     Status = "Pending"
  StatusInProgress  Status = "In-progress"
  StatusCompleted   Status = "Completed"
)

type Task struct {
  ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // MongoDB ObjectID
  Title       string             `json:"title"` // Name or title
  Description string             `json:"description"` // Details
  DueDate     time.Time          `json:"due_date"` // Due date
  Status      Status             `json:"status"` // Status (Pending, In-progress, Completed)
}
```

## Error Handling
All errors are returned as JSON objects with an `error` field describing the issue.

## Example Usage
- Get all tasks:
  ```bash
  curl http://localhost:8080/tasks
  ```
- Create a task:
  ```bash
  curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Test","description":"Test task","due_date":"2025-08-17T12:00:00Z"}'
  ```
- Update a task:
  ```bash
  curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title":"Updated","status":"Completed"}'
  ```

## Notes
- All data is stored in MongoDB.
- The API is suitable for local development and small-scale use.
