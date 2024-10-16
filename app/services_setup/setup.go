package setup

import (
	"task-scheduler/adapters/in/api/services"
	"task-scheduler/adapters/out/db"
	create_task "task-scheduler/app/logic/create_task"
)

func SetupCreateTaskService() {
	db := db.NewRedisRepository()
	service := create_task.NewCreateTaskService(db)
	services.CreateTaskAPIService(service)
}