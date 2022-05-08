// Queue Linked-List
// description: based on `geeksforgeeks` description A Queue is a linear structure which follows a particular order in which the operations are performed.
// 	The order is First In First Out (FIFO).
// details:
// 	Queue Data Structure : https://www.geeksforgeeks.org/queue-data-structure/
//  Queue (abstract data type) : https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see queuearray.go, queuelinkedlistwithlist.go, queue_test.go

package queue

// Node will be store the value and the next node as well
type Node struct {
	Data any
	Next *Node
}

// Queue structure is tell us what our head is and what tail should be with length of the list
type Queue struct {
	head   *Node
	tail   *Node
	length int
}

// enqueue it will be added new value into queue
func (ll *Queue) enqueue(n any) {
	var newNode Node // create new Node
	newNode.Data = n // set the data

	if ll.tail != nil {
		ll.tail.Next = &newNode
	}

	ll.tail = &newNode

	if ll.head == nil {
		ll.head = &newNode
	}
	ll.length++
}

// dequeue it will be removed the first value into queue (First In First Out)
func (ll *Queue) dequeue() any {
	if ll.isEmpty() {
		return -1 // if is empty return -1
	}
	data := ll.head.Data

	ll.head = ll.head.Next

	if ll.head == nil {
		ll.tail = nil
	}

	ll.length--
	return data
}

// isEmpty it will check our list is empty or not
func (ll *Queue) isEmpty() bool {
	return ll.length == 0
}

// len is return the length of queue
func (ll *Queue) len() int {
	return ll.length
}

// frontQueue it will return the front data
func (ll *Queue) frontQueue() any {
	return ll.head.Data
}

// backQueue it will return the back data
func (ll *Queue) backQueue() any {
	return ll.tail.Data
}
