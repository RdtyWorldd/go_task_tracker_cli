package actions

import (
	"fmt"
)

type HelpAction struct {
	program_name string
}

func NewHelpAction(_program_name string) *HelpAction {
	return &HelpAction{program_name: _program_name}
}

func (act *HelpAction) Do() error {
	help_text :=
		`usage %s:
		list - show all tasks
		list [todo, in-progress, done] - show all tasks with certain progress status

		add [task description] - add task to the task list
		update [task ID] [updated task description] - update description of task with certain ID
		delete [task ID] - delete task with certain ID from the task list
		
		mark-in-progress [task ID] - change progress status to IN-PROGRESS of task with certain ID
		mark-done [task ID] - change progress status to DONE of task with certain ID
	`
	fmt.Printf(help_text, act.program_name)
	return nil
}
