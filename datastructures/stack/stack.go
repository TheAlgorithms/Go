// Package stack an abstract data type
package stack

/*Stack is an abstract data type that serves as a collection of elements.
	The order in which elements come off a stack gives rise to its alternative name, LIFO (last in, first out).
reference: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
*/
type Stack []interface{}

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(element interface{}) {
	*s = append(*s, element) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element, true
}
