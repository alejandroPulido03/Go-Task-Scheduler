package worker

import task_storage "task-scheduler/app/logic/tasks_storage"

type Worker struct {
	task_store task_storage.TaskTreap
}

func NewWorker(storage task_storage.TaskTreap) *Worker {
	return &Worker{
		task_store: storage,
	}
}

func (w *Worker) Run() {
	// Do some work
}