package app

import (
	"task-scheduler/adapters/in/api/services"
	"task-scheduler/adapters/out/db"
	create_task "task-scheduler/app/logic/create_task"
	task_storage "task-scheduler/app/logic/tasks_storage"
	"task-scheduler/app/logic/worker"
)

func SetupCreateTaskService() {
	db := db.NewRedisRepository()
	storage := task_storage.NewTaskTreapStorage()

	service := create_task.NewCreateTaskService(db, storage)
	worker_tasks := worker.NewWorker(storage)
	
	services.CreateTaskAPIService(service)
	worker_tasks.Run()

}