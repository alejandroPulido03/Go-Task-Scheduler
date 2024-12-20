package create_task

import (
	"fmt"
	"task-scheduler/app/entities"
	"task-scheduler/app/logic/repository"
	"time"
)



type TaskHandler struct {
	task_repo repository.Repository
	
}

func (c *TaskHandler) CreateTask(task *entities.Task) (*entities.Task, error) {

	fmt.Println("Task created")

	if task.Exp_time.Before(time.Now().Add(1 * time.Minute)){
		return &entities.Task{}, fmt.Errorf("task expiration time must be at least 1 minute in the future")
	}

	err := c.task_repo.Save(task)
	if err != nil {
		return &entities.Task{}, err
	}

	return task, nil
}


func NewCreateTaskService(task_repo repository.Repository) *TaskHandler {
	return &TaskHandler{ task_repo, }
}