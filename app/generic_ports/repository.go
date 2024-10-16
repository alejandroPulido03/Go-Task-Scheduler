package generic_ports

import "task-scheduler/app/entities"


type TaskRepository interface {
	Save(*entities.Task) error
	SaveRecovery(*entities.Task) error
	GetByTime(time string) (entities.Task, error)
	GetByInterval(start string, end string) (entities.Task, error)
}