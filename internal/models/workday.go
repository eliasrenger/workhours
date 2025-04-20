package models

import "github.com/google/uuid"

type Workday struct {
	Id                  uuid.UUID
	Date                string // ISO format (e.g. "2025-04-19")
	IsActive            bool
	NumberOfQuickBreaks int
	LastQuickBreak      *int64 // INTEGER (Unix timestamp)
	Notes               string
}
