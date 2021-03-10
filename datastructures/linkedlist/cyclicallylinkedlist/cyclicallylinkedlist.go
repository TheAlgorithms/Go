// Package cyclicallylinkedlist demonstration of Linked List Cycle in golang

package cyclicallylinkedlist

import "fmt"

// Node of List.
type ClNode struct {
	val  interface{}
	prev *ClNode
	next *ClNode
}

// The Cycling list in this implementation is addressed
// by means of the element located at the current position.
type ClList struct {
	size        int
	currentItem *ClNode
}

// Create new list.
func NewList() *ClList {
	return &ClList{0, nil}
}

// Create new node.
func NewNode(val interface{}) *ClNode {
	return &ClNode{val, nil, nil}
}

// Inserting the first node is a special case. It will
// point to itself. For other cases, the node will be added
// to the end of the list. End of the list is prev field of
// current item. Complexity O(1).
func (cl *ClList) Add(val interface{}) {
	n := NewNode(val)
	cl.size++
	if cl.currentItem == nil {
		n.prev = n
		n.next = n
		cl.currentItem = n
	} else {
		n.prev = cl.currentItem.prev
		n.next = cl.currentItem
		cl.currentItem.prev.next = n
		cl.currentItem.prev = n
	}
}

// Rotate method is interesting for optimization.
// Let us need to rotate the list by P places
// and list have N items. For first optimization
// we must decrease P value so that is in range from
// 0 to N-1. For this we need to use the operation of
// division modulo. But be careful if P is less than 0.
// if it is - make it positive. This can be done without
// violating the meaning of the number by adding to it
// a multiple of N. Now you can decrease P modulo N to
// rotate the list by the minimum number of places.
// We use the fact that moving forward in a circle by P
// places is the same as moving N - P places back.
// Therefore, if P > N / 2, you can turn the list by N-P places back.
// Complexity O(n).
func (cl *ClList) Rotate(places int) {
	if cl.size > 0 {
		if places < 0 {
			multiple := cl.size - 1 - places/cl.size
			places += multiple * cl.size
		}
		places %= cl.size

		if places > cl.size/2 {
			places = cl.size - places
			for i := 0; i < places; i++ {
				cl.currentItem = cl.currentItem.prev
			}
		} else if places == 0 {
			return
		} else {
			for i := 0; i < places; i++ {
				cl.currentItem = cl.currentItem.next
			}

		}
	}
}

// Delete the current item.
func (cl *ClList) Delete() bool {
	var deleted bool
	var prevItem, thisItem, nextItem *ClNode

	if cl.size == 0 {
		return deleted
	}

	deleted = true
	thisItem = cl.currentItem
	nextItem = thisItem.next
	prevItem = thisItem.prev

	if cl.size == 1 {
		cl.currentItem = nil
	} else {
		cl.currentItem = nextItem
		nextItem.prev = prevItem
		prevItem.next = nextItem
	}
	cl.size--

	return true
}

// Destroy all items in the list.
func (cl *ClList) Destroy() {
	for cl.Delete() == true {
		continue
	}
}

// Show list body.
func (cl *ClList) Walk() *ClNode {
	var start *ClNode
	start = cl.currentItem

	for i := 0; i < cl.size; i++ {
		fmt.Printf("%v \n", start.val)
		start = start.next
	}
	return start
}

// https://en.wikipedia.org/wiki/Josephus_problem
// This is a struct-based solution for Josephus problem.
func JosephusProblem(cl *ClList, k int) int {
	for cl.size > 1 {
		cl.Rotate(k)
		cl.Delete()
		cl.Rotate(-1)
	}
	retval := cl.currentItem.val.(int)
	cl.Destroy()
	return retval
}
