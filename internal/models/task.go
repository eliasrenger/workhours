package models

import "github.com/google/uuid"

type Task struct {
	Id                uuid.UUID
	Title             string
	Description       string
	Priority          int // 1-5
	Category          string
	CreatedAt         int64
	CompletedAt       *int64 // INTEGER (Unix timestamp)
	Status            string
	EstimatedDuration int
	Notes             string
}
