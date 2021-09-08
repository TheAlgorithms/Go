package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	t.Run("Test NewNode()", func(t *testing.T) {
		var value = 1
		expected := &Node{
			val:  value,
			next: nil,
			prev: nil,
		}
		actual := NewNode(value)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got: %v, want: %v", actual, expected)
		}
	})
}

func TestDoubleLinkedList_AddAtBeg(t *testing.T) {
	ll := DoubleLinkedList{}

	t.Run("Test AddAtBeg()", func(t *testing.T) {
		var value = 1
		expected := &Node{
			val:  value,
			next: nil,
			prev: nil,
		}
		ll.AddAtBeg(value)
		actual := ll.head
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got: %v, want: %v", actual, expected)
		}
	})
}

func TestDoubleLinkedList_AddAtEnd(t *testing.T) {
	ll := DoubleLinkedList{}

	t.Run("Test AddAtEnd()", func(t *testing.T) {
		var value = 1
		expected := &Node{
			val:  value,
			next: nil,
			prev: nil,
		}
		ll.AddAtBeg(2)
		ll.AddAtEnd(value)

		cur := ll.head
		for ; cur.next != nil; cur = cur.next {
		}
		actual := cur
		if !reflect.DeepEqual(actual.val, expected.val) {
			t.Errorf("got: %v, want: %v", actual, expected)
		}
	})
}

func TestDoubleLinkedList_DelAtBeg(t *testing.T) {
	ll := DoubleLinkedList{}

	t.Run("Test DelAtBeg()", func(t *testing.T) {
		ll.AddAtBeg(2)
		ll.AddAtBeg(1)
		ll.AddAtBeg(3)
		ll.AddAtBeg(2)
		ll.AddAtBeg(2)
		ll.DelAtBeg()

		expectedSize := 4
		expectedValue := 1
		cur := ll.head
		var size = 0
		for {
			size++
			if cur.next == nil {
				break
			}
			cur = cur.next
		}

		if !reflect.DeepEqual(size, expectedSize) {
			t.Errorf("got doublylinkedlist size: %v, want doublylinkedlist size: %v", size, expectedSize)
			t.Errorf("got: %v, want: %v", ll.head.val, expectedValue)
		}
	})
}

func TestDoubleLinkedList_DelAtEnd(t *testing.T) {
	ll := DoubleLinkedList{}

	t.Run("Test DelAtEnd()", func(t *testing.T) {
		ll.AddAtBeg(2)
		ll.AddAtBeg(1)
		ll.AddAtBeg(3)
		ll.AddAtBeg(2)
		ll.AddAtBeg(6)
		ll.DelAtEnd()

		expectedSize := 4
		expectedValue := 1
		cur := ll.head
		var size = 0
		for {
			size++
			if cur.next == nil {
				break
			}
			cur = cur.next
		}

		if !reflect.DeepEqual(size, expectedSize) || !reflect.DeepEqual(cur.val, expectedValue) {
			t.Errorf("got doublylinkedlist size: %v, want doublylinkedlist size: %v", size, expectedSize)
			t.Errorf("got: %v, want: %v", cur.val, expectedValue)
		}
	})
}

func TestDoubleLinkedList_Count(t *testing.T) {
	ll := DoubleLinkedList{}

	t.Run("Test Count()", func(t *testing.T) {
		ll.AddAtBeg(2)
		ll.AddAtBeg(1)
		ll.AddAtBeg(3)
		ll.AddAtBeg(2)
		ll.AddAtBeg(6)

		actual := ll.Count()
		expected := 5

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got: %v, want: %v", actual, expected)
		}
	})
}

func TestDoubleLinkedList_Reverse(t *testing.T) {
	ll := DoubleLinkedList{}

	t.Run("Test Reverse()", func(t *testing.T) {
		ll.AddAtBeg(2)
		ll.AddAtBeg(1)
		ll.AddAtBeg(3)
		ll.AddAtBeg(2)
		ll.AddAtBeg(6)

		expectedHeadValue := 6
		expectedTailValue := 2
		cur := ll.head
		for {
			if cur.next == nil {
				break
			}
			cur = cur.next
		}

		if !reflect.DeepEqual(expectedHeadValue, ll.head.val) || !reflect.DeepEqual(expectedTailValue, cur.val) {
			t.Errorf("got head value: %v, want head value: %v", ll.head.val, expectedHeadValue)
			t.Errorf("got tail value: %v, want tail value: %v", cur.val, expectedTailValue)
		}
	})
}
