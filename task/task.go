package task

import (
	"fmt"
	"time"
)

type Progress string

const (
	TODO       Progress = "todo"
	INPROGRESS Progress = "in-progress"
	DONE       Progress = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Progress  `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t Task) Print() {
	format := "task id:%d\n\t%s\n\tStatus: %s\n\tCreated At: %s\n\tUpdated At: %s\n"
	fmt.Printf(format, t.ID,
		t.Description,
		t.Status,
		t.CreatedAt.Format("Monday, 2.January.2006, at 15:04"),
		t.UpdatedAt.Format("Monday, 2.January.2006, at 15:04"))
}
