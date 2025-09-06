# Task Manager API Documentation

## Overview
The Task Manager API is a RESTful service built with Go and Gin for managing tasks. It allows clients to create, read, update, and delete tasks, each with properties such as title, description, due date, and status. Task status is an enum: Pending, In-progress, or Completed. IDs are MongoDB ObjectIDs and request/response formats use DTOs. Data is stored in MongoDB.

## Base URL
```
http://localhost:8080
```


## Authentication

All protected endpoints require a valid JWT token in the `Authorization` header:

```
Authorization: Bearer <token>
```

Obtain a token by registering and logging in:

- `POST /register` — Register a new user
- `POST /login` — Login and receive a JWT

## Endpoints

### Get All Tasks
- **URL:** `/api/tasks/`
- **Method:** `GET`
- **Auth:** Required (JWT)
- **Description:** Returns a list of all tasks.
- **Response:**
  - `200 OK`: Array of task objects

### Get Task by ID
- **URL:** `/api/tasks/:id`
- **Method:** `GET`
- **Auth:** Required (JWT)
- **Description:** Returns a single task by its MongoDB ObjectID.
- **Response:**
  - `200 OK`: Task object
  - `404 Not Found`: Task not found

### Create Task
- **URL:** `/api/tasks/`
- **Method:** `POST`
- **Auth:** Required (JWT)
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
- **URL:** `/api/tasks/:id`
- **Method:** `PUT`
- **Auth:** Required (JWT)
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
- **URL:** `/api/tasks/:id`
- **Method:** `DELETE`
- **Auth:** Required (JWT)
- **Description:** Deletes a task by ObjectID.
- **Response:**
  - `200 OK`: Success message
  - `404 Not Found`: Task not found
### Get All Users (Admin Only)
- **URL:** `/api/users`
- **Method:** `GET`
- **Auth:** Required (JWT, Admin role)
- **Description:** Returns a list of all users. Only accessible to users with the Admin role.
- **Response:**
  - `200 OK`: Array of user objects
  - `403 Forbidden`: Admins only

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
- Register a user:
  ```bash
  curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"yourpassword"}'
  ```
- Login and get JWT:
  ```bash
  curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"yourpassword"}'
  ```
- Get all tasks (with JWT):
  ```bash
  curl -H "Authorization: Bearer <token>" http://localhost:8080/api/tasks/
  ```
- Create a task (with JWT):
  ```bash
  curl -X POST http://localhost:8080/api/tasks/ -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title":"Test","description":"Test task","due_date":"2025-09-06T12:00:00Z"}'
  ```
- Update a task (with JWT):
  ```bash
  curl -X PUT http://localhost:8080/api/tasks/<id> -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title":"Updated","status":"Completed"}'
  ```
- Get all users (admin only):
  ```bash
  curl -H "Authorization: Bearer <admin_token>" http://localhost:8080/api/users
  ```
  ```bash
  curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title":"Updated","status":"Completed"}'
  ```

## Notes
- All data is stored in MongoDB.
- The API is suitable for local development and small-scale use.
