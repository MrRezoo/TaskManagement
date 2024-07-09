package models

import (
	"github.com/google/uuid"
	"time"
)

// Status represents the status of a task.
type Status string

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status" validate:"required,oneof=pending completed failed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
