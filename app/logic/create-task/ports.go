package create_task

import "task-scheduler/app/entities"

type CreateTaskService interface {
	CreateTask(entities.Task) (entities.Task, error)
}

type TaskRepository interface {
	Save(entities.Task) error
}