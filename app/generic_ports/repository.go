package generic_ports

import "task-scheduler/app/entities"


type SecondaryStorage interface {
	Save(*entities.Task) error
	SaveRecovery(*entities.Task) error
	GetByTime(time string) (entities.Task, error)
	GetByInterval(start string, end string) (entities.Task, error)
	Remove(*entities.Task) error
}

type MemoryStorage interface {
	AddTask(task *entities.Task) error
	PopNextTask() (*entities.Task, error)
	PopLastTask() (*entities.Task, error)
	PopTask(task *entities.Task) error
	SearchTask(task *entities.Task) (*entities.Task, error)
	ReplaceLastTask(task *entities.Task) (*entities.Task, error)
	Size() (int, error)
	Min() (*entities.Task, error)
	Max() (*entities.Task, error)
}

