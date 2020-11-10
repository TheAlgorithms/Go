package singlelinkedlist

// demonstration of singly linked list in golang
import "fmt"

type node struct {
	Val  interface{}
	Next *node
}

type singlelinkedlist struct {
	length int
	Head   *node
}

// CreateList returns a new instance of a linked list
func CreateList() *singlelinkedlist {
	return &singlelinkedlist{}
}

// to avoid mistakes when using pointer vs struct for new node creation
func newNode(val interface{}) *node {
	return &node{val, nil}
}

// AddAtBeg adds a new node with given value at the beginning of the list.
func (ll *singlelinkedlist) AddAtBeg(val interface{}) {
	n := newNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

// AddAtEnd adds a new node with given value at the end of the list.
func (ll *singlelinkedlist) AddAtEnd(val int) {
	n := newNode(val)

	if ll.Head == nil {
		ll.Head = n
		ll.length++
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = n
	ll.length++
}

// DelAtBeg deletes the node at the head(beginning) of the list and returns its value. Returns -1 if the list is empty.
func (ll *singlelinkedlist) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.Val
}

// DelAtEnd deletes the node at the tail(end) of the list and returns its value. Returns -1 if the list is empty.
func (ll *singlelinkedlist) DelAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}

	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head

	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	cur.Next = nil
	ll.length--
	return retval

}

// Count returns the current size of the list.
func (ll *singlelinkedlist) Count() int {
	return ll.length
}

// Reverse reverses the list.
func (ll *singlelinkedlist) Reverse() {
	var prev, Next *node
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = prev
		prev = cur
		cur = Next
	}

	ll.Head = prev
}

// Display prints out the elements of the list.
func (ll *singlelinkedlist) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}
