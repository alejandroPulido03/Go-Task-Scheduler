package worker

import (
	"fmt"
	"net/http"
	"sort"
	"task-scheduler/adapters/in-out/mem_storage"
	"task-scheduler/app/entities"
	"task-scheduler/app/logic/repository"
	"testing"
	"time"
)

type mockRepository struct {
	tasks []*entities.Task
}

type mockSecondaryStorage struct {
	tasks []*entities.Task
}

func (m *mockSecondaryStorage) Save(task *entities.Task) error {
	return nil
}

func (m *mockSecondaryStorage) SaveRecovery(task *entities.Task) error {
	return nil
}

func (m *mockSecondaryStorage) GetFirst() (*entities.Task, error) {
	return nil, nil
}

func (m *mockSecondaryStorage) RemoveRecovery(task *entities.Task) error {
	return nil
}


func (m *mockRepository) Save(task *entities.Task) error {
	m.tasks = append(m.tasks, task)
	sort.Slice(m.tasks, func(i, j int) bool {
		return m.tasks[i].Exp_time.Before(m.tasks[j].Exp_time)
	})

	return nil
}

func (m *mockRepository) PushFirst() (*entities.Task, error) {
	if len(m.tasks) == 0 {
		return nil, nil
	}

	task := m.tasks[0]
	m.tasks = m.tasks[1:]
	return task, nil
}

func (m *mockRepository) GetFirst() (*entities.Task, error) {
	if len(m.tasks) == 0 {
		return nil, nil
	}

	return m.tasks[0], nil
}

func (m *mockRepository) DeleteTask(task *entities.Task) error {
	for i, t := range m.tasks {
		if t == task {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}


func newMockWorker(t []*entities.Task) *Worker {
	return NewWorker(&mockRepository{t,})
}

func mockTask(t time.Time) *entities.Task {
	return &entities.Task{
		Url: "http://example.com",
		Method: "GET",
		Body: []byte(""),
		Headers: map[string][]string{},
		Exp_time: t,
		Client_id: "1",
		Web_hook: "http://example.com",
		Uuid: "1",
	}
}

func newServerMock(ch chan map[string]string) {

	response := make(map[string]string)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		response["status"] = "ok"
	
		for k, v := range r.Header {
			response[k] = v[0]
		}

		ch <- response
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
	
}

func TestGetNextTasksEmpty(t *testing.T) {
	w := newMockWorker(make([]*entities.Task, 0))
	
	tasks, ch, err := getNextTasks(w)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}

	if ch == nil {
		t.Errorf("Expected channel, got nil")
	}
	
}

func TestGetNextTasksOne(t *testing.T) {
	t1 := mockTask(time.Now())
	w := newMockWorker([]*entities.Task{t1})
	
	tasks, ch, err := getNextTasks(w)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 tasks, got %d", len(tasks))
	}

	if ch == nil {
		t.Errorf("Expected channel, got nil")
	}
	
}

func TestGetNextTasksSeveralSameTime(t *testing.T) {
	t1 := mockTask(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	t2 := mockTask(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	t3 := mockTask(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	w := newMockWorker([]*entities.Task{t1, t2, t3})
	
	tasks, ch, err := getNextTasks(w)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
	}

	if ch == nil {
		t.Errorf("Expected channel, got nil")
	}
	
}


func TestGetNextTasksSeveralDifferentTime(t *testing.T) {
	t1 := mockTask(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	t2 := mockTask(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Second))
	t3 := mockTask(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Second * 2))
	w := newMockWorker([]*entities.Task{t1, t2, t3})
	
	tasks, ch, err := getNextTasks(w)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 tasks, got %d", len(tasks))
	}

	if ch == nil {
		t.Errorf("Expected channel, got nil")
	}
	
}



func TestMakeRequest(t *testing.T) {
	w := newMockWorker(make([]*entities.Task, 0))
	ch := make(chan *entities.Task)
	task := entities.Task{
		Url: "http://localhost:8000",
		Method: "GET",
		Body: []byte(""),
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
			"Custom-Header": {"value"},
		},
		Exp_time: time.Now(),
		Client_id: "1",
	}

	res_chan := make(chan map[string]string)
	go newServerMock(res_chan)
	fmt.Println("Server started")

	go makeRequest(ch, w)
	ch <- &task

	response := <- res_chan
	if response["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", response["Content-Type"])
	}

	if response["Custom-Header"] != "value" {
		t.Errorf("Expected Custom-Header value, got %s", response["Custom-Header"])
	}

	if response["status"] != "ok" {
		t.Errorf("Expected status ok, got %s", response["status"])
	}
	
}

func TestMakeRequestUpdateFirstRequest(t *testing.T) {
	w := NewWorker(repository.NewTaskRepository(
		&mockSecondaryStorage{},
		mem_storage.NewTaskTreapStorage()))

	task_1 := entities.Task{
		Url: "http://localhost:8000",
		Method: "GET",
		Body: []byte(""),
		Headers: map[string][]string{
			"Order": {"2"},
		},
		Exp_time: time.Now().Add(time.Minute * 20),
		Client_id: "1",
	}

	task_2 := entities.Task{
		Url: "http://localhost:8000",
		Method: "GET",
		Body: []byte(""),
		Headers: map[string][]string{
			"Order": {"1"},
		},
		Exp_time: time.Now().Add(time.Minute * 10),
		Client_id: "1",
	}

	res_chan := make(chan map[string]string)
	go newServerMock(res_chan)
	fmt.Println("Server started")
	go schedulerTick(w)
	fmt.Println("Scheduler started")

	w.task_repository.Save(&task_1)
	w.task_repository.Save(&task_2)
	fmt.Println("Tasks saved")

	response := <- res_chan
	if response["Order"] != "1" {
		t.Errorf("Expected Order 1, got %s", response["Order"])
	}

	response = <- res_chan
	if response["Order"] != "2" {
		t.Errorf("Expected Order 2, got %s", response["Order"])
	}

	
}

