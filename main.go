package main

import (
	setup "task-scheduler/app"
	"task-scheduler/server"
)

func main() {

	server.NewServer()

	setup.SetupCreateTaskService()

	server.StartServer()

	
  
}
