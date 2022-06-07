// Queue Array
// description: based on `geeksforgeeks` description A Queue is a linear structure which follows a particular order in which the operations are performed.
// 	The order is First In First Out (FIFO).
// details:
// 	Queue Data Structure : https://www.geeksforgeeks.org/queue-data-structure/
//  Queue (abstract data type) : https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see queuelinkedlist.go, queuelinkedlistwithlist.go, queue_test.go

package queue

var ListQueue []any

// EnQueue it will be added new value into our list
func EnQueue(n any) {
	ListQueue = append(ListQueue, n)
}

// DeQueue it will be removed the first value that added into the list
func DeQueue() any {
	data := ListQueue[0]
	ListQueue = ListQueue[1:]
	return data
}

// FrontQueue return the Front value
func FrontQueue() any {
	return ListQueue[0]
}

// BackQueue return the Back value
func BackQueue() any {
	return ListQueue[len(ListQueue)-1]
}

// LenQueue will return the length of the queue list
func LenQueue() int {
	return len(ListQueue)
}

// IsEmptyQueue check our list is empty or not
func IsEmptyQueue() bool {
	return len(ListQueue) == 0
}
