package services

import (
	"net/http"
	"task-scheduler/app/entities"
	create_task "task-scheduler/app/logic/create-task"
	"task-scheduler/core/interfaces"

	"github.com/labstack/echo"
)

type taskDTO struct {
	Url string `json:"url"`
	Method string `json:"method"`
	Payload map[string]string `json:"payload"`
	Headers map[string]string `json:"headers"`
	Exp_time string `json:"exp_time"`
	Client_id string `json:"client_id"`
	Web_hook string `json:"web_hook"`
}

func createTaskHandler(c echo.Context) error {
	t := new(taskDTO)
	if err := c.Bind(t); err != nil {
		return err
	}
	task := parseTask(t)
	create_task.NewCreateTaskHandler().CreateTask(task)

	return c.JSON(http.StatusCreated, t)
}

func parseTask(t *taskDTO) entities.Task{
	return entities.Task{
		Url: t.Url,
		Method: t.Method,
		Payload: t.Payload,
		Headers: t.Headers,
		Exp_time: t.Exp_time,
		Client_id: t.Client_id,
		Web_hook: t.Web_hook,
	}
}

var CreateTask = interfaces.Service{
	Method: "POST",
	Path: "/task",
	Handler: createTaskHandler,
}