package main

import (
	"fmt"
	"os"

	"github.com/RdtyWorldd/go_task_tracker_cli/actions"
	"github.com/RdtyWorldd/go_task_tracker_cli/dao/file"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Arguments incorect. \n Use -h flag to get help information")
	}
	_, err := os.Stat("./json")
	if err != nil {
		os.Mkdir("json", 0755)
	}

	var action actions.Action
	if args[1] == actions.HELP {
		action = actions.NewHelpAction(args[0])
		err := action.Do()
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
	action = actions.NewTaskAction(args, file.NewFileDao("json/tasks.json"))
	err = action.Do()
	if err != nil {
		fmt.Println(err.Error())
	}
}
