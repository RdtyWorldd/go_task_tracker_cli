package actions

import (
	"github.com/RdtyWorldd/go_task_tracker_cli/dao"
)

type TaskAction struct {
	args []string
	dao  dao.Crud_dao
}

func NewTaskAction(_args []string, _dao dao.Crud_dao) TaskAction {
	return TaskAction{_args, _dao}
}

func (act TaskAction) add() {

}

func (act TaskAction) update() {

}

func (act TaskAction) delete() {

}

func (act TaskAction) Do() (int, error) {
	return 0, nil
}
