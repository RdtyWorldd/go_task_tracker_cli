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
	ID          int       ``
	Description string    ``
	Status      Progress  ``
	CreatedAt   time.Time ``
	UpdatedAt   time.Time ``
}
