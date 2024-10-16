package task_storage

import (
	"task-scheduler/app/entities"
	"time"
)

type Tasks interface{
	AddTask(task *entities.Task)
	PopNextTask() *entities.Task
	SeekNextTaskTime() string
	PopLastTask() *entities.Task
	PopTask(task *entities.Task)
	SearchTask(task *entities.Task) *entities.Task
	Min() *entities.Task
	Max() *entities.Task
}

func NewTaskTreapStorage() TaskTreap{
	return TaskTreap{}
}

func (t *TaskTreap) Min() *entities.Task{
	return t.min()
}

func (t *TaskTreap) Max() *entities.Task{
	return t.max()
}

func (t *TaskTreap) SearchTask(task *entities.Task) *entities.Task{
	if node := t.search(task); node == nil{
		return nil
	}else{
		return node.task
	}
}

func (t *TaskTreap) PopTask(task *entities.Task){
	t.delete(task)
}


func (t *TaskTreap) AddTask(task *entities.Task){
	t.insert(task)
}

func (t *TaskTreap) PopNextTask() *entities.Task{
	task := t.min()
	t.delete(task)
	return task
}

func (t *TaskTreap) SeekNextTaskTime() time.Time{
	return t.min().Exp_time
}

func (t *TaskTreap) PopLastTask() *entities.Task{
	task := t.max()
	t.delete(task)
	return task
}