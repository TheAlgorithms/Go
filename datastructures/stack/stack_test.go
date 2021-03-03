package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	var stackTestData = []struct {
		description string
		operation   func(s *Stack)
		expected    Stack
	}{
		{
			"test 1 push elements to stack",
			func(s *Stack) {
				s.Push(1)
				s.Push(2)
				s.Push(3)
				s.Push(4)
				s.Push(4)
				s.Push(4)
			},
			[]interface{}{1, 2, 3, 4, 4, 4},
		},
		{
			"test 1 push and pop elements to stack",
			func(s *Stack) {
				s.Push(1)
				s.Push(2)
				s.Push(3)
				s.Push(4)
				s.Push(4)
				s.Push(4)
				s.Pop()
				s.Pop()
				s.Pop()
				s.Pop()
			},
			[]interface{}{1, 2},
		},
		{
			"test 1 pop empty stack",
			func(s *Stack) {
				s.Pop()
				s.Push(1)
				s.Pop()
				s.Pop()
				s.Push(5)
			},
			[]interface{}{5},
		},
	}

	for _, test := range stackTestData {
		var s Stack
		t.Run(test.description, func(t *testing.T) {
			test.operation(&s)
			if !reflect.DeepEqual(s, test.expected) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Was expecting '%v' but result was '%v'",
					test.expected, s)
			}
		})
	}
}
