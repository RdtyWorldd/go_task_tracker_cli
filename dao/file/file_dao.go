package file

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/RdtyWorldd/go_task_tracker_cli/task"
)

type FileDao struct {
	path    string
	taskMap map[int]task.Task
}

func NewFileDao(path string) FileDao {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var tasks []task.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	dao := FileDao{path, make(map[int]task.Task)}
	for i, value := range tasks {
		dao.taskMap[i] = value
	}
	return dao
}

// question
// нужно ли проверять индекс или доверяться обработчикам комманд
func (dao FileDao) Create(task task.Task) {
	file, err := os.OpenFile(dao.path, os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Seek(-1, 2)
	task_json, err := json.Marshal(task)
	if err != nil {

	}
	write_data := string(task_json) + "]"
	if len(dao.taskMap) != 0 {
		write_data = "," + write_data
	}
	io.WriteString(file, write_data)
	//file.WriteString(write_data)
}

func (dao FileDao) Read(id int) (task.Task, error) {
	if id < 0 || id > len(dao.taskMap) {
		return task.Task{}, errors.New("index out of range") //пусть пока повисит пустая таска
	}
	return dao.taskMap[id], nil
}

func (dao FileDao) Update(task task.Task) error {
	return nil
}

func (dao FileDao) Delete(id int) error {
	return nil
}
