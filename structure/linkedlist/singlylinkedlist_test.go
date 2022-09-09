package linkedlist

import (
	"reflect"
	"testing"
)

func TestSingly(t *testing.T) {
	list := NewSingly[int]()
	list.AddAtBeg(1)
	list.AddAtBeg(2)
	list.AddAtBeg(3)

	t.Run("Test AddAtBeg()", func(t *testing.T) {
		want := []any{3, 2, 1}
		got := []any{}
		current := list.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list.AddAtEnd(4)

	t.Run("Test AddAtEnd()", func(t *testing.T) {
		want := []any{3, 2, 1, 4}
		got := []any{}
		current := list.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtBeg()", func(t *testing.T) {
		want := any(3)
		got, ok := list.DelAtBeg()
		if !ok {
			t.Error("unexpected not-ok")
		}
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtEnd()", func(t *testing.T) {
		want := any(4)
		got, ok := list.DelAtEnd()
		if !ok {
			t.Error("unexpected not-ok")
		}
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Count()", func(t *testing.T) {
		want := 2
		got := list.Count()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list2 := Singly[int]{}
	list2.AddAtBeg(1)
	list2.AddAtBeg(2)
	list2.AddAtBeg(3)
	list2.AddAtBeg(4)
	list2.AddAtBeg(5)
	list2.AddAtBeg(6)

	t.Run("Test Reverse()", func(t *testing.T) {
		want := []any{1, 2, 3, 4, 5, 6}
		got := []any{}
		list2.Reverse()
		current := list2.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test ReversePartition()", func(t *testing.T) {
		want := []any{1, 5, 4, 3, 2, 6}
		got := []any{}
		err := list2.ReversePartition(2, 5)

		if err != nil {
			t.Errorf("Incorrect boundary conditions entered%v", err)
		}
		current := list2.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
