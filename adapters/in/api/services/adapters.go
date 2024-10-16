package services

import (
	"task-scheduler/adapters/in/api/core"
	"task-scheduler/adapters/in/api/core/interfaces"
	create_task "task-scheduler/app/logic/create_task"
	"task-scheduler/server"
)

func CreateTaskAPIService(service create_task.CreateTaskService) {
	t := TaskService{
		s: service,
	}

	APIService := interfaces.APIService{
		Method: "POST",
		Path: "/task",
		Handler: t.createTaskHandler,
	}

	core.RegisterRoute(server.GetCtx(), APIService)
}
