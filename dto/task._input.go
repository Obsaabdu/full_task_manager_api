package dto

import (
	"time"
	"task_manager/models"
)

type CreateTaskInput struct {
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	DueDate 		time.Time 	`json:"due_date"`
}

type UpdateTaskInput struct {
	Title 			string 				`json:"title"`
	Description 	string 				`json:"description"`
	DueDate 		time.Time 			`json:"due_date"`
	Status 			models.Status 		`json:"status"`
}
