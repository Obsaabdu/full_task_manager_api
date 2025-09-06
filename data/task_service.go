// task_service.go
// Provides functions to read and write tasks to persistent storage (JSON file).
package data

import (
	"context"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	Collection *mongo.Collection
}

func (s *TaskService) CreateTask(task models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.Collection.InsertOne(ctx, task)
	
	return err
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}

	var task models.Task
	err = s.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	return task, err
}

func (s *TaskService) UpdateTask(id string, task models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":		task.Title,
			"description": 	task.Description,
			"due_date":   	task.DueDate,
			"status": 		task.Status,
		},
	}
	_, err = s.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (s *TaskService) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}