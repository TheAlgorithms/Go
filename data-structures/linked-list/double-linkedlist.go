// demonstration of doubly linked list in golang

package doublelinkedlist

// package main

import "fmt"

type node struct {
	val  int
	next *node
	prev *node
}

type doublelinkedlist struct {
	head *node
}

// to avoid mistakes when using pointer vs struct for new node creation
func newNode(val int) *node {
	n := &node{}
	n.val = val
	n.next = nil
	n.prev = nil
	return n
}

func (ll *doublelinkedlist) addAtBeg(val int) {
	n := newNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *doublelinkedlist) addAtEnd(val int) {
	n := newNode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
	n.prev = cur
}

func (ll *doublelinkedlist) delAtBeg() int {
	if ll.head == nil {
		return -1
	}

	cur := ll.head
	ll.head = cur.next

	if ll.head != nil {
		ll.head.prev = nil
	}

	return cur.val
}

func (ll *doublelinkedlist) delAtEnd() int {
	// no item
	if ll.head == nil {
		return -1
	}

	// only one item
	if ll.head.next == nil {
		return ll.delAtBeg()
	}

	// more than one, go to second last
	cur := ll.head
	for ; cur.next.next != nil; cur = cur.next {
	}

	retval := cur.next.val
	cur.next = nil
	return retval
}

func (ll *doublelinkedlist) count() int {
	var ctr int = 0

	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}

	return ctr
}

func (ll *doublelinkedlist) reverse() {
	var prev, next *node
	cur := ll.head

	for cur != nil {
		next = cur.next
		cur.next = prev
		cur.prev = next
		prev = cur
		cur = next
	}

	ll.head = prev
}

func (ll *doublelinkedlist) display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

func (ll *doublelinkedlist) displayReverse() {
	if ll.head == nil {
		return
	}
	var cur *node
	for cur = ll.head; cur.next != nil; cur = cur.next {
	}

	for ; cur != nil; cur = cur.prev {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

/*
func main() {
	ll := doublelinkedlist{}

	ll.addAtBeg(10)
	ll.addAtEnd(20)
	ll.display()
	ll.addAtBeg(30)
	ll.display()

	ll.reverse()
	ll.display()
	ll.displayReverse()

	fmt.Print(ll.delAtBeg(), "\n")
	fmt.Print(ll.delAtEnd(), "\n")
	fmt.Print("Display")
	ll.display()
	fmt.Print(ll.delAtBeg(), "\n")
	ll.display()
	fmt.Print(ll.delAtBeg(), "\n")
	ll.display()
}
*/
