// Stack Linked-List
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stacklinkedlistwithlist.go, stackarray.go, stack_test.go

package stack

// Node structure
type Node struct {
	Val  any
	Next *Node
}

// Stack has jost top of node and with length
type Stack struct {
	top    *Node
	length int
}

// push add value to last index
func (ll *Stack) Push(n any) {
	newStack := &Node{} // new node

	newStack.Val = n
	newStack.Next = ll.top

	ll.top = newStack
	ll.length++
}

// pop remove last item as first output
func (ll *Stack) Pop() any {
	result := ll.top.Val
	if ll.top.Next == nil {
		ll.top = nil
	} else {
		ll.top.Val, ll.top.Next = ll.top.Next.Val, ll.top.Next.Next
	}

	ll.length--
	return result
}

// isEmpty to check our array is empty or not
func (ll *Stack) IsEmpty() bool {
	return ll.length == 0
}

// len use to return length of our stack
func (ll *Stack) Length() int {
	return ll.length
}

// peak return last input value
func (ll *Stack) Peek() any {
	return ll.top.Val
}

// show all value as an interface array
func (ll *Stack) Show() (in []any) {
	current := ll.top

	for current != nil {
		in = append(in, current.Val)
		current = current.Next
	}
	return
}
