package work_day_utils

import (
	"time"

	"example.com/workhours/config"
	"example.com/workhours/internal/models"
)

var cfg config.Config = config.LoadConfig()

func GetActiveWorkDay() (models.WorkDay, bool, error) {
	var failedReturn models.WorkDay
	workDayData, err := ReadWorkDays(cfg.WorkDaysFilePath)
	if err != nil {
		return failedReturn, false, err
	}
	for _, workDay := range workDayData {
		if workDay.FinnishedAt.Year() == 1 && workDay.StartedAt.Year() != 1 {
			return workDay, true, nil
		}
	}
	return failedReturn, false, nil
}

func IsWorkDayActive(workDay models.WorkDay) bool {
	if workDay.FinnishedAt.Year() == 1 {
		return true
	} else {
		return false
	}
}

func IsLastSessionActive(workDay models.WorkDay) bool {
	lastSession := workDay.TimeSessions[len(workDay.TimeSessions)-1]
	if lastSession.FinnishedAt.Year() == 1 {
		return true
	} else {
		return false
	}
}

func UpdateTask(workDay models.WorkDay) {
	currentTime := time.Now()
	// Duration
	var duration time.Duration
	if !IsWorkDayActive(workDay) {
		duration = workDay.FinnishedAt.Sub(workDay.StartedAt)
	} else {
		duration = currentTime.Sub(workDay.StartedAt)
	}
	workDay.Duration = duration

	// WorkDuration
	var workDuration time.Duration
	var lastDuration time.Duration

	timeSessions := workDay.TimeSessions
	lastSessionIdx := len(timeSessions) - 1
	lastSession := timeSessions[lastSessionIdx]

	if IsLastSessionActive(workDay) {
		lastDuration = currentTime.Sub(lastSession.StartedAt)
	} else {
		lastDuration = lastSession.FinnishedAt.Sub(lastSession.StartedAt)
	}
	for idx, timeSession := range timeSessions {
		if idx == lastSessionIdx {
			break
		}
		duration := timeSession.FinnishedAt.Sub(timeSession.StartedAt)
		workDuration += duration
	}
	workDuration += lastDuration
	workDay.WorkDuration = workDuration

	// Break duration
	breakDuration := duration - workDuration
	workDay.BreakDuration = breakDuration
}
