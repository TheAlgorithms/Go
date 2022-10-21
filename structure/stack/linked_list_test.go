package stack

import "testing"

// TestStackLinkedList for testing stack implementation using singly linked list
func TestStack_SinglyLinkedList(t *testing.T) {
	st := NewLinkedList[int]()

	st.Push(1)
	st.Push(2)

	t.Run("Stack Push", func(t *testing.T) {
		result := st.ToSlice()
		expected := []int{2, 1}
		for x := range result {
			if result[x] != expected[x] {
				t.Errorf("Expected stack elements to be %v. Current elements: %v", expected, result)
			}
		}
	})

	t.Run("Stack isEmpty", func(t *testing.T) {
		if st.Empty() {
			t.Error("Stack shouldn't be emtpy")
		}
	})

	t.Run("Stack Length", func(t *testing.T) {
		if st.Len() != 2 {
			t.Errorf("Expected stack length to be 2, got %d", st.Len())
		}
	})

	st.Pop()
	pop, _ := st.Pop()

	t.Run("Stack Pop", func(t *testing.T) {
		if pop != 1 {
			t.Errorf("Expected 1 from Pop operation, got %d", pop)
		}
	})

	st.Push(52)
	st.Push(23)
	st.Push(99)
	t.Run("Stack Peek", func(t *testing.T) {
		if val, _ := st.Peek(); val != 99 {
			t.Errorf("Expected 99 from Peek operation, got %d", val)
		}
	})
}
