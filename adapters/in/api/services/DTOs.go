package services

import "task-scheduler/app/entities"

type taskDTO struct {
	Url string `json:"url"`
	Method string `json:"method"`
	Payload map[string]string `json:"payload"`
	Headers map[string]string `json:"headers"`
	Exp_time string `json:"exp_time"`
	Client_id string `json:"client_id"`
	Web_hook string `json:"web_hook"`
}

func (t *taskDTO) Parse() (entities.Task, error){
	return entities.Task{
		Url: t.Url,
		Method: t.Method,
		Payload: t.Payload,
		Headers: t.Headers,
		Exp_time: t.Exp_time,
		Client_id: t.Client_id,
		Web_hook: t.Web_hook,
	}, nil
}

