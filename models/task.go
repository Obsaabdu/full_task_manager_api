package models

import "time"

type Status string

const (
	StatusPending Status = "Pending"
	StatusInProgress Status = "In-progress"
	StatusCompleted Status = "Completed"
)
// Task represents a single task with its properties.
type Task struct {
	ID 				int 		`json:"id"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	DueDate 		time.Time 	`json:"due_date"`
	Status 			Status 		`json:"status"`
}
