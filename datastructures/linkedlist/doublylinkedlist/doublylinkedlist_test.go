package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	newList := DoubleLinkedList{}

	newList.AddAtBeg(1)
	newList.AddAtBeg(2)
	newList.AddAtBeg(3)

	t.Run("Test AddAtBeg", func(t *testing.T) {
		wantNext := []int{3, 2, 1}
		wantPrev := []int{1, 2, 3}
		got := []int{}

		// check from next address
		current := newList.head

		got = append(got, current.val)

		for current.next != nil {
			current = current.next
			got = append(got, current.val)
		}

		if !reflect.DeepEqual(got, wantNext) {
			t.Errorf("got: %v, want: %v", got, wantNext)
		}

		// check from prev address
		got = []int{}
		got = append(got, current.val)

		for current.prev != nil {
			current = current.prev
			got = append(got, current.val)
		}

		if !reflect.DeepEqual(got, wantPrev) {
			t.Errorf("got: %v, want: %v", got, wantPrev)
		}
	})

	newList.AddAtEnd(4)

	t.Run("Test AddAtEnd", func(t *testing.T) {
		want := []int{3, 2, 1, 4}
		got := []int{}
		current := newList.head
		got = append(got, current.val)
		for current.next != nil {
			current = current.next
			got = append(got, current.val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtBeg", func(t *testing.T) {
		want := 3
		got := newList.DelAtBeg()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtEnd", func(t *testing.T) {
		want := 4
		got := newList.DelAtEnd()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Count", func(t *testing.T) {
		want := 2
		got := newList.Count()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}
