package dao

import (
	"github.com/RdtyWorldd/go_task_tracker_cli/task"
)

type Crud_dao interface {
	Create(task.Task)
	Read(id int) (task.Task, error)
	Update(task.Task) error
	Delete(id int) error
}
