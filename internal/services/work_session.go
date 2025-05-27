package services

import (
	"errors"
	"time"

	"github.com/eliasrenger/workhours/internal/db"
	"github.com/eliasrenger/workhours/internal/models"
	"github.com/google/uuid"
)

// Define errors as package-level variables
var (
	ErrWorkSessionAlreadyActive = errors.New("work session already active")
	ErrNoWorkSessionActive      = errors.New("no active work session")
)

func StartWork(currentTime time.Time) error {
	hasActiveWorkSession, err := db.HasActiveWorkSession()
	if err != nil {
		return err
	}
	if hasActiveWorkSession {
		return ErrWorkSessionAlreadyActive
	}

	workSession := models.WorkSession{
		Id:                  uuid.New(),
		Date:                currentTime.Format("2006-01-02"),
		StartedAt:           currentTime.Unix(),
		EndedAt:             0,
		NumberOfQuickBreaks: 0,
		LastQuickBreak:      0,
		Notes:               "",
	}

	err = db.InsertWorkSession(workSession)
	if err != nil {
		return err
	}

	return nil
}

func StopWork(currentTime time.Time) error {
	activeWorkSession, err := db.GetActiveWorkSession()
	if err != nil {
		return err
	}
	if activeWorkSession.Id == uuid.Nil {
		return ErrNoWorkSessionActive
	}

	activeWorkSession.EndedAt = currentTime.Unix()
	err = db.UpdateWorkSessionById(activeWorkSession)
	if err != nil {
		return err
	}

	return nil
}

func AddQuickBreak(currentTime time.Time) error {
	activeWorkSession, err := db.GetActiveWorkSession()
	if err != nil {
		return err
	}
	if activeWorkSession.Id == uuid.Nil {
		return ErrNoWorkSessionActive
	}

	activeWorkSession.NumberOfQuickBreaks++
	activeWorkSession.LastQuickBreak = currentTime.Unix()
	err = db.UpdateWorkSessionById(activeWorkSession)

	if err != nil {
		return err
	}

	return nil
}

func GetSecondsWorkedToday() (int64, error) {
	currentTime := time.Now()
	secondsWorked, err := db.GetWorktimeByDate(currentTime.Format("2006-01-02"))
	if err != nil {
		return 0, err
	}

	return secondsWorked, nil
}
