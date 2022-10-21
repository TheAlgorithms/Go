package stack

import (
	"container/list"
)

// doublyLinkedList is an implementation of stack.Interface using the doubly linked list provided by `container/list` as its underlying storage.
type doublyLinkedList[T any] struct {
	stack *list.List
}

func NewDoublyLinkedList[T any]() Interface[T] {
	return &doublyLinkedList[T]{
		stack: list.New(),
	}
}

// Push add a value into our stack
func (dl *doublyLinkedList[T]) Push(val T) {
	dl.stack.PushFront(val)
}

// Peek return last inserted element(top of the stack) without removing it from the stack
// If the stack is empty, ErrStackEmpty error is returned
func (dl *doublyLinkedList[T]) Peek() (T, error) {
	var result T
	if dl.Empty() {
		return result, ErrStackEmpty
	}

	element := dl.stack.Front()
	if element == nil {
		return result, ErrStackEmpty
	}

	result = element.Value.(T)
	return result, nil
}

// Pop is return last value that insert into our stack
// also it will remove it in our stack
func (dl *doublyLinkedList[T]) Pop() (T, error) {
	var result T
	if dl.Empty() {
		return result, ErrStackEmpty
	}

	element := dl.stack.Front()
	if element == nil {
		return result, ErrStackEmpty
	}

	dl.stack.Remove(element)
	result = element.Value.(T)
	return result, nil
}

// Length returns the number of elements in the stack
func (dl *doublyLinkedList[T]) Len() int {
	if dl == nil {
		return 0
	}
	return dl.stack.Len()
}

// Empty returns true if stack has no elements and false otherwise.
func (dl *doublyLinkedList[T]) Empty() bool {
	if dl == nil {
		return true
	}
	return dl.stack.Len() == 0
}

// Clear initializes the underlying storage with a new empty doubly linked list, thus clearing the underlying storage.
func (dl *doublyLinkedList[T]) Clear() {
	if dl == nil {
		return
	}
	dl.stack = list.New()
}

// ToSlice returns the elements of stack as a slice
func (dl *doublyLinkedList[T]) ToSlice() []T {
	var result []T
	if dl == nil {
		return result
	}

	for e := dl.stack.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(T))
	}
	return result
}
