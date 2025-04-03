package work_day_utils

import (
	"time"

	"github.com/eliasrenger/workhours/config"
	"github.com/eliasrenger/workhours/internal/models"
)

var cfg config.Config = config.LoadConfig()

func GetActiveWorkDay() (models.WorkDay, bool) {
	var failedReturn models.WorkDay
	workDayData := ReadWorkDays()
	for _, workDay := range workDayData {
		if workDay.FinishedAt.Year() == 1 && workDay.StartedAt.Year() != 1 {
			return workDay, true
		}
	}
	return failedReturn, false
}

func IsWorkDayActive(workDay models.WorkDay) bool {
	if workDay.FinishedAt.Year() == 1 {
		return true
	} else {
		return false
	}
}

func IsLastSessionActive(workDay models.WorkDay) bool {
	lastSession := workDay.TimeSessions[len(workDay.TimeSessions)-1]
	if lastSession.FinishedAt.Year() == 1 {
		return true
	} else {
		return false
	}
}

func UpdateWorkDay(workDay models.WorkDay) models.WorkDay {
	currentTime := time.Now()
	// Duration
	var duration time.Duration
	if !IsWorkDayActive(workDay) {
		duration = workDay.FinishedAt.Sub(workDay.StartedAt)
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
		lastDuration = lastSession.FinishedAt.Sub(lastSession.StartedAt)
	}
	for idx, timeSession := range timeSessions {
		if idx == lastSessionIdx {
			break
		}
		duration := timeSession.FinishedAt.Sub(timeSession.StartedAt)
		workDuration += duration
	}
	workDuration += lastDuration
	workDay.WorkDuration = workDuration

	// Break duration
	breakDuration := duration - workDuration
	workDay.BreakDuration = breakDuration

	return workDay
}
