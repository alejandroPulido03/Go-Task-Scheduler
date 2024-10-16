package task_storage

import (
	"task-scheduler/app/entities"
	"testing"
	"time"
)

func TestTaskTreapAddTask(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task := &entities.Task{Exp_time: baseTime}
	treap.AddTask(task)

	if treap.SearchTask(task) == nil {
		t.Error("Task was not added successfully.")
	}
}

func TestTaskTreapPopNextTask(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task1 := &entities.Task{Exp_time: baseTime}
	task2 := &entities.Task{Exp_time: baseTime.Add(time.Minute)}

	treap.AddTask(task1)
	treap.AddTask(task2)

	poppedTask := treap.PopNextTask()

	if poppedTask != task1 {
		t.Errorf("Expected %v, but got %v", task1, poppedTask)
	}
}

func TestTaskTreapSeekNextTaskTime(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task1 := &entities.Task{Exp_time: baseTime}
	task2 := &entities.Task{Exp_time: baseTime.Add(time.Minute)}

	treap.AddTask(task1)
	treap.AddTask(task2)

	nextTaskTime := treap.SeekNextTaskTime()

	if nextTaskTime != baseTime {
		t.Errorf("Expected next task time %v, but got %v", baseTime, nextTaskTime)
	}
}

func TestTaskTreapPopLastTask(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task1 := &entities.Task{Exp_time: baseTime}
	task2 := &entities.Task{Exp_time: baseTime.Add(time.Minute)}

	treap.AddTask(task1)
	treap.AddTask(task2)

	poppedTask := treap.PopLastTask()

	if poppedTask != task2 {
		t.Errorf("Expected %v, but got %v", task2, poppedTask)
	}
}

func TestTaskTreapMinMax(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task1 := &entities.Task{Exp_time: baseTime}
	task2 := &entities.Task{Exp_time: baseTime.Add(time.Minute)}
	task3 := &entities.Task{Exp_time: baseTime.Add(2 * time.Minute)}

	treap.AddTask(task1)
	treap.AddTask(task2)
	treap.AddTask(task3)

	// Verificamos el mínimo
	minTask := treap.Min()
	if minTask != task1 {
		t.Errorf("Expected min task %v, but got %v", task1, minTask)
	}

	// Verificamos el máximo
	maxTask := treap.Max()
	if maxTask != task3 {
		t.Errorf("Expected max task %v, but got %v", task3, maxTask)
	}
}
