package dtos

import (
	"encoding/json"
	"task-scheduler/app/entities"
	"time"
)

type TaskJSON struct {
	Url string `json:"url"`
	Method string `json:"method"`
	Payload map[string]string `json:"payload"`
	Headers map[string]string `json:"headers"`
	Exp_time string `json:"exp_time"`
	Client_id string `json:"client_id"`
	Web_hook string `json:"web_hook"`
}

type TaskDTO struct {
	JSON TaskJSON
	Entity entities.Task
}

func (t *TaskDTO) ToEntity() (entities.Task, error){
	exp_time, err := time.Parse(time.UnixDate, t.JSON.Exp_time)
	
	if err != nil{
		return entities.Task{}, err
	}

	return entities.Task{
		Url: t.JSON.Url,
		Method: t.JSON.Method,
		Payload: t.JSON.Payload,
		Headers: t.JSON.Headers,
		Exp_time: exp_time,
		Client_id: t.JSON.Client_id,
		Web_hook: t.JSON.Web_hook,
	}, nil
}

func (t *TaskDTO) ToJson() ([]byte, error){
	t.JSON = TaskJSON{
		Url: t.Entity.Url,
		Method: t.Entity.Method,
		Payload: t.Entity.Payload,
		Headers: t.Entity.Headers,
		Exp_time: t.Entity.Exp_time.String(),
		Client_id: t.Entity.Client_id,
		Web_hook: t.Entity.Web_hook,
	}
	return json.Marshal(t.JSON)
}