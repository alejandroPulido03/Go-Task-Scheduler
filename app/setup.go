package app

import (
	"task-scheduler/adapters/in-out/db"
	"task-scheduler/adapters/in-out/mem_storage"
	"task-scheduler/adapters/in/api/services"
	"task-scheduler/app/logic/create_task"
	"task-scheduler/app/logic/repository"
	"task-scheduler/app/logic/worker"
)

func SetupCreateTaskService() {

	//External adapters
	db := db.NewRedisRepository()
	storage := mem_storage.NewTaskTreapStorage()

	//Core services
	repo := repository.NewTaskRepository(db, storage)
	service := create_task.NewCreateTaskService(repo)
	worker_tasks := worker.NewWorker(repo)
	
	services.CreateTaskAPIService(service)
	worker_tasks.Run()

}