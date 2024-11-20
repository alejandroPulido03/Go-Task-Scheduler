package create_task

import (
	"fmt"
	"task-scheduler/app/entities"
	"task-scheduler/app/repository"
)



type TaskHandler struct {
	task_repo repository.Repository
	
}

func (c *TaskHandler) CreateTask(task *entities.Task) (*entities.Task, error) {

	fmt.Println("Task created")

	err := c.task_repo.Save(task)
	if err != nil {
		return &entities.Task{}, err
	}

	return task, nil
}


func NewCreateTaskService(task_repo repository.Repository) *TaskHandler {
	return &TaskHandler{ task_repo, }
}