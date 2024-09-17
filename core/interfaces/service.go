package interfaces

import "github.com/labstack/echo"

type Service struct {
	
	Method string
	Path string
	Middleware []echo.MiddlewareFunc
	Handler echo.HandlerFunc
}