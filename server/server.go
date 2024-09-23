package server

import "github.com/labstack/echo"

type Server struct {
	Engine *echo.Echo
}

var server *Server

func NewServer() {
	server = &Server{
		Engine: echo.New(),
	}
}

func StartServer() {
	server.Engine.Start(":8080")
}

func GetCtx() *echo.Echo {
	
	return server.Engine
}

