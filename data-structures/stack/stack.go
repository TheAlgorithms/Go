// Package stack implements the Stack data structure, LIFO.
package stack

// Stack is the type that implements the Stack data structure.
type Stack struct {
	items []interface{}
}

// Push pushes an item onto the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Top returns the item last inserted into the stack.
func (s *Stack) Top() interface{} {
	if !s.IsEmpty() {
		return s.items[len(s.items)-1]
	}
	return nil
}

// Pop removes the last item from the stack.
func (s *Stack) Pop() {
	if !s.IsEmpty() {
		s.items = s.items[:len(s.items)-1]
	}
}

// IsEmpty returns a bool denoting whether or not the stack
// is empty or not.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
