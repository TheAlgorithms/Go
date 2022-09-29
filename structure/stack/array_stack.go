// Stack Array
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stacklinkedlist.go, stacklinkedlistwithlist.go, stack_test.go

package stack

/*
The methods can also be implemented directly on the slice.
```
type ArrayStack[T any] []T
```
However, this exposes the underlaying storage (slice) outside the package.
A struct is used instead, so that the underlying storage is not accessible outside the package.
*/

// ArrayStack is an implementation of stack with slice as underlying storage.
// ```
// stack := new(ArrayStack[int])
// ```
type ArrayStack[T any] struct {
	store []T
}

func NewStackArray[T any]() *ArrayStack[T] {
	return new(ArrayStack[T])
}

// Push inserts a new element to the stack
// Push on a nil stack will panic
func (s *ArrayStack[T]) Push(val T) {
	s.store = append(s.store, val)
}

// Peek the last inserted element without removing it from the stack
// Return type is a pointer to represent an optional value.
// If the stack is empty, nil value is returned
func (s *ArrayStack[T]) Peek() *T {
	if s.Empty() {
		return nil
	}
	//do not return &s.store[s.Len()-1] directly. A Pop(), followed by another Push() can overwrite
	// that memory location in the underlaying array of the slice. Assigning it to a new variable creates a copy.
	element := s.store[s.Len()-1]
	return &element
}

func (s *ArrayStack[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.store)
}

func (s *ArrayStack[T]) Empty() bool {
	return s.Len() == 0
}

// Pop returns last inserted element and removes it from the underlaying storage
// The return type is a pointer to represent optional value. If the stack is empty, nil is returned
func (s *ArrayStack[T]) Pop() *T {
	if s.Empty() {
		return nil
	}
	element := s.store[s.Len()-1]
	s.store = s.store[:s.Len()-1]
	return &element
}

// Clear removes all elements.
// The allocated capacity remains the same and will be reused for subsequent push operations
func (s *ArrayStack[T]) Clear() {
	if s == nil {
		return
	}
	s.store = s.store[:0]
}

// Truncate removes all elements and underlaying storage
func (s *ArrayStack[T]) Truncate() {
	if s == nil {
		return
	}
	s.store = nil
}
