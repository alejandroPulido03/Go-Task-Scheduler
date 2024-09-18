package main

import (
	"task-scheduler/adapters/in/api/routes"
	"task-scheduler/core"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	core.RegisterRoutes(e, routes.Routes)

	e.Start(":8080")
  
}
