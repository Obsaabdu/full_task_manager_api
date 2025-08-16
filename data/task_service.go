// task_service.go
// Provides functions to read and write tasks to persistent storage (JSON file).
package data

import (
	"encoding/json"
	"errors"

	"os"
	"sync"

	"task_manager/models"
)

var filePath = "data/tasks.json"
var mutex sync.Mutex

// Read loads all tasks from the JSON file. Returns slice of tasks and error if reading or unmarshalling fails.
func ReadTasks() ([]models.Task, error) {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	defer data.Close()
	var tasks []models.Task

	info,_ := data.Stat()
	if info.Size() == 0 {
		return tasks, nil
	}

	err = json.NewDecoder(data).Decode(&tasks)
	return tasks, err
}

// Write saves the given slice of tasks to the JSON file. Returns error if marshalling or writing fails.
func WriteTasks(tasks []models.Task) error {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer data.Close()

	encoder := json.NewEncoder(data)
	encoder.SetIndent("", " ")

	return encoder.Encode(tasks)
}

func AddTask(task models.Task) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}
	task.ID = GetNextID(tasks)
	tasks = append(tasks, task)

	return WriteTasks(tasks)
}

func GetNextID(tasks []models.Task) int {
	if len(tasks) == 0 {
		return 1
	}

	maxID := tasks[0].ID

	for _, t := range tasks {
		if  t.ID > maxID {
			maxID = t.ID
		}

	}
	return maxID + 1
}


func GetTaskByID(id int) (*models.Task, error) {
	tasks, err := ReadTasks()
	if err != nil {
		return nil, err
	}
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("Task not found")
}

func UpdateTask(id int, updated models.Task) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	updatedList := make([]models.Task, 0)
	found := false
	for _, t := range tasks {
		if t.ID == id {
			updated.ID = id
			updatedList = append(updatedList, updated)
			found = true
		} else {
			updatedList = append(updatedList, t)
		}
	}

	if !found {
		return errors.New("Task not found")
	}

	return WriteTasks(updatedList)
}

func DeleteTask(id int) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	newList := make([]models.Task, 0)
	found := false
	for _, t := range tasks {
		if t.ID != id {
			newList = append(newList, t)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("Task not found")
	}

	return WriteTasks(newList)
}


