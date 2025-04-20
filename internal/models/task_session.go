package models

import "github.com/google/uuid"

type TaskSession struct {
	Id        uuid.UUID
	TaskId    uuid.UUID
	StartedAt int64
	EndedAt   *int64
	Notes     string
}
