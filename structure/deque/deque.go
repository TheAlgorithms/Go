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

import "fmt"

type DoublyEndedQueue struct {
	Deque []any
}

func newDeque() *DoublyEndedQueue {
	return &DoublyEndedQueue{}
}

// enQueueFront adds an item at the front of Deque.
func (dq *DoublyEndedQueue) enQueueFront(n any) {
	dq.Deque = append([]any{n}, dq.Deque...)
}

// enQueueRear adds an item at the rear of Deque.
func (dq *DoublyEndedQueue) enQueueRear(n any) {
	dq.Deque = append(dq.Deque, n)
}

// deQueueFront deletes an item from front of Deque and returns it.
func (dq *DoublyEndedQueue) deQueueFront() (any, error) {
	if len(dq.Deque) == 0 {
		return nil, fmt.Errorf("DoublyEnded queue is empty so can't dequeue")
	}
	data := dq.Deque[0]
	dq.Deque = dq.Deque[1:]
	return data, nil
}

// deQueueRear deletes an item from rear of Deque and returns it.
func (dq *DoublyEndedQueue) deQueueRear() (any, error) {
	if len(dq.Deque) == 0 {
		return nil, fmt.Errorf("DoublyEnded queue is empty so can't dequeue")
	}
	data := dq.Deque[len(dq.Deque)-1]
	dq.Deque = dq.Deque[:len(dq.Deque)-1]
	return data, nil
}

// getFront gets the front item from queue.
func (dq *DoublyEndedQueue) getFront() (any, error) {
	if (len(dq.Deque)) == 0 {
		return nil, fmt.Errorf("DoublyEnded queue is empty so can't get front")
	}
	return dq.Deque[0], nil
}

// getRear gets the last item from queue.
func (dq *DoublyEndedQueue) getRear() (any, error) {
	if (len(dq.Deque)) == 0 {
		return nil, fmt.Errorf("DoublyEnded queue is empty so can't get rear")
	}
	return dq.Deque[len(dq.Deque)-1], nil
}

// isEmpty checks whether Deque is empty or not.
func (dq *DoublyEndedQueue) isEmpty() bool {
	return len(dq.Deque) == 0
}

// getLength gets the length of Deque.
func (dq *DoublyEndedQueue) getLength() int {
	return len(dq.Deque)
}
