package create_task

import (
	"fmt"
	"task-scheduler/app/entities"
)

type CreateTaskHandler struct {

}

func (c *CreateTaskHandler) CreateTask(task entities.Task) (entities.Task, error) {
	fmt.Println("Task created")
	
	return task, nil
}

func NewCreateTaskHandler() *CreateTaskHandler {
	return &CreateTaskHandler{}
}