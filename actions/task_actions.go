package actions

import (
	"errors"
	"fmt"
	"strconv"
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
	tasks := act.dao.ReadAll()
	id := len(tasks) + 1
	status := task.TODO
	created := time.Now()
	updated := created

	err := act.dao.Create(task.Task{
		ID:          id,
		Description: desc,
		Status:      status,
		CreatedAt:   created,
		UpdatedAt:   updated})
	if err != nil {
		return err
	}
	fmt.Printf("Task added sucessfully (ID: %d)\n", id)
	return nil
}

func (act TaskAction) update(id int, upd_desc string) error {
	task, err := act.dao.Read(id)
	if err != nil {
		return err
	}
	task.Description = upd_desc
	task.UpdatedAt = time.Now()
	err = act.dao.Update(id, task)
	if err != nil {
		return err
	}
	fmt.Printf("Task updateded sucessfully (ID: %d)\n", id)
	return nil
}

func (act TaskAction) delete(id int) error {
	err := act.dao.Delete(id)
	if err != nil {
		return err
	}
	fmt.Printf("Task delete sucessfully (ID: %d)\n", id)
	return nil
}

func (act TaskAction) list(progress any) {
	task_list := act.dao.ReadAll()
	if len(task_list) == 0 {
		fmt.Println("There are no tasks in your list.")
		return
	}
	var res_list []task.Task
	if value, ok := progress.(task.Progress); ok {
		for _, t := range task_list {
			if t.Status == value {
				res_list = append(res_list, t)
			}
		}
	} else {
		res_list = task_list
	}
	for _, t := range res_list {
		t.Print()
	}
}

func (act TaskAction) mark(id int, progress task.Progress) error {
	mark_task, _ := act.dao.Read(id)
	mark_task.Status = progress
	err := act.dao.Update(id, mark_task)
	if err != nil {
		return err
	}
	fmt.Printf("Task progress udated (ID: %d)\n", mark_task.ID)
	return nil
}

func (act TaskAction) Do() error {
	var e error
	switch act.args[1] {
	case ADD:
		{
			if len(act.args) < 3 {
				return errors.New("there are too few arguments")
			}
			e = act.add(string(act.args[2]))
		}
	case UPD:
		{
			if len(act.args) < 4 {
				return errors.New("there are too few arguments")
			}
			id, err := strconv.Atoi(act.args[2])
			if err != nil {
				return err
			}
			upd_desc := string(act.args[3])
			e = act.update(id, upd_desc)
		}
	case DEL:
		{
			id, err := strconv.Atoi(act.args[2])
			if err != nil {
				return err
			}
			e = act.delete(id)
		}
	case IN_PROGRESS:
		{
			if len(act.args) < 3 {
				return errors.New("there are too few arguments")
			}
			id, err := strconv.Atoi(act.args[2])
			if err != nil {
				return err
			}
			act.mark(id, task.INPROGRESS)
		}
	case DONE:
		{
			if len(act.args) < 3 {
				return errors.New("there are too few arguments")
			}
			id, err := strconv.Atoi(act.args[2])
			if err != nil {
				return err
			}
			act.mark(id, task.DONE)
		}
	case LIST:
		{
			if len(act.args) == 3 {
				if act.args[2] == string(task.DONE) || act.args[2] == string(task.INPROGRESS) || act.args[2] == string(task.TODO) {
					act.list(task.Progress(act.args[2]))
				} else {
					fmt.Printf("Unknown parameter: %s, use this key words: (%s, %s, %s)\n",
						act.args[2], task.DONE, task.INPROGRESS, task.TODO)
				}
			} else {
				act.list(nil)
			}
		}
	default:
		fmt.Println("Unkniwn command. Try -h flag to wath all available command")
	}
	return e
}
