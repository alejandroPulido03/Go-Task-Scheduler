package create_task

import (
	"fmt"
	"os"
	"strconv"
	"task-scheduler/app/entities"
	"task-scheduler/app/generic_ports"
	task_storage "task-scheduler/app/logic/worker/tasks_storage"
)



type CreateTaskHandler struct {
	repository generic_ports.TaskRepository
	task_store task_storage.TaskTreap
	MAX_IN_MEMORY_TASKS int
}

func (c *CreateTaskHandler) CreateTask(task *entities.Task) (*entities.Task, error) {

	fmt.Println("Task created")

	var err error

	if c.task_store.Size() < c.MAX_IN_MEMORY_TASKS {
		err = c.task_store.AddTask(task)
		c.repository.SaveRecovery(task)
	} else if last, max_err := c.task_store.Max(); max_err == nil && task.Exp_time.Before(last.Exp_time) {
		var last_task *entities.Task
		last_task, err = c.task_store.ReplaceLastTask(task)
		if err == nil {
			err = c.repository.Save(last_task)
		}
	} else {
		err = c.repository.Save(task)
	}

	if err == nil {
		
		return task, nil
	}
	return &entities.Task{}, err
}

func NewCreateTaskService(repository generic_ports.TaskRepository) *CreateTaskHandler {
	max_in_memory_tasks, err := strconv.Atoi(os.Getenv("MAX_IN_MEMORY_TASKS"))
	if err != nil {
		max_in_memory_tasks = 1000
	}

	return &CreateTaskHandler{
		repository: repository,
		task_store: task_storage.NewTaskTreapStorage(),
		MAX_IN_MEMORY_TASKS: max_in_memory_tasks,
	}
}