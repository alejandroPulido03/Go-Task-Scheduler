package task_storage

import (
	"fmt"
	"task-scheduler/app/entities"
	"testing"
	"time"
)

const NUM_NODES = 1000000

func verifyHeapProperty(t *taskNode, test *testing.T) bool {
	if t == nil {
		return true
	}

	// Verificamos si el hijo izquierdo cumple con la propiedad del heap
	if t.left != nil && t.left.priority > t.priority {
		test.Errorf("Heap property violated at node: %v. Left child has higher priority.", t.task)
		return false
	}

	// Verificamos si el hijo derecho cumple con la propiedad del heap
	if t.right != nil && t.right.priority > t.priority {
		test.Errorf("Heap property violated at node: %v. Right child has higher priority.", t.task)
		return false
	}

	// Recursión para el hijo izquierdo y derecho
	return verifyHeapProperty(t.left, test) && verifyHeapProperty(t.right, test)
}

func verifyTreeProperty(t *taskNode, test *testing.T) bool {
	if t == nil {
		return true
	}

	// Verificamos si el hijo izquierdo cumple con la propiedad del árbol
	if t.left != nil && compare(t.left.task, t.task) == 1 {
		test.Errorf("Tree property violated at node: %v. Left child has higher expiration time.", t.task)
		return false
	}

	// Verificamos si el hijo derecho cumple con la propiedad del árbol
	if t.right != nil && compare(t.right.task, t.task) == -1 {
		test.Errorf("Tree property violated at node: %v. Right child has lower expiration time.", t.task)
		return false
	}

	// Recursión para el hijo izquierdo y derecho
	return verifyTreeProperty(t.left, test) && verifyTreeProperty(t.right, test)
}

func verifyNodes(t *taskNode, test *testing.T, list_nodes *[]int) {
	if t != nil {
		*list_nodes = append(*list_nodes, t.task.Exp_time.Minute())
		verifyNodes(t.left, test, list_nodes)
		verifyNodes(t.right, test, list_nodes)
	}
}

func checkInvariants(treap *TaskTreap, t *testing.T, numNodes int) {
	minutes := make([]int, 0)
	verifyNodes(treap.root, t, &minutes)
	if len(minutes) != numNodes {
		t.Errorf("Number of nodes in treap is %d, expected %d.", len(minutes), numNodes)
	}
	
	if !verifyTreeProperty(treap.root, t) {
		t.Errorf("Tree property violated after inserting %d nodes.", numNodes)
	}

	if !verifyHeapProperty(treap.root, t) {
		t.Errorf("Heap property violated after inserting %d nodes.", numNodes)
	}
}

func getMaxDepth(t *taskNode) int {
	if t == nil {
		return 0
	}

	leftDepth := getMaxDepth(t.left)
	rightDepth := getMaxDepth(t.right)

	return 1 + max(leftDepth, rightDepth)
}


func insertNodes(treap *TaskTreap, numNodes int, baseTime time.Time) {
	for i := 0; i < numNodes; i++ {
		task := &entities.Task{
			Exp_time: baseTime.Add(time.Duration(i) * time.Minute),
		}
		treap.AddTask(task)
	}
}


func TestTaskTreapInsert(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := &TaskTreap{}

	insertNodes(treap, NUM_NODES, baseTime)

	fmt.Println("Max depth:", getMaxDepth(treap.root))
	checkInvariants(treap, t, NUM_NODES)

}




func TestTaskTreapDelete(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := &TaskTreap{}

	insertNodes(treap, NUM_NODES, baseTime)

	for i := 0; i < NUM_NODES / 2; i ++ {
		treap.PopTask(&entities.Task{Exp_time: baseTime.Add(time.Duration(i) * time.Minute)})
	}
	
	checkInvariants(treap, t, NUM_NODES / 2)
}

func TestTaskTreapSearch(t *testing.T) {
	baseTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	treap := &TaskTreap{}

	insertNodes(treap, NUM_NODES, baseTime)

	for i := 1; i < NUM_NODES; i++ {
		found := treap.search(&entities.Task{Exp_time: baseTime.Add(time.Duration(i) * time.Minute)})
		if found == nil{
			t.Errorf("Task with ID %d not found in treap.", i)
		}
	}

	// Search for a non-existing task
	notFound := treap.search(&entities.Task{})
	if notFound != nil {
		t.Error("Non-existing task found in treap.")
	}

	checkInvariants(treap, t, NUM_NODES)
}

