package main

import (
	setup "task-scheduler/app/services_setup"
	"task-scheduler/server"
)

func main() {

	server.NewServer()

	setup.SetupCreateTaskService()

	server.StartServer()

	
  
}
