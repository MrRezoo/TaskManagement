package models

import (
	"github.com/google/uuid"
	"time"
)

// Status represents the status of a task.
type Status string

const (
	StatusPending    Status = "pending"
	StatusCompleted  Status = "completed"
	StatusInProgress Status = "in_progress"
)

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Status      Status    `json:"status" validate:"required,oneof=pending completed in_progress"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
