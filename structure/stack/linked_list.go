package stack

type node[T any] struct {
	Val  T
	Next *node[T]
}

// linkedList implements stack.Interface using a singly linked list as the underlying storage
type linkedList[T any] struct {
	top    *node[T]
	length int
}

func NewLinkedList[T any]() Interface[T] {
	return new(linkedList[T])
}

// Push value to the top of the stack
func (ll *linkedList[T]) Push(n T) {
	newStack := new(node[T])

	newStack.Val = n
	newStack.Next = ll.top

	ll.top = newStack
	ll.length++
}

// Pop returns last inserted element and removes it from the underlying storage
// If the stack is empty, ErrStackEmpty error is returned
func (ll *linkedList[T]) Pop() (T, error) {
	var element T
	if ll.Empty() {
		return element, ErrStackEmpty
	}
	element = ll.top.Val
	ll.top = ll.top.Next
	ll.length--
	return element, nil
}

// Empty returns true if stack has no elements and false otherwise.
func (ll *linkedList[T]) Empty() bool {
	return ll.length == 0
}

// Len returns length of the stack
func (ll *linkedList[T]) Len() int {
	return ll.length
}

// Peek return last inserted element(top of the stack) without removing it from the stack
// If the stack is empty, ErrStackEmpty error is returned
func (ll *linkedList[T]) Peek() (T, error) {
	var element T
	if ll == nil || ll.length == 0 {
		return element, ErrStackEmpty
	}
	return ll.top.Val, nil
}

// ToSlice returns the elements of stack as a slice
func (ll *linkedList[T]) ToSlice() []T {
	var elements []T
	if ll == nil {
		return elements
	}

	current := ll.top
	for current != nil {
		elements = append(elements, current.Val)
		current = current.Next
	}
	return elements
}

func (ll *linkedList[T]) Clear() {
	if ll == nil {
		return
	}
	ll.top = nil
	ll.length = 0
}
