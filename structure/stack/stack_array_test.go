package stack

import (
	"errors"
	"testing"
)

func Test_StackArray(t *testing.T) {
	stack := NewStackArray[int]()
	_, err := stack.Peek()
	if !errors.Is(err, ErrStackEmpty) {
		t.Errorf("Expected error ErrStackEmpty from Peek operation, got %v", err)
	}

	_, err = stack.Pop()
	if !errors.Is(err, ErrStackEmpty) {
		t.Errorf("Expected error ErrStackEmpty from Pop operation, got %v", err)
	}

	stack.Push(2)
	stack.Push(3)
	pop, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no errors in Pop operation, got %v", err)
	}
	if stack.Len() != 1 {
		t.Errorf("Expected stack length 1, got %d", stack.Len())
	}
	if pop != 3 {
		t.Errorf("Expected popped element to be 3, got %d", pop)
	}

	peek, err := stack.Peek()
	if err != nil {
		t.Errorf("Expected no errors in Peek operation, got %v", err)
	}
	if peek != 2 {
		t.Errorf("Expected peek operation to return element 3, got %d", peek)
	}

	stack.Clear()
	if stack.Len() != 0 && !stack.Empty() {
		t.Errorf("Expected stack to be emtpy after Clear. Got len=%d, empty=%t", stack.Len(), stack.Empty())
	}

	stack.Truncate()
	storeCapacity := cap(stack.store)
	if storeCapacity != 0 {
		t.Errorf("Expected store capacity to be zero after truncate. Got capacity=%d", storeCapacity)
	}
}
