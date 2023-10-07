// description: Double Ended Queue is a generalized version of Queue data structure that allows insert and delete at both ends.
// basic Operations:
//	EnQueueFront(): Adds an item at the front of Deque.
//	EnQueueRear(): Adds an item at the rear of Deque.
//	DeQueueFront(): Deletes an item from front of Deque.
//	DeQueueRear(): Deletes an item from rear of Deque.
//	Front(): Gets the front item from queue.
//	Rear(): Gets the last item from queue.
//	IsEmpty(): Checks whether Deque is empty or not.
//	Length(): Gets the length of Deque.
// References:
//	Wikipedia : https://en.wikipedia.org/wiki/Double-ended_queue
//	Github: https://www.geeksforgeeks.org/deque-set-1-introduction-applications/
// author [Sayan](https://github.com/bose-sayan)

// Package deque implements a Double Ended Queue data structure.
package deque

import (
	"errors"
)

// Custom error for handling cases when some dequeuing operation is performed on an empty deque.
var ErrEmptyDeQueue = errors.New("DoublyEnded queue is empty, so can't perform this operation")

type DoublyEndedQueue[T any] struct {
	deque []T
}

// NewDeque returns a new DoublyEndedQueue.
func NewDeque[T any]() *DoublyEndedQueue[T] {
	return &DoublyEndedQueue[T]{deque: make([]T, 0)}
}

// EnQueueFront adds an item at the front of Deque.
func (dq *DoublyEndedQueue[T]) EnQueueFront(item T) {
	dq.deque = append([]T{item}, dq.deque...)
}

// EnQueueRear adds an item at the rear of Deque.
func (dq *DoublyEndedQueue[T]) EnQueueRear(item T) {
	dq.deque = append(dq.deque, item)
}

// DeQueueFront deletes an item from front of Deque and returns it.
func (dq *DoublyEndedQueue[T]) DeQueueFront() (T, error) {
	if len(dq.deque) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDeQueue
	}
	frontElement := dq.deque[0]
	dq.deque = dq.deque[1:]
	return frontElement, nil
}

// DeQueueRear deletes an item from rear of Deque and returns it.
func (dq *DoublyEndedQueue[T]) DeQueueRear() (T, error) {
	if len(dq.deque) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDeQueue
	}
	rearElement := dq.deque[len(dq.deque)-1]
	dq.deque = dq.deque[:len(dq.deque)-1]
	return rearElement, nil
}

// Front gets the front item from queue.
func (dq *DoublyEndedQueue[T]) Front() (T, error) {
	if (len(dq.deque)) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDeQueue
	}
	return dq.deque[0], nil
}

// Rear gets the last item from queue.
func (dq *DoublyEndedQueue[T]) Rear() (T, error) {
	if (len(dq.deque)) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDeQueue
	}
	return dq.deque[len(dq.deque)-1], nil
}

// IsEmpty checks whether Deque is empty or not.
func (dq *DoublyEndedQueue[T]) IsEmpty() bool {
	return len(dq.deque) == 0
}

// Length gets the length of Deque.
func (dq *DoublyEndedQueue[T]) Length() int {
	return len(dq.deque)
}
