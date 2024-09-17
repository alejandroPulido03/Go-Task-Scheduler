package main

import (
	"task-scheduler/core"
	"task-scheduler/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	core.RegisterRoutes(e, routes.Routes)

	e.Start(":8080")
  
}
