package models

import "github.com/google/uuid"

type WorkSession struct {
	Id                  uuid.UUID
	Date                string // ISO format (e.g. "2025-04-19")
	StartedAt           int64
	EndedAt             *int64
	NumberOfQuickBreaks int
	LastQuickBreak      *int64
	Notes               string
}
