package stack

import (
	"testing"
)

func Test_StackArray(t *testing.T) {
	stack := NewStackArray[int]()
	if stack.Peek() != nil {
		t.Error("Expected nil value from peak operation")
	}
	if stack.Pop() != nil {
		t.Error("Expected nil value from pop operation")
	}
	stack.Push(2)
	stack.Push(3)
	pop := stack.Pop()
	if stack.Len() != 1 {
		t.Errorf("Expected stack length 1, got %d", stack.Len())
	}
	if *pop != 3 {
		t.Errorf("Expected popped element to be 3, got %d", *pop)
	}

	peek := stack.Peek()
	if *peek != 2 {
		t.Errorf("Expected peek operation return element 3, got %d", *peek)
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
