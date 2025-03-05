package tasks_utils

import (
	"encoding/json"
	"os"
	"slices"

	"example.com/workhours/internal/models"
)

func ReadTasks(filePath string) ([]models.Task, error) {
	var result []models.Task

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return result, err
	}

	// Unmarshal the JSON
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func AppendTasks(filePath string, newTasks []models.Task) error {
	oldData, err := ReadTasks(filePath)
	if err != nil {
		return err
	}
	tasks := append(oldData, newTasks...)
	return SaveTasks(filePath, tasks)
}

func SaveTasks(filePath string, tasks []models.Task) error {
	formattedTasks, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, formattedTasks, 0644)
}

func EditTask(filePath string, newTask models.Task) error {
	tasks, err := ReadTasks(filePath)
	if err != nil {
		return err
	}
	var editedTasks []models.Task
	var cleanedTasks []models.Task
	var foundTargetTask bool
	for idx, task := range tasks {
		if task.Id == newTask.Id {
			cleanedTasks = slices.Delete(tasks, idx, idx)
			foundTargetTask = true
			break
		}
	}
	if !foundTargetTask {
		var Error error
		return Error
	}
	editedTasks = append(cleanedTasks, newTask)
	return SaveTasks(filePath, editedTasks)
}
