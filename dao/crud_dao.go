package dao

import (
	"github.com/RdtyWorldd/go_task_tracker_cli/task"
)

type Crud_dao interface {
	Create(task.Task)
	Read(id int) (task.Task, error)
	Update(id int, upd_task task.Task) error
	Delete(id int) error
	ReadAll() []task.Task
}
