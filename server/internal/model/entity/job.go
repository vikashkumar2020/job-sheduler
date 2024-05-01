package model

import (
	"time"

	"github.com/google/uuid"
)

// Job struct having the details related to a job
type Job struct {
	ID        uuid.UUID
	Name      string
	Duration  time.Duration
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}
