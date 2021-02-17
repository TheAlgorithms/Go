// demonstration of doubly linked list in golang

package doublelinkedlist

// package main

import "fmt"

// TODO: Not the ways its supposed to be done but can't think of a work around for tests to not worry about redeclaration
type dnode struct {
	val  int
	next *dnode
	prev *dnode
}

// DoubleLinkedList structure with just the head node
type DoubleLinkedList struct {
	head *dnode
}

// to avoid mistakes when using pointer vs struct for new node creation
func newNode(val int) *dnode {
	n := &dnode{}
	n.val = val
	n.next = nil
	n.prev = nil
	return n
}

func (ll *DoubleLinkedList) addAtBeg(val int) {
	n := newNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *DoubleLinkedList) addAtEnd(val int) {
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

func (ll *DoubleLinkedList) delAtBeg() int {
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

func (ll *DoubleLinkedList) delAtEnd() int {
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

func (ll *DoubleLinkedList) count() int {
	var ctr int = 0

	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}

	return ctr
}

func (ll *DoubleLinkedList) reverse() {
	var prev, next *dnode
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

func (ll *DoubleLinkedList) display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

func (ll *DoubleLinkedList) displayReverse() {
	if ll.head == nil {
		return
	}
	var cur *dnode
	for cur = ll.head; cur.next != nil; cur = cur.next {
	}

	for ; cur != nil; cur = cur.prev {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

/*
func main() {
	ll := DoubleLinkedList{}

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
