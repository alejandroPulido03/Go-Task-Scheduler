package entities

type Task struct {
	Url string
	Method string
	Payload map[string]string
	Headers map[string]string
	Exp_time string
	Client_id string
	Web_hook string
}