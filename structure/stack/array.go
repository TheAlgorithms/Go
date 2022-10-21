package stack

import "errors"

var ErrStackEmpty = errors.New("stack is empty")

// Array is an implementation of stack with slice as underlying storage.
// ```
// stack := stack.NewArray[int]()
// ```
// Note that the type `Array` could also be implemented directly using a slice.
// ```
// type Array[T any] []T
// ```
// However, this exposes the underlying storage (slice) outside the package.
// A struct is used instead, so that the underlying storage is not accessible
// outside the package.
type Array[T any] struct {
	store []T
}

func NewArray[T any]() Interface[T] {
	return new(Array[T])
}

// Push inserts a new element to the stack
// Push on a nil stack will panic
func (s *Array[T]) Push(val T) {
	s.store = append(s.store, val)
}

// Peek the last inserted element without removing it from the stack
// If the stack is empty, ErrStackEmpty error is returned
func (s *Array[T]) Peek() (T, error) {
	var element T
	if s.Empty() {
		return element, ErrStackEmpty
	}
	return s.store[s.Len()-1], nil
}

func (s *Array[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.store)
}

func (s *Array[T]) Empty() bool {
	return s.Len() == 0
}

// Pop returns last inserted element and removes it from the underlying storage
// If the stack is empty, ErrStackEmpty error is returned
func (s *Array[T]) Pop() (T, error) {
	var element T
	if s.Empty() {
		return element, ErrStackEmpty
	}
	element = s.store[s.Len()-1]
	s.store = s.store[:s.Len()-1]
	return element, nil
}

// Clear removes all elements.
// The allocated capacity remains the same and will be reused for subsequent push operations
func (s *Array[T]) Clear() {
	if s == nil {
		return
	}
	s.store = s.store[:0]
}

func (s *Array[T]) ToSlice() []T {
	out := make([]T, len(s.store))
	copy(out, s.store)
	return out
}
