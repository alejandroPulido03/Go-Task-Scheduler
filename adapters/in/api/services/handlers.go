package services

import (
	"net/http"
	"task-scheduler/adapters/in/api/core"
	create_task "task-scheduler/app/logic/create-task"

	"github.com/labstack/echo"
)

type TaskService struct{
	s create_task.CreateTaskService
}

func (ts TaskService) createTaskHandler(c echo.Context) error {
	
	t := new(taskDTO)

	if err := c.Bind(t); err != nil {
		return err
	}
	parsed_task, err := t.Parse();	
	if err != nil {
		return c.JSON(http.StatusBadRequest, core.ErrorMessage(err))
	}

	created_task, err := ts.s.CreateTask(parsed_task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, core.ErrorMessage(err))
	}

	return c.JSON(http.StatusCreated, created_task)
		
	
}