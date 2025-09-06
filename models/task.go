package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


// Task represents a single task with its properties.
type Task struct {
	ID 				primitive.ObjectID	`json:"id,omitempty" bson:"_id,omitempty"`
	Title 			string 				`json:"title"`
	Description 	string 				`json:"description"`
	DueDate 		time.Time 			`json:"due_date"`
	Status 			string 				`json:"status"`
}
