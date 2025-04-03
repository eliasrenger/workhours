package task_utils

import (
	"log"
	"time"

	"github.com/eliasrenger/workhours/config"
	"github.com/eliasrenger/workhours/internal/models"
)

var cfg config.Config = config.LoadConfig()

func GetActiveTask() (models.Task, bool) {
	var failedReturn models.Task
	tasks := ReadTasks()
	for _, task := range tasks {
		if IsTaskActive(task) {
			return task, true
		}
	}
	return failedReturn, false
}

func GetTaskByName(name string) (models.Task, bool) {
	var failedReturn models.Task
	tasks := ReadTasks()
	for _, task := range tasks {
		if task.Name == name {
			return task, true
		}
	}
	return failedReturn, false
}

func IsTaskFinished(task models.Task) bool {
	if task.FinishedAt.Year() == 1 {
		return false
	} else {
		return true
	}
}

func IsTaskActive(task models.Task) bool {
	lastSession := task.TimeSessions[len(task.TimeSessions)-1]
	if lastSession.FinishedAt.Year() == 1 {
		return true
	} else {
		return false
	}
}

func IsLastSessionActive(task models.Task) bool {
	lastSession := task.TimeSessions[len(task.TimeSessions)-1]
	if lastSession.FinishedAt.Year() == 1 {
		return true
	} else {
		return false
	}
}

func UpdateTask(task models.Task) models.Task {
	currentTime := time.Now()

	// WorkDuration
	var taskDuration time.Duration
	var lastDuration time.Duration

	timeSessions := task.TimeSessions
	lastSessionIdx := len(timeSessions) - 1
	lastSession := timeSessions[lastSessionIdx]

	if IsLastSessionActive(task) {
		lastDuration = currentTime.Sub(lastSession.StartedAt)
	} else {
		lastDuration = lastSession.FinishedAt.Sub(lastSession.StartedAt)
	}
	for idx, timeSession := range timeSessions {
		if idx == lastSessionIdx {
			break
		}
		sessionDuration := timeSession.FinishedAt.Sub(timeSession.StartedAt)
		taskDuration += sessionDuration
	}
	taskDuration += lastDuration
	task.Duration = taskDuration
	task.DurationDiff = taskDuration - task.EstimatedDuration

	task.WorkCount = uint(len(task.TimeSessions))

	return task
}

func PauseActiveTask(name string, currentTime time.Time) {
	activeTask, ok := GetActiveTask()
	if !ok || activeTask.Name != name {
		log.Fatalln("there is no active task with name:", name)
	}
	activeTask.TimeSessions[len(activeTask.TimeSessions)-1].FinishedAt = currentTime
	updatedTask := UpdateTask(activeTask)
	EditTask(updatedTask)
}
