package mem_storage

import (
	"fmt"
	"task-scheduler/app/entities"
	"testing"
	"time"
)

func TestTaskTreapAddTask(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task := &entities.Task{Exp_time: baseTime}
	err := treap.AddTask(task)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	if task, err := treap.SearchTask(task); err != nil || task == nil {
		t.Error("Task was not added successfully.")
	}
}

func TestTaskTreapPopNextTask(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task1 := &entities.Task{Exp_time: baseTime}
	task2 := &entities.Task{Exp_time: baseTime.Add(time.Minute)}

	err := treap.AddTask(task1)
	if err != nil {
		t.Fatalf("Failed to add task1: %v", err)
	}
	err = treap.AddTask(task2)
	if err != nil {
		t.Fatalf("Failed to add task2: %v", err)
	}

	poppedTask, err := treap.PopNextTask()
	if err != nil {
		t.Fatalf("Failed to pop next task: %v", err)
	}

	if poppedTask != task1 {
		t.Errorf("Expected %v, but got %v", task1, poppedTask)
	}

	if treap.size != 1 {
		t.Errorf("Expected size 1, but got %d", treap.size)
	}

	sec_task, err := treap.PopNextTask()
	if err != nil {
		t.Fatalf("Failed to pop next task: %v", err)
	}

	if sec_task != task2 {
		fmt.Println(sec_task == task1)
		t.Errorf("Expected %v, but got %v", task2, sec_task)
	}

	if treap.size != 0 {
		t.Errorf("Expected size 0, but got %d", treap.size)
	}

	no_task, err := treap.PopNextTask()
	if err != nil {
		t.Fatalf("Failed to pop next task: %v", err)
	}

	if no_task != nil {
		t.Errorf("Expected nil, but got %v", no_task)
	}

	if treap.size != 0 {
		t.Errorf("Expected size 0, but got %d", treap.size)
	}


}

func TestTaskTreapPopLastTask(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task1 := &entities.Task{Exp_time: baseTime}
	task2 := &entities.Task{Exp_time: baseTime.Add(time.Minute)}

	err := treap.AddTask(task1)
	if err != nil {
		t.Fatalf("Failed to add task1: %v", err)
	}
	err = treap.AddTask(task2)
	if err != nil {
		t.Fatalf("Failed to add task2: %v", err)
	}

	poppedTask, err := treap.PopLastTask()
	if err != nil {
		t.Fatalf("Failed to pop last task: %v", err)
	}

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

	err := treap.AddTask(task1)
	if err != nil {
		t.Fatalf("Failed to add task1: %v", err)
	}
	err = treap.AddTask(task2)
	if err != nil {
		t.Fatalf("Failed to add task2: %v", err)
	}
	err = treap.AddTask(task3)
	if err != nil {
		t.Fatalf("Failed to add task3: %v", err)
	}

	// Verificamos el mínimo
	minTask, err := treap.Min()
	if err != nil {
		t.Fatalf("Failed to get min task: %v", err)
	}
	if minTask != task1 {
		t.Errorf("Expected min task %v, but got %v", task1, minTask)
	}

	// Verificamos el máximo
	maxTask, err := treap.Max()
	if err != nil {
		t.Fatalf("Failed to get max task: %v", err)
	}
	if maxTask != task3 {
		t.Errorf("Expected max task %v, but got %v", task3, maxTask)
	}
}

func TestTaskTreapReplaceLastTask(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := NewTaskTreapStorage()

	task1 := &entities.Task{Exp_time: baseTime}
	task2 := &entities.Task{Exp_time: baseTime.Add(time.Minute)}
	newTask := &entities.Task{Exp_time: baseTime.Add(2 * time.Minute)}

	err := treap.AddTask(task1)
	if err != nil {
		t.Fatalf("Failed to add task1: %v", err)
	}
	err = treap.AddTask(task2)
	if err != nil {
		t.Fatalf("Failed to add task2: %v", err)
	}

	replacedTask, err := treap.ReplaceLastTask(newTask)
	if err != nil {
		t.Fatalf("Failed to replace last task: %v", err)
	}

	if replacedTask != task2 {
		t.Errorf("Expected replaced task %v, but got %v", task2, replacedTask)
	}

	// Verify the new task is now the last task
	maxTask, err := treap.Max()
	if err != nil {
		t.Fatalf("Failed to get max task: %v", err)
	}
	if maxTask != newTask {
		t.Errorf("Expected max task %v, but got %v", newTask, maxTask)
	}
}


