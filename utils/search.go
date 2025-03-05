package utils

import (
	"example.com/workhours/config"
	"example.com/workhours/internal/models"
)

var cfg config.Config = config.LoadConfig()

func GetOngoingWorkDay() (models.WorkDay, error) {
	var failedReturn models.WorkDay
	workDayData, err := ReadWorkDays(cfg.WorkDaysFilePath)
	if err != nil {
		return failedReturn, err
	}
	for _, workDay := range workDayData {
		if workDay.FinnishedAt.Year() == 1 {
			return workDay, nil
		}
	}
	return failedReturn, nil
}

func GetOngoingTask() (models.Task, error) {
	var failedReturn models.Task
	tasks, err := ReadTasks(cfg.TasksFilePath)
	if err != nil {
		return failedReturn, err
	}
	for _, task := range tasks {
		if task.Ongoing {
			return task, nil
		}
	}
	return failedReturn, nil
}
