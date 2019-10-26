// Package queue implements the Queue data structure, FIFO.
package queue

// Queue is the type that holds the methods for the Queue data structure.
type Queue struct {
	items []interface{}
}

// IsEmpty returns a bool denoting whether or not the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Enqueue adds an item into the queue.
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes an item from the queue.
func (q *Queue) Dequeue() interface{} {
	if !q.IsEmpty() {
		temp := q.GetFront()
		q.items = append(q.items[:0], q.items[1:]...)
		return temp
	}

	return nil
}

// GetFront returns the next item in the queue, but does not dequeue it.
func (q *Queue) GetFront() interface{} {
	if !q.IsEmpty() {
		return q.items[0]
	}

	return nil
}

// Clear clears the entire queue.
func (q *Queue) Clear() {
	q.items = nil
}
