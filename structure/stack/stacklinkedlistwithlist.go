// Stack Linked-List with standard library (Container/List)
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stacklinkedlist.go, stackarray.go, stack_test.go

package stack

import (
	"container/list"
	"fmt"
)

// SList is our struct that point to stack with container/list.List library
type SList struct {
	stack *list.List
}

// Push add a value into our stack
func (sl *SList) Push(val interface{}) {
	sl.stack.PushFront(val)
}

// Peak is return last value that insert into our stack
func (sl *SList) Peak() (interface{}, error) {
	if !sl.Empty() {
		element := sl.stack.Front()
		return element.Value, nil
	}
	return "", fmt.Errorf("stack list is empty")
}

// Pop is return last value that insert into our stack
//also it will remove it in our stack
func (sl *SList) Pop() (interface{}, error) {
	if !sl.Empty() {
		// get last element that insert into stack
		element := sl.stack.Front()
		// remove element in stack
		sl.stack.Remove(element)
		// return element value
		return element.Value, nil
	}
	return "", fmt.Errorf("stack list is empty")
}

// Length return length of our stack
func (sl *SList) Length() int {
	return sl.stack.Len()
}

// Empty check our stack has value or not
func (sl *SList) Empty() bool {
	// check our stack is empty or not
	// if is 0 it means our stack is empty otherwise is not empty
	return sl.stack.Len() == 0
}
