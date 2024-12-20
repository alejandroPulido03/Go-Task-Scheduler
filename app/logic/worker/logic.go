package worker

import (
	"bytes"
	"fmt"
	"net/http"
	"task-scheduler/app/entities"
	"task-scheduler/app/logic/repository"
	"time"
)



type Worker struct {
	task_repository repository.Repository
}

func NewWorker(task_repository repository.Repository) *Worker {
	return &Worker{ task_repository, }
}

func (w *Worker) Run() {
	go schedulerTick(w)
	fmt.Println("Worker started")
}

func schedulerTick(w *Worker) error{
	var tasks []*entities.Task
	var ch chan *entities.Task
	var err error
	for {

		if first, search_err := w.task_repository.GetFirst(); search_err == nil && first != nil && first.Exp_time.Before(time.Now().Add(1 * time.Minute)){
			// If there is no tasks in the queue, we will wait for the next task to be executed in the next minute (that need the ban of request of scheduling by less than a minute than the current time)

			tasks, ch, err = getNextTasks(w)

			if err != nil{
				break
			}
		}else if search_err != nil{
			return search_err
		}


		if len(tasks) != 0{
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

	return err
}

func getNextTasks(w *Worker) ([]*entities.Task, chan *entities.Task,  error){
	next_tasks := make([]*entities.Task, 0)
	ch := make(chan *entities.Task)
	t, err := w.task_repository.PushFirst()
	if err != nil{
		return nil, nil, err
	}

	if t != nil{
		next_tasks = append(next_tasks, t)
	}
	for{
		next_t, err := w.task_repository.PushFirst()
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
		go makeRequest(ch, w)
		
	}

	return next_tasks, ch, nil
}

func makeRequest(ch chan *entities.Task, worker *Worker) error {
	// Request to the server, we could implement pre load of the request for large payloads
	client := &http.Client{}

	task := <- ch
	fmt.Println("Task received")
	// Do something with the task
	body := bytes.NewReader(task.Body)
	req, err := http.NewRequest(task.Method, task.Url, body)
	if err != nil{
		return err
	}

	for k, v := range task.Headers{
		for _, h := range v{
			req.Header.Add(k, h)
		}
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