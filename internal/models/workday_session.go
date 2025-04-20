package models

import "github.com/google/uuid"

type WorkdaySession struct {
	Id        uuid.UUID
	WorkdayId uuid.UUID
	StartedAt int64
	EndedAt   *int64
	Notes     string
}
