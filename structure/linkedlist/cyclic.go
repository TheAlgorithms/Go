package linkedlist

import "fmt"

// Cyclic Struct which cycles the linked list in this implementation.
type Cyclic[T any] struct {
	Size int
	Head *Node[T]
}

// Create new list.
func NewCyclic[T any]() *Cyclic[T] {
	return &Cyclic[T]{}
}

// Inserting the first node is a special case. It will
// point to itself. For other cases, the node will be added
// to the end of the list. End of the list is Prev field of
// current item. Complexity O(1).
func (cl *Cyclic[T]) Add(val T) {
	n := NewNode(val)
	cl.Size++
	if cl.Head == nil {
		n.Prev = n
		n.Next = n
		cl.Head = n
	} else {
		n.Prev = cl.Head.Prev
		n.Next = cl.Head
		cl.Head.Prev.Next = n
		cl.Head.Prev = n
	}
}

// Rotate list by P places.
// This method is interesting for optimization.
// For first optimization we must decrease
// P value so that it ranges from 0 to N-1.
// For this we need to use the operation of
// division modulo. But be careful if P is less than 0.
// if it is - make it positive. This can be done without
// violating the meaning of the number by adding to it
// a multiple of N. Now you can decrease P modulo N to
// rotate the list by the minimum number of places.
// We use the fact that moving forward in a circle by P
// places is the same as moving N - P places back.
// Therefore, if P > N / 2, you can turn the list by N-P places back.
// Complexity O(n).
func (cl *Cyclic[T]) Rotate(places int) {
	if cl.Size > 0 {
		if places < 0 {
			multiple := cl.Size - 1 - places/cl.Size
			places += multiple * cl.Size
		}
		places %= cl.Size

		if places > cl.Size/2 {
			places = cl.Size - places
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Prev
			}
		} else if places == 0 {
			return
		} else {
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Next
			}

		}
	}
}

// Delete the current item.
func (cl *Cyclic[T]) Delete() bool {
	var deleted bool
	var prevItem, thisItem, nextItem *Node[T]

	if cl.Size == 0 {
		return deleted
	}

	deleted = true
	thisItem = cl.Head
	nextItem = thisItem.Next
	prevItem = thisItem.Prev

	if cl.Size == 1 {
		cl.Head = nil
	} else {
		cl.Head = nextItem
		nextItem.Prev = prevItem
		prevItem.Next = nextItem
	}
	cl.Size--

	return deleted
}

// Destroy all items in the list.
func (cl *Cyclic[T]) Destroy() {
	for cl.Delete() {
		continue
	}
}

// Show list body.
func (cl *Cyclic[T]) Walk() *Node[T] {
	var start *Node[T]
	start = cl.Head

	for i := 0; i < cl.Size; i++ {
		fmt.Printf("%v \n", start.Val)
		start = start.Next
	}
	return start
}

// https://en.wikipedia.org/wiki/Josephus_problem
// This is a struct-based solution for Josephus problem.
func JosephusProblem(cl *Cyclic[int], k int) int {
	for cl.Size > 1 {
		cl.Rotate(k)
		cl.Delete()
		cl.Rotate(-1)
	}
	retval := cl.Head.Val
	cl.Destroy()
	return retval
}
