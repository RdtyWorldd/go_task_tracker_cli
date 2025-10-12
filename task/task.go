package task

import (
	"time"
)

type Progress int

const (
	TODO Progress = iota
	INPROGRESS
	DONE
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Progress  `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
