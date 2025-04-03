package task_utils

import (
	"encoding/json"
	"log"
	"os"
	"slices"

	"github.com/eliasrenger/workhours/internal/models"
)

func ReadTasks() []models.Task {
	var result []models.Task

	// Read the file
	data, err := os.ReadFile(cfg.TasksFilePath)
	if err != nil {
		log.Fatalln("failed to read tasks from file", cfg.TasksFilePath)
	}

	// Unmarshal the JSON
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalln("failed to unmarshal the json object")
	}

	return result
}

func AppendTask(newTask models.Task) {
	oldData := ReadTasks()
	tasks := append(oldData, newTask)
	SaveTasks(tasks)
}

func SaveTasks(tasks []models.Task) {
	formattedTasks, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatalln("failed formatting tasks")
	}
	err = os.WriteFile(cfg.TasksFilePath, formattedTasks, 0644)
	if err != nil {
		log.Fatalln("failed to write tasks to file")
	}
}

func EditTask(newTask models.Task) {
	tasks := ReadTasks()

	var cleanedTasks []models.Task
	var foundTargetTask bool
	for idx, task := range tasks {
		if task.Id == newTask.Id {
			cleanedTasks = slices.Delete(tasks, idx, idx+1)
			foundTargetTask = true
			break
		}
	}
	if !foundTargetTask {
		log.Fatalf("failed to find task %v in saved tasks", newTask.Name)
	}
	editedTasks := append(cleanedTasks, newTask)
	SaveTasks(editedTasks)
}
