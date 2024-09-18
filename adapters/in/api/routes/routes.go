package routes

import (
	"task-scheduler/adapters/in/api/services"
	"task-scheduler/core/interfaces"
)


var Routes = []interfaces.Service{
	services.CreateTask,
}