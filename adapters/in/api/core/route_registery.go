package core

import (
	"errors"
	"task-scheduler/adapters/in/api/core/interfaces"

	"github.com/labstack/echo"
)



func RegisterRoute(e *echo.Echo, route interfaces.APIService) error{
	
	if route.Method != "" && route.Path != "" {
		if route.Middleware != nil {
			e.Add(route.Method, route.Path, route.Handler, route.Middleware...)
		} else {
			e.Add(route.Method, route.Path, route.Handler)
		}
		return nil
	}else {
		return errors.New("invalid route")
	}

}
