// Stack Array
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stacklinkedlist.go, stacklinkedlistwithlist.go, stack_test.go

package stack

type Array[T any] struct {
	elements []T
}

// NewStack creates and returns a new stack.
func NewStack[T any]() *Array[T] {
	return &Array[T]{}
}

// Push adds an element to the top of the stack.
func (s *Array[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Size returns the number of elements in the stack.
func (s *Array[T]) Length() int {
	return len(s.elements)
}

// Peek returns the top element of the stack without removing it.
func (s *Array[T]) Peek() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue // Stack is empty
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Array[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Pop removes and returns the top element from the stack.
func (s *Array[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue // Stack is empty
	}
	index := len(s.elements) - 1
	popped := s.elements[index]
	s.elements = s.elements[:index]
	return popped
}
