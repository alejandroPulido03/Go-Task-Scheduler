package mem_storage

import (
	"task-scheduler/app/entities"
)

func NewTaskTreapStorage() *TaskTreap {
	return  &TaskTreap{}
}

func (t * TaskTreap) GetFirst() (*entities.Task, error) {
	if t.minNode == nil {
		return nil, nil
	}
	return t.minNode.task, nil
}

func (t * TaskTreap) GetMax() (*entities.Task, error) {
	if t.maxNode == nil {
		return nil, nil
	}
	return t.maxNode.task, nil
}

func (t * TaskTreap) Size() (int, error) {
	return t.size, nil
}

func (t * TaskTreap) SearchTask(task *entities.Task) (*entities.Task, error) {
	if node := t.search(task); node == nil {
		return nil, nil
	} else {
		return node.task, nil
	}
}

func (t * TaskTreap) PopTask(task *entities.Task) error {
	t.delete(task)
	return nil
}

func (t * TaskTreap) AddTask(task *entities.Task) error {
	t.insert(task)
	return nil
}

func (t * TaskTreap) PopNextTask() (*entities.Task, error) {
	min_node := t.minNode
	if min_node == nil {
		return nil, nil
	}
	task := min_node.task
	
	t.delete(task)
	return task, nil
}


func (t * TaskTreap) PopLastTask() (*entities.Task, error) {
	task := t.maxNode.task
	t.delete(task)
	return task, nil
}

func (t * TaskTreap) ReplaceLastTask(task *entities.Task) (*entities.Task, error) {
	oldTask := t.max()
	t.delete(oldTask)
	t.insert(task)
	return oldTask, nil
}