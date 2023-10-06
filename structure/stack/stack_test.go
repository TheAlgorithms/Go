// Stack Test
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stackarray.go, stacklinkedlist.go, stacklinkedlistwithlist.go

package stack_test

import (
	"container/list"
	"github.com/TheAlgorithms/Go/structure/stack"
	"reflect"
	"testing"
)

// TestStackLinkedList for testing Stack with LinkedList
func TestStackLinkedList(t *testing.T) {
	var newStack stack.Stack

	newStack.Push(1)
	newStack.Push(2)

	t.Run("Stack Push", func(t *testing.T) {
		result := newStack.Show()
		expected := []any{2, 1}
		for x := range result {
			if result[x] != expected[x] {
				t.Errorf("Stack Push is not work, got %v but expected %v", result, expected)
			}
		}
	})

	t.Run("Stack isEmpty", func(t *testing.T) {
		if newStack.IsEmpty() {
			t.Error("Stack isEmpty is returned true but expected false", newStack.IsEmpty())
		}

	})

	t.Run("Stack Length", func(t *testing.T) {
		if newStack.Length() != 2 {
			t.Error("Stack Length should be 2 but got", newStack.Length())
		}
	})

	newStack.Pop()
	pop := newStack.Pop()

	t.Run("Stack Pop", func(t *testing.T) {
		if pop != 1 {
			t.Error("Stack Pop should return 1 but is returned", pop)
		}

	})

	newStack.Push(52)
	newStack.Push(23)
	newStack.Push(99)
	t.Run("Stack Peek", func(t *testing.T) {
		if newStack.Peek() != 99 {
			t.Error("Stack Peak should return 99 but got ", newStack.Peek())
		}
	})
}

// TestStackArray for testing Stack with Array
func TestStackArray(t *testing.T) {
	newStack := stack.NewStack[int]()
	t.Run("Stack With Array", func(t *testing.T) {

		newStack.Push(2)
		newStack.Push(3)

		t.Run("Stack Push", func(t *testing.T) {
			var stackElements []any
			for i := 0; i < 2; i++ {
				poppedElement := newStack.Pop()
				stackElements = append(stackElements, poppedElement)
			}

			if !reflect.DeepEqual([]any{3, 2}, stackElements) {
				t.Errorf("Stack Push is not work we expected %v but got %v", []any{3, 2}, newStack)
			}

			newStack.Push(2)
			newStack.Push(3)
		})

		pop := newStack.Pop()

		t.Run("Stack Pop", func(t *testing.T) {
			if newStack.Length() == 2 && pop != 3 {
				t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
			}
		})

		newStack.Push(2)
		newStack.Push(83)

		t.Run("Stack Peak", func(t *testing.T) {
			if newStack.Peek() != 83 {
				t.Errorf("Stack Peek is not work we expected %v but got %v", 83, newStack.Peek())
			}
		})

		t.Run("Stack Length", func(t *testing.T) {
			if newStack.Length() != 3 {
				t.Errorf("Stack Length is not work we expected %v but got %v", 3, newStack.Length())
			}
		})

		t.Run("Stack Empty", func(t *testing.T) {
			if newStack.IsEmpty() == true {
				t.Errorf("Stack Empty is not work we expected %v but got %v", false, newStack.IsEmpty())
			}

			newStack.Pop()
			newStack.Pop()
			newStack.Pop()

			if newStack.IsEmpty() == false {
				t.Errorf("Stack Empty is not work we expected %v but got %v", true, newStack.IsEmpty())
			}
		})
	})
}

// TestStackLinkedListWithList for testing Stack with Container/List Library (STL)
func TestStackLinkedListWithList(t *testing.T) {
	stackList := &stack.SList{
		Stack: list.New(),
	}

	t.Run("Stack Push", func(t *testing.T) {

		stackList.Push(2)
		stackList.Push(3)

		if stackList.Length() != 2 {
			t.Errorf("Stack Push is not work we expected %v but got %v", 2, stackList.Length())
		}
	})

	t.Run("Stack Pop", func(t *testing.T) {
		pop, _ := stackList.Pop()

		if stackList.Length() == 1 && pop != 3 {
			t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
		}
	})

	t.Run("Stack Peak", func(t *testing.T) {

		stackList.Push(2)
		stackList.Push(83)
		peak, _ := stackList.Peek()
		if peak != 83 {
			t.Errorf("Stack Peak is not work we expected %v but got %v", 83, peak)
		}
	})

	t.Run("Stack Length", func(t *testing.T) {
		if stackList.Length() != 3 {
			t.Errorf("Stack Length is not work we expected %v but got %v", 3, stackList.Length())
		}
	})

	t.Run("Stack Empty", func(t *testing.T) {
		if stackList.IsEmpty() == true {
			t.Errorf("Stack Empty is not work we expected %v but got %v", false, stackList.IsEmpty())
		}

		d1, err := stackList.Pop()
		d2, _ := stackList.Pop()
		d3, _ := stackList.Pop()

		if err != nil {
			t.Errorf("got an unexpected error %v, pop1: %v, pop2: %v, pop3: %v", err, d1, d2, d3)
		}

		if stackList.IsEmpty() == false {
			t.Errorf("Stack Empty is not work we expected %v but got %v", true, stackList.IsEmpty())
		}
	})
}
