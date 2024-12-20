package repository

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"task-scheduler/app/entities"
)

type Repository interface {
	Save(*entities.Task) error
	GetFirst() (*entities.Task, error)
	PushFirst() (*entities.Task, error)
	DeleteTask(*entities.Task) error
}

type TaskRepository struct {
	secondary SecondaryStorage
	main MemoryStorage
	MAX_IN_MEMORY_TASKS int
	mu sync.Mutex
}

func NewTaskRepository(secondary SecondaryStorage, main MemoryStorage) *TaskRepository {
	max_in_memory_tasks, err := strconv.Atoi(os.Getenv("MAX_IN_MEMORY_TASKS"))
	if err != nil {
		max_in_memory_tasks = 3
	}
	return &TaskRepository{
		secondary: secondary,
		main: main,
		MAX_IN_MEMORY_TASKS: max_in_memory_tasks,
	}
}

func (t *TaskRepository) Save(task *entities.Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.save_op(task)
}

func (t *TaskRepository) save_op(task *entities.Task) error {
	size, size_err := t.main.Size()
	if size_err != nil {
		return size_err
	}

	if size + 1 <= t.MAX_IN_MEMORY_TASKS {
		add_err := t.main.AddTask(task)
		if add_err != nil{
			return add_err
		}

		save_r_err := t.secondary.SaveRecovery(task)
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
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.main.GetFirst()
}

func (t *TaskRepository) PushFirst() (*entities.Task, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.main.PopNextTask()
}

func (t *TaskRepository) DeleteTask(task *entities.Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	err := t.secondary.RemoveRecovery(task)
	

	main_size, size_err := t.main.Size()
	if size_err != nil {
		return size_err
	}

	if main_size + 1 <= t.MAX_IN_MEMORY_TASKS {
		next_t, sec_err := t.secondary.GetFirst()	
		if sec_err != nil {
			return sec_err
		}
		if next_t != nil {
			t.save_op(next_t)
		}
	}

	return err
}


func (t *TaskRepository) switchTaskStorage(task *entities.Task) error{
	last, max_err := t.main.GetMax()
	if max_err != nil {
		return max_err
	}

	if task.Exp_time.Before(last.Exp_time) {
		
		last_task, rep_err := t.main.ReplaceLastTask(task)
		if rep_err != nil {
			return rep_err
		}

		s_err := t.secondary.Save(last_task)
		fmt.Println("Task switched")
		if s_err != nil {
			return s_err
		}

	} else {
		s_err := t.secondary.Save(task)
		fmt.Println("Task saved")
		if s_err != nil {
			return s_err
		}
	}

	return nil
}