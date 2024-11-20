package repository

import (
	"os"
	"strconv"
	"task-scheduler/app/entities"
	"task-scheduler/app/generic_ports"
)

type Repository interface {
	Save(*entities.Task) error
	GetFirst() (*entities.Task, error)
	DeleteTask(*entities.Task) error
}

type TaskRepository struct {
	recovery generic_ports.SecondaryStorage
	main generic_ports.MemoryStorage
	MAX_IN_MEMORY_TASKS int
}

func NewTaskRepository(recovery generic_ports.SecondaryStorage, main generic_ports.MemoryStorage) *TaskRepository {
	max_in_memory_tasks, err := strconv.Atoi(os.Getenv("MAX_IN_MEMORY_TASKS"))
	if err != nil {
		max_in_memory_tasks = 1000
	}
	return &TaskRepository{
		recovery, main, max_in_memory_tasks,
	}
}

func (t *TaskRepository) Save(task *entities.Task) error {
	size, size_err := t.main.Size()
	if size_err != nil {
		return size_err
	}


	if size < t.MAX_IN_MEMORY_TASKS {
		add_err := t.main.AddTask(task)
		if add_err != nil{
			return add_err
		}

		save_r_err := t.recovery.SaveRecovery(task)
		if save_r_err != nil {
			return save_r_err
		}
		
	} else {
		switch_err := t.switchTaskStorage(task)
		if switch_err != nil {
			return switch_err
		}
	}

	return nil
}

func (t *TaskRepository) GetFirst() (*entities.Task, error) {
	return t.main.PopNextTask()
}

func (t *TaskRepository) DeleteTask(task *entities.Task) error {
	return t.recovery.Remove(task)
}


func (t *TaskRepository) switchTaskStorage(task *entities.Task) error{
	last, max_err := t.main.Max()
	if max_err != nil {
		return max_err
	}

	if task.Exp_time.Before(last.Exp_time) {
		
		last_task, rep_err := t.main.ReplaceLastTask(task)
		if rep_err != nil {
			return rep_err
		}

		s_err := t.recovery.Save(last_task)
		if s_err != nil {
			return s_err
		}

	} else {
		s_err := t.recovery.Save(task)
		if s_err != nil {
			return s_err
		}
	}

	return nil
}