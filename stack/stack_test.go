package stack

import "testing"

func TestStack(t *testing.T) {
	// helper function to avoid repetition
	assertMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	}
	t.Run("Should be able to add elements to the stack", func(t *testing.T) {
		var s = Stack{}
		s.Unshift("John")
		s.Unshift("Mary")
		s.Unshift("Brian")
		size := s.Size()
		got := size
		want := 3
		assertMessage(t, got, want)
	})
	t.Run("Should be able to peek elements in the stack", func(t *testing.T) {
		var s = Stack{}
		s.Unshift("John")
		s.Unshift("Mary")
		ele := s.Peek()
		want := []string{"Mary", "John"}
		for i := 0; i < s.Size(); i++ {
			if ele[i] != want[i] {
				t.Errorf("got %s but expected %s", ele[i], want[i])
			}
		}
	})
	t.Run("Should be able to remove elements from the stack", func(t *testing.T) {
		var s = Stack{}
		s.Unshift("John")
		s.Unshift("Mary")
		s.Shift()
		size := s.Size()
		got := size
		want := 1
		assertMessage(t, got, want)
	})
	t.Run("Should be able to know the number of elements in the stack", func(t *testing.T) {
		var s = Stack{}
		size := s.Size()
		got := size
		want := 0
		assertMessage(t, got, want)
	})
	t.Run("Should be able to know if the stack is empty", func(t *testing.T) {
		var s = Stack{}
		isEmpty := s.Empty()
		got := isEmpty
		want := true
		if got != want {
			t.Errorf("got %v but expected %v", got, want)
		}
	})
}
