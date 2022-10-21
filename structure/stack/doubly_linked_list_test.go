package stack

import (
	"testing"
)

// TestStackLinkedListWithList for testing Stack with Container/List Library (STL)
func TestStackLinkedListWithList(t *testing.T) {
	st := NewDoublyLinkedList[int]()

	t.Run("Stack Push", func(t *testing.T) {

		st.Push(2)
		st.Push(3)

		if st.Len() != 2 {
			t.Errorf("Expected 2 elements in the stack, found %d", st.Len())
		}
	})

	t.Run("Stack Pop", func(t *testing.T) {
		pop, _ := st.Pop()

		if pop != 3 {
			t.Errorf("Expected 3 from Pop operation, got %d", pop)
		}

		if st.Len() != 1 {
			t.Errorf("Expected stack length to be 1 after Pop operation, got %d", st.Len())
		}
	})

	t.Run("Stack Peek", func(t *testing.T) {
		st.Push(2)
		st.Push(83)
		peek, _ := st.Peek()
		if peek != 83 {
			t.Errorf("Expected value 83 from Peek operation, got %d", peek)
		}
	})

	t.Run("Stack Len", func(t *testing.T) {
		if st.Len() != 3 {
			t.Errorf("Expected stack length to be 3, got %d", st.Len())
		}
	})

	t.Run("Stack Empty", func(t *testing.T) {
		if st.Empty() {
			t.Error("Stack should not be empty")
		}
		st.Clear()
		if !st.Empty() {
			t.Error("Stack is expected to be empty")
		}
	})
}
