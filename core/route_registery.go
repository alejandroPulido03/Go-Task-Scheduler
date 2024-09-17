package core

import (
	"errors"
	"task-scheduler/core/interfaces"

	"github.com/labstack/echo"
)



func RegisterRoutes(e *echo.Echo, routes []interfaces.Service) error{
	for _, route := range routes {
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

	return nil
}
