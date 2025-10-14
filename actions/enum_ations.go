package actions

const (
	ADD         = "add"
	UPD         = "update"
	DEL         = "delete"
	IN_PROGRESS = "mark-in-progress"
	DONE        = "mark-done"
	LIST        = "list"
	HELP        = "-h"
)

type Action interface {
	Do() error
}
