package create_task

import (
	"fmt"
	"task-scheduler/app/entities"
	"task-scheduler/app/generic_ports"
)

type CreateTaskHandler struct {
	repository generic_ports.TaskRepository
}

func (c *CreateTaskHandler) CreateTask(task entities.Task) (entities.Task, error) {

	fmt.Println("Task created")
	err := c.repository.Save(task)

	if err == nil {
			
		return task, nil
	}
	return entities.Task{}, err
}

func NewCreateTaskService(repository generic_ports.TaskRepository) *CreateTaskHandler {
	return &CreateTaskHandler{
		repository: repository,
	}
}