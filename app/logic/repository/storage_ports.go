package repository

import "task-scheduler/app/entities"


type SecondaryStorage interface {
	Save(*entities.Task) error
	SaveRecovery(*entities.Task) error
	GetFirst() (*entities.Task, error)
	RemoveRecovery(*entities.Task) error
}

type MemoryStorage interface {
	AddTask(task *entities.Task) error
	PopNextTask() (*entities.Task, error)
	PopLastTask() (*entities.Task, error)
	PopTask(task *entities.Task) error
	SearchTask(task *entities.Task) (*entities.Task, error)
	ReplaceLastTask(task *entities.Task) (*entities.Task, error)
	Size() (int, error)
	GetFirst() (*entities.Task, error)
	GetMax() (*entities.Task, error)
}

