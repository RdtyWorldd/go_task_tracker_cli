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
	dao := FileDao{path, make(map[int]task.Task)}
	if len(data) != 0 {
		var tasks []task.Task
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			panic(err)
		}
		for i, value := range tasks {
			dao.taskMap[i] = value
		}
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
		panic(err)
	}
	write_data := string(task_json) + "]"
	if len(dao.taskMap) != 0 {
		write_data = "," + write_data
	} else {
		write_data = "[" + write_data
	}
	_, err = io.WriteString(file, write_data)
	if err != nil {
		panic(err)
	}
}

func (dao FileDao) Read(id int) (task.Task, error) {
	if id < 0 || id > len(dao.taskMap) {
		return task.Task{}, errors.New("index out of range") //пусть пока повисит пустая таска
	}
	if value, ok := dao.taskMap[id]; ok {
		return value, nil
	} else {
		return task.Task{}, errors.New("index out of range")
	}
}

func (dao FileDao) ReadAll() []task.Task {
	res := make([]task.Task, 0, len(dao.taskMap))
	for _, value := range dao.taskMap {
		res = append(res, value)
	}
	return res
}

func (dao FileDao) Update(id int, upd_task task.Task) error {
	if id < 0 || id > len(dao.taskMap) {
		return errors.New("index out of range") //пусть пока повисит пустая таска
	}
	dao.taskMap[id] = upd_task
	file, err := os.OpenFile(dao.path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(dao.taskMap)
	if err != nil {
		return err
	}
	_, err = io.WriteString(file, string(data))
	if err != nil {
		return err
	}
	return nil
}

func (dao FileDao) Delete(id int) error {
	if id < 0 || id > len(dao.taskMap) {
		return errors.New("index out of range") //пусть пока повисит пустая таска
	}
	delete(dao.taskMap, id)
	file, err := os.OpenFile(dao.path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(dao.taskMap)
	if err != nil {
		return err
	}
	_, err = io.WriteString(file, string(data))
	if err != nil {
		return err
	}
	return nil
}
