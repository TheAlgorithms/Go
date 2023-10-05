/*
Package deque implements a Double Ended Queue data structure.

Description:
	Double Ended Queue is a generalized version of Queue data structure
	that allows insert and delete at both ends.

Basic Operations:
	enQueueFront(): Adds an item at the front of Deque.
	enQueueRear(): Adds an item at the rear of Deque.
	deQueueFront(): Deletes an item from front of Deque.
	deQueueRear(): Deletes an item from rear of Deque.
	getFront(): Gets the front item from queue.
	getRear(): Gets the last item from queue.
	isEmpty(): Checks whether Deque is empty or not.
	getLength(): Gets the length of Deque.

References:
	Wikipedia : https://en.wikipedia.org/wiki/Double-ended_queue
	Github: https://www.geeksforgeeks.org/deque-set-1-introduction-applications/

Author: [Sayan](https://github.com/bose-sayan)
*/

package deque

import (
	"errors"
)

var ErrEmptyDeQueue = errors.New("DoublyEnded queue is empty, so can't perform this operation")

type DoublyEndedQueue[T any] struct {
	deque []T
}

func NewDeque[T any]() *DoublyEndedQueue[T] {
	return &DoublyEndedQueue[T]{deque: make([]T, 0)}
}

// enQueueFront adds an item at the front of Deque.
func (dq *DoublyEndedQueue[T]) EnQueueFront(item T) {
	dq.deque = append([]T{item}, dq.deque...)
}

// enQueueRear adds an item at the rear of Deque.
func (dq *DoublyEndedQueue[T]) EnQueueRear(item T) {
	dq.deque = append(dq.deque, item)
}

// deQueueFront deletes an item from front of Deque and returns it.
func (dq *DoublyEndedQueue[T]) DeQueueFront() (T, error) {
	if len(dq.deque) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDeQueue
	}
	frontElement := dq.deque[0]
	dq.deque = dq.deque[1:]
	return frontElement, nil
}

// deQueueRear deletes an item from rear of Deque and returns it.
func (dq *DoublyEndedQueue[T]) DeQueueRear() (T, error) {
	if len(dq.deque) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDeQueue
	}
	rearElement := dq.deque[len(dq.deque)-1]
	dq.deque = dq.deque[:len(dq.deque)-1]
	return rearElement, nil
}

// getFront gets the front item from queue.
func (dq *DoublyEndedQueue[T]) Front() (T, error) {
	if (len(dq.deque)) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDeQueue
	}
	return dq.deque[0], nil
}

// getRear gets the last item from queue.
func (dq *DoublyEndedQueue[T]) Rear() (T, error) {
	if (len(dq.deque)) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDeQueue
	}
	return dq.deque[len(dq.deque)-1], nil
}

// isEmpty checks whether Deque is empty or not.
func (dq *DoublyEndedQueue[T]) Empty() bool {
	return len(dq.deque) == 0
}

// getLength gets the length of Deque.
func (dq *DoublyEndedQueue[T]) Length() int {
	return len(dq.deque)
}
