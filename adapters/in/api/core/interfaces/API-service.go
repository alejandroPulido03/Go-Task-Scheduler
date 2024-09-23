package interfaces

import (
	"github.com/labstack/echo"
)

type APIService struct {	
	Method string
	Path string
	Middleware []echo.MiddlewareFunc
	Handler echo.HandlerFunc
	Service interface{}
}