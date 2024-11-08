package services

import (
	"net/http"
	dto "task-scheduler/adapters/DTOs"
	"task-scheduler/adapters/in/api/core"
	create_task "task-scheduler/app/logic/create_task"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type TaskService struct{
	s create_task.CreateTaskService
}

func (ts TaskService) createTaskHandler(c echo.Context) error {
	
	t := dto.TaskDTO{
		JSON: &dto.TaskJSON{},
	}

	new_uuid, err := uuid.NewRandom()
	if err != nil {
		return c.JSON(http.StatusBadRequest, core.ErrorMessage(err))
	}

	t.JSON.Uuid = new_uuid.String()

	if err := c.Bind(t.JSON); err != nil {
		return c.JSON(http.StatusBadRequest, core.ErrorMessage(err))
	}
	task, err := t.ToEntity();
	if err != nil {
		return c.JSON(http.StatusBadRequest, core.ErrorMessage(err))
	}

	created_task, err := ts.s.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, core.ErrorMessage(err))
	}

	return c.JSON(http.StatusCreated, created_task)
		
	
}