package worker

import (
	"bytes"
	"fmt"
	"net/http"
	"task-scheduler/app/entities"
	"task-scheduler/app/repository"
	"time"
)



type Worker struct {
	task_repository repository.Repository
}

func NewWorker(task_repository repository.Repository) *Worker {
	return &Worker{ task_repository, }
}

func (w *Worker) Run() {
	go scheduler_tick(w)
	fmt.Println("Worker started")
}

func scheduler_tick(w *Worker) error{
	var tasks []*entities.Task
	var ch chan *entities.Task
	var err error
	for {
		if len(tasks) == 0{
			tasks, ch, err = get_next_tasks(w)
			if err != nil{
				break
			}
		}


		if len(tasks) != 0{
			fmt.Println("Next tasks time: ", tasks[0].Exp_time)
			tasks_time := tasks[0].Exp_time
			if tasks_time.Before(time.Now()){
				for i := 0; i < len(tasks); i++{
					ch <- tasks[i]
					fmt.Println("Task sent")
				}
				tasks = nil
			}
		}

		time.Sleep(1 * time.Second)
		
	}
	fmt.Println("Unexpected end", err)
	return err
}

func get_next_tasks(w *Worker) ([]*entities.Task, chan *entities.Task,  error){
	next_tasks := make([]*entities.Task, 0)
	ch := make(chan *entities.Task)
	t, err := w.task_repository.GetFirst()
	if err != nil{
		return nil, nil, err
	}

	if t != nil{
		next_tasks = append(next_tasks, t)
	}
	for{
		next_t, err := w.task_repository.GetFirst()
		if err != nil{
			return nil, nil, err
		}

		if next_t != nil && t.Exp_time.Equal(next_t.Exp_time){
			next_tasks = append(next_tasks, next_t)
		}else{
			break
		}
	}

	for i := 0; i < len(next_tasks); i++{
		go make_request(ch, w)
		
	}

	return next_tasks, ch, nil
}

func make_request(ch chan *entities.Task, worker *Worker) error {
	// Request to the server, we could implement pre load of the request for large payloads
	client := &http.Client{}

	task := <- ch
	fmt.Println("Task received")
	// Do something with the task
	body := bytes.NewReader(task.Payload)
	req, err := http.NewRequest(task.Method, task.Url, body)
	if err != nil{
		return err
	}

	for k, v := range task.Headers{
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil{
		return err
	}

	fmt.Println(task.Exp_time)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Body:", resp.Body)
	worker.task_repository.DeleteTask(task)
	

	return nil
}