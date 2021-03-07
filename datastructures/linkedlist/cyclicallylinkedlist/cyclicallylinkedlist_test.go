package cyclicallylinkedlist

import (
	"reflect"
	"testing"
)

type testCase struct {
	param        int
	wantToReturn int
	listCount    int
}

func fillList(list *ClList, n int) {
	for i := 1; i <= n; i++ {
		list.Add(i)
	}
}

// 100% coverage.
func TestLinkedList(t *testing.T) {
	list := NewList()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	t.Run("Test Add()", func(t *testing.T) {
		want := []interface{}{1, 2, 3}
		var got []interface{}
		var start *ClNode
		start = list.currentItem

		for i := 0; i < list.size; i++ {
			got = append(got, start.val)
			start = start.next
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Walk()", func(t *testing.T) {
		want := 1
		got := list.Walk()

		if got.val != want {
			t.Errorf("got: %v, want: nil", got)
		}
	})

	t.Run("Test Rotate()", func(t *testing.T) {
		testCases := []testCase{
			{1, 2, 0},
			{3, 2, 0},
			{6, 2, 0},
			{7, 3, 0},
			{-2, 1, 0},
			{5, 3, 0},
			{8, 2, 0},
			{-8, 3, 0},
		}
		for idx, tCase := range testCases {
			list.Rotate(tCase.param)
			got := list.currentItem.val
			if got != tCase.wantToReturn {
				t.Errorf("got: %v, want: %v for test id %v", got, tCase.wantToReturn, idx)
			}
		}

	})

	t.Run("Test Delete()", func(t *testing.T) {
		want := 1
		list.Delete()
		got := list.currentItem.val
		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Destroy()", func(t *testing.T) {
		list.Destroy()
		got := list.currentItem
		if got != nil {
			t.Errorf("got: %v, want: nil", got)
		}
	})

	t.Run("Test JosephusProblem()", func(t *testing.T) {
		testCases := []testCase{
			{5, 4, 8},
			{3, 8, 8},
			{8, 5, 8},
			{8, 5, 8},
			{2, 14, 14},
			{13, 56, 58},
			{7, 5, 5},
		}
		for _, tCase := range testCases {
			list.Destroy()
			fillList(list, tCase.listCount)
			got := JosephusProblem(list, tCase.param)
			if got != tCase.wantToReturn {
				t.Errorf("got: %v, want: %v", got, tCase.wantToReturn)
			}
		}
	})
}
