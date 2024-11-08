package entities

import "time"

type Task struct {
	Url string
	Method string
	Payload map[string]string
	Headers map[string]string
	Exp_time time.Time
	Client_id string
	Web_hook string
	Uuid string
}