package actions

import (
	"time"

	"github.com/RdtyWorldd/go_task_tracker_cli/dao"
	"github.com/RdtyWorldd/go_task_tracker_cli/task"
)

type TaskAction struct {
	args []string
	dao  dao.Crud_dao
}

func NewTaskAction(_args []string, _dao dao.Crud_dao) TaskAction {
	return TaskAction{_args, _dao}
}

func (act TaskAction) add(desc string) error {
	// if len(act.args) < 3 {
	// 	return errors.New("there are too few arguments")
	// }
	tasks := act.dao.ReadAll()
	id := len(tasks) + 1
	status := task.TODO
	created := time.Now()
	updated := created

	act.dao.Create(task.Task{
		ID:          id,
		Description: desc,
		Status:      status,
		CreatedAt:   created,
		UpdatedAt:   updated})
	return nil
}

func (act TaskAction) update(id int, upd_desc string) error {
	task, err := act.dao.Read(id)
	if err != nil {
		panic(err)
	}
	task.Description = upd_desc
	task.UpdatedAt = time.Now()
	act.dao.Update(id, task)
	return nil
}

func (act TaskAction) delete(id int) error {
	err := act.dao.Delete(id)
	return err
}

func (act TaskAction) Do() (int, error) {
	return 0, nil
}
