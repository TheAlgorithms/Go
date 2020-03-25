package stack

// Stack : holds down the stack elements
type Stack struct {
	stack []string
}

// Unshift : add an element to the stack(front)
func (s *Stack) Unshift(element string) {
	newStack := append([]string{element}, s.stack...)
	s.stack = newStack
}

// Shift : remove element form the stack(front)
func (s *Stack) Shift() {
	newStack := s.stack[1:]
	s.stack = newStack
}

// Peek : helper function to view elements in the stack
func (s Stack) Peek() []string {
	return s.stack
}

// Size : helper function to return the size of the stack
func (s Stack) Size() int {
	return len(s.stack)
}

// Empty : helper function th return bool if stack is empty
func (s Stack) Empty() bool {
	return s.Size() == 0
}
