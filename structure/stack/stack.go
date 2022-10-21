// Stack is a linear data structure in which the push and pop operations occur only at one end of the structure, referred to as the top of the stack.
// The order in which an element added to or removed from a stack is described as last in, first out (LIFO).
//
//	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
//	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
package stack

// Interface defines the specifications for a stack implementation.
// Currently, there are 3 implmentatios of this stack.Interface in this package.
// stack.Array : Uses a slice as its underlying storage
// stack.LinkedList : Uses linked list as its underlying storage
// stack.DoublyLinkedList : Uses the doubly linked list provided by std library `container/list` as the underlying storage.
type Interface[T any] interface {
	// Push value to the top of the stack
	Push(T)

	// Pop returns last inserted element and removes it from the underlying storage
	// If the stack is empty, ErrStackEmpty error is returned
	Pop() (T, error)

	// Peek the last inserted element without removing it from the stack
	// If the stack is empty, ErrStackEmpty error is returned
	Peek() (T, error)

	// Len returns the number of elements in the stack
	Len() int

	// Empty returns true if stack has no elements and false otherwise.
	Empty() bool

	// Clear removes all elements by re-initializing the underlying storage
	// If the underlying storage of the stack is an array/slice (created using `stack.NewArray[T]()`),
	// the allocated capacity remains the same and will be reused for subsequent push operations
	Clear()

	// ToSlice returns the elements of stack as a slice
	ToSlice() []T
}
