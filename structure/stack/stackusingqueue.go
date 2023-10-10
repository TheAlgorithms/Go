// Stack Implementation Using Queue
// description: Implemenating the operations of Stack using Queue
// details:
/*
* Based on the Geeks-for-Geeks topic: [Stack Using Queue](https://www.geeksforgeeks.org/implement-stack-using-queue/)
 */
// author(s): [Anuj Kumar Karmakar](https://www.github.com/anujkkarmakar)
// see also: stackusingqueue_test.go

package stack

import (
	"fmt"
)

// Queue is a struct that represents a queue data structure
type Queue struct {
	items []int
}

// NewQueue returns a new empty queue
func NewQueue() *Queue {
	return &Queue{items: []int{}}
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Printf("Stack is Empty")
		return -1 // indicates empty queue
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty returns true if the queue is empty, false otherwise
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// Stack is a struct that represents a stack data structure using two queues
type Stack1 struct {
	q1 *Queue // main queue
	q2 *Queue // auxiliary queue
}

// NewStack returns a new empty stack
func NewStackUsingQueue() *Stack1 {
	return &Stack1{q1: NewQueue(), q2: NewQueue()}
}

// Push adds an item to the top of the stack
func (s *Stack1) Push(item int) {
	// enqueue the item to the main queue
	s.q1.Enqueue(item)
}

// Pop removes and returns the item at the top of the stack
func (s *Stack1) Pop() int {
	// if the main queue is empty, return -1 to indicate empty stack
	if s.q1.IsEmpty() {
		return -1
	}
	// dequeue all items from the main queue except the last one, and enqueue them to the auxiliary queue
	for s.q1.Size() > 1 {
		s.q2.Enqueue(s.q1.Dequeue())
	}
	// dequeue and store the last item from the main queue, which is the top of the stack
	item := s.q1.Dequeue()
	// swap the main and auxiliary queues
	s.q1, s.q2 = s.q2, s.q1
	// return the top item
	return item
}

// IsEmpty returns true if the stack is empty, false otherwise
func (s *Stack1) IsEmpty() bool {
	// return true if the main queue is empty, false otherwise
	return s.q1.IsEmpty()
}

// Size returns the number of items in the stack
func (s *Stack1) Size() int {
	// return the size of the main queue
	return s.q1.Size()
}
