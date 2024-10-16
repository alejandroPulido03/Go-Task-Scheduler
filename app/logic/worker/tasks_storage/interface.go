package task_storage

import (
	"task-scheduler/app/entities"
)

type Tasks interface {
	AddTask(task *entities.Task) error
	PopNextTask() (*entities.Task, error)
	PopLastTask() (*entities.Task, error)
	PopTask(task *entities.Task) error
	SearchTask(task *entities.Task) (*entities.Task, error)
	Size() (int, error)
	Min() (*entities.Task, error)
	Max() (*entities.Task, error)
}

func NewTaskTreapStorage() TaskTreap {
	return TaskTreap{}
}

func (t *TaskTreap) Min() (*entities.Task, error) {
	return t.min(), nil
}

func (t *TaskTreap) Max() (*entities.Task, error) {
	return t.max(), nil
}

func (t *TaskTreap) Size() int {
	return t.size
}

func (t *TaskTreap) SearchTask(task *entities.Task) (*entities.Task, error) {
	if node := t.search(task); node == nil {
		return nil, nil
	} else {
		return node.task, nil
	}
}

func (t *TaskTreap) PopTask(task *entities.Task) error {
	t.delete(task)
	return nil
}

func (t *TaskTreap) AddTask(task *entities.Task) error {
	t.insert(task)
	return nil
}

func (t *TaskTreap) PopNextTask() (*entities.Task, error) {
	task := t.min()
	t.delete(task)
	return task, nil
}


func (t *TaskTreap) PopLastTask() (*entities.Task, error) {
	task := t.max()
	t.delete(task)
	return task, nil
}

func (t *TaskTreap) ReplaceLastTask(task *entities.Task) (*entities.Task, error) {
	oldTask := t.max()
	t.delete(oldTask)
	t.insert(task)
	return oldTask, nil
}