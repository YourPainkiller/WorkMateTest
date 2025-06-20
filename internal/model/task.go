package model

import "time"

type Status string

const (
	StatusCreated   Status = "created"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Task struct {
	ID          string
	Status      Status
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
	Result      string
	Error       error
}

func (t Task) Duration() time.Duration {
	switch t.Status {
	case StatusCompleted, StatusFailed:
		return t.CompletedAt.Sub(t.StartedAt)
	case StatusRunning:
		return time.Since(t.StartedAt)
	default:
		return 0
	}
}
