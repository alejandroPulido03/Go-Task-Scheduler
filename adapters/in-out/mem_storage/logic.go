package mem_storage

import (
	"math/rand"
	"task-scheduler/app/entities"
)

type TaskTreap struct {
	root *taskNode
	minNode *taskNode
	maxNode *taskNode
	size int
}

type taskNode struct {
	task *entities.Task
	left *taskNode
	right *taskNode
	parent *taskNode
	priority float32
	id int
}

func (t *TaskTreap) min() *entities.Task{
	
	node := t.root
	for node != nil && node.left != nil{
		node = node.left
	}
	t.minNode = node
	
	if t.minNode == nil{
		return nil
	}

	return t.minNode.task
}

func (t *TaskTreap) max() *entities.Task{
	
	node := t.root
	for node != nil && node.right != nil{
		node = node.right
	}
	t.maxNode = node

	if t.maxNode == nil{
		return nil
	}

	return t.maxNode.task
}

func (t *TaskTreap) insert(task *entities.Task){
	new_node := &taskNode{
		task: task,
		priority: rand.Float32(),
		id: task.Exp_time.Minute(),
	}

	if t.root == nil{
		t.root = new_node
		t.minNode = new_node
		t.maxNode = new_node
	}else{
		if compare(t.minNode.task, task) == 1{
			t.minNode = new_node
		}
		if compare(t.maxNode.task, task) == -1{
			t.maxNode = new_node
		}
		t.dive(t.root, new_node)
	}
	t.size++
}

func (t *TaskTreap) search(task *entities.Task) *taskNode{
	node := t.root
	for node != nil && compare(node.task, task) != 0{
		if compare(node.task, task) == -1{
			node = node.right
		}else{
			node = node.left
		}
	}

	if node != nil{
		return node
	}else{
		return nil
	}
}

func (t *TaskTreap) delete(task *entities.Task){
	node := t.search(task)
	if node != nil{
		node.priority = -1
		for node.left != nil || node.right != nil{
			if node.left != nil && (node.right == nil || node.left.priority > node.right.priority){
				t.rotateRight(node.left)
			}else if node.right != nil && (node.left == nil || node.right.priority > node.left.priority){
				t.rotateLeft(node.right)
			}
		}
		
		if node.parent != nil {
			if node.parent.left == node{
				node.parent.left = nil
			}else if node.parent.right == node{
				node.parent.right = nil
			}
		}else {
			t.root = nil
		}
		t.size--
	}

	t.min()
	t.max()
}


func (t *TaskTreap) dive(node *taskNode, new_node *taskNode){
	if compare(node.task, new_node.task) == 1{
		if node.left == nil{
			node.left = new_node
			new_node.parent = node
		}else{
			t.dive(node.left, new_node)
		}
	}else{
		if node.right == nil{
			node.right = new_node
			new_node.parent = node

		}else{
			t.dive(node.right, new_node)
		}
	}	

	t.checkRotate(new_node)
}

func (t *TaskTreap) checkRotate(node *taskNode){

	if node.parent != nil && node.parent.priority < node.priority{
		if node.parent.left == node{
			t.rotateRight(node)
		}else{
			t.rotateLeft(node)
		}
	}
}

func(t *TaskTreap) rotateRight(node *taskNode){
	old_parent := node.parent
	new_parent := node.parent.parent
	old_right := node.right

	// Father to Right Son
	node.right = old_parent
	
	old_parent.parent = node

	// Right Son to Left Old Parent
	old_parent.left = old_right
	if(old_right != nil){
		old_right.parent = old_parent
	}

	// New Parent to node
	node.parent = new_parent
	if new_parent != nil{
		if new_parent.left == old_parent{
			new_parent.left = node
		}else{
			new_parent.right = node
		}
	}else{
		t.root = node
	}

	t.checkRotate(node)
}

func(t *TaskTreap) rotateLeft(node *taskNode){
	
	old_parent := node.parent
	new_parent := node.parent.parent
	old_left := node.left

	//Father to Left Son
	node.left = old_parent
	old_parent.parent = node

	//Left Son to Right New Left Son
	old_parent.right = old_left
	if(old_left != nil){
		old_left.parent = old_parent
	}

	//New Parent to node
	node.parent = new_parent
	if new_parent != nil{
		if new_parent.left == old_parent{
			new_parent.left = node
		}else{
			new_parent.right = node
		}
	}else{
		t.root = node
	}

	t.checkRotate(node)

}


func compare (task_a *entities.Task, task_b *entities.Task) int{
	if task_a.Exp_time.Before(task_b.Exp_time){
		return -1
	}else if task_a.Exp_time.After(task_b.Exp_time){
		return 1
	}else{
		return 0
	}
	// TODO: Implement task priority comparison to break ties
	
}



