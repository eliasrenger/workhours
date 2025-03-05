package models

import (
	"time"
	//"github.com/google/uuid"
)

type TimeSession struct {
	StartedAt   time.Time `json:"started_at"`
	FinnishedAt time.Time `json:"finnished_at"`
}

// Task which will be the db row
type Task struct {
	Id                uint64        `json:"id"` //equiv
	Name              string        `json:"name"`
	Duration          time.Duration `json:"duration"`
	EstimatedDuration time.Duration `json:"estimated_duration"`
	DurationDiff      time.Duration `json:"duration_diff"`
	StartedAt         time.Time     `json:"started_at"`
	FinnishedAt       time.Time     `json:"finnished_at"`
	Occurances        uint          `json:"occurances"`
	TimeSessions      []TimeSession `json:"time_sessions"`
	BlockedBy         []string      `json:"blocked_by"` // will not be used, just for the future
	Blocking          []string      `json:"blocking"`   // will not be used, just for the future
}

type WorkDay struct {
	Id                  uint64        `json:"id"`
	StartedAt           time.Time     `json:"started_at"`
	FinnishedAt         time.Time     `json:"finnished_at"`
	TasksWorkedOn       []string      `json:"tasks_worked_on"` // task uuid
	TimeSessions        []TimeSession `json:"time_sessions"`   // time session uuid
	NumberOfQuickBreaks uint          `json:"number_of_breaks"`
	LastQuickBreak      time.Time     `json:"last_quick_break"`
	BreakDuration       time.Duration `json:"break_duration"`
	WorkDuration        time.Duration `json:"work_duration"`
	Duration            time.Duration `json:"total_duration"`
}
