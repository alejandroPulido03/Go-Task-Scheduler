package entities

import "time"

type Task struct {
	Url string
	Method string
	Body []byte
	Headers map[string][]string
	Exp_time time.Time
	Client_id string
	Web_hook string
	Uuid string
}