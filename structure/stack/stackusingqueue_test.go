// Stack Implementation Using Queue Test
// description: Implemenating the operations of Stack using Queue
// details:
/*
* Based on the Geeks-for-Geeks topic: [Stack Using Queue](https://www.geeksforgeeks.org/implement-stack-using-queue/)
 */
// author(s): [Anuj Kumar Karmakar](https://www.github.com/anujkkarmakar)
// see also: stackusingqueue.go

package stack

import (
	"testing"
)

// TestPush tests the Push method of the stack
func TestPush(t *testing.T) {
	// create a new stack
	s := NewStackUsingQueue()
	// push some items to the stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	// check if the stack size is correct
	if s.Size() != 3 {
		t.Errorf("Expected stack size to be 3, got %d", s.Size())
	}
}

// TestPop tests the Pop method of the stack
func TestPop(t *testing.T) {
	// create a new stack
	s := NewStackUsingQueue()
	// push some items to the stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	// pop and check the items from the stack
	if s.Pop() != 30 {
		t.Errorf("Expected pop to return 30, got %d", s.Pop())
	}
	if s.Pop() != 20 {
		t.Errorf("Expected pop to return 20, got %d", s.Pop())
	}
	if s.Pop() != 10 {
		t.Errorf("Expected pop to return 10, got %d", s.Pop())
	}
	// check if the stack is empty
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, got %v", s.IsEmpty())
	}
	// check if popping from an empty stack returns -1
	if s.Pop() != -1 {
		t.Errorf("Expected pop to return -1, got %d", s.Pop())
	}
}

// TestIsEmpty tests the IsEmpty method of the stack
func TestIsEmpty(t *testing.T) {
	// create a new stack
	s := NewStackUsingQueue()
	// check if the stack is initially empty
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, got %v", s.IsEmpty())
	}
	// push an item to the stack
	s.Push(10)
	// check if the stack is not empty
	if s.IsEmpty() {
		t.Errorf("Expected stack to be not empty, got %v", s.IsEmpty())
	}
}

// TestSize tests the Size method of the stack
func TestSize(t *testing.T) {
	// create a new stack
	s := NewStackUsingQueue()
	// check if the stack size is initially zero
	if s.Size() != 0 {
		t.Errorf("Expected stack size to be 0, got %d", s.Size())
	}
	// push some items to the stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	// check if the stack size is correct
	if s.Size() != 3 {
		t.Errorf("Expected stack size to be 3, got %d", s.Size())
	}
}
