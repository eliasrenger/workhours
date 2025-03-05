package tasks_utils

import (
	"example.com/workhours/config"
	"example.com/workhours/internal/models"
)

var cfg config.Config = config.LoadConfig()

func GetActiveTask() (models.Task, bool, error) {
	var failedReturn models.Task
	tasks, err := ReadTasks(cfg.TasksFilePath)
	if err != nil {
		return failedReturn, false, err
	}
	for _, task := range tasks {
		if task.FinnishedAt.Year() == 1 && task.StartedAt.Year() != 1 {
			return task, true, nil
		}
	}
	return failedReturn, false, nil
}

func IsTaskActive(task models.Task) bool {
	if task.FinnishedAt.Year() == 1 {
		return true
	} else {
		return false
	}
}

func IsLastSessionActive(task models.Task) bool {
	lastSession := task.TimeSessions[len(task.TimeSessions)-1]
	if lastSession.FinnishedAt.Year() == 1 {
		return true
	} else {
		return false
	}
}

func UpdateTask(task models.Task) {}