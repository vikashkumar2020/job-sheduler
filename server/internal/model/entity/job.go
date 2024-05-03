package model

import (
	"time"

	"github.com/google/uuid"
)

// Job struct having the details related to a job

type Job struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Duration  time.Duration `json:"duration"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Status    string      `json:"status"`
}