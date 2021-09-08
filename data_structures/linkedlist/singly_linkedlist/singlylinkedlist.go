// Package singlelinkedlist Single Linked List
package singly_linkedlist

// demonstration of singly linked list in golang
import (
	"errors"
	"fmt"
)

// TODO: This is not how it should be done but I don't see a way out of this
type snode struct {
	Val  interface{}
	Next *snode
}

// SingleLinkedList structure with length of the list and its head
type SingleLinkedList struct {
	length int
	Head   *snode
}

// CreateList returns a new instance of a linked list
func CreateList() *SingleLinkedList {
	return &SingleLinkedList{}
}

// to avoid mistakes when using pointer vs struct for new snode creation
func NewNode(val interface{}) *snode {
	return &snode{val, nil}
}

// AddAtBeg adds a new snode with given value at the beginning of the list.
func (ll *SingleLinkedList) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

// AddAtEnd adds a new snode with given value at the end of the list.
func (ll *SingleLinkedList) AddAtEnd(val int) {
	n := NewNode(val)

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

// DelAtBeg deletes the snode at the head(beginning) of the list and returns its value. Returns -1 if the list is empty.
func (ll *SingleLinkedList) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.Val
}

// DelAtEnd deletes the snode at the tail(end) of the list and returns its value. Returns -1 if the list is empty.
func (ll *SingleLinkedList) DelAtEnd() interface{} {
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
func (ll *SingleLinkedList) Count() int {
	return ll.length
}

// Reverse reverses the list.
func (ll *SingleLinkedList) Reverse() {
	var prev, Next *snode
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = prev
		prev = cur
		cur = Next
	}

	ll.Head = prev
}

// ReversePartition Reverse the linked list from the ath to the bth node
func (ll *SingleLinkedList) ReversePartition(left, right int) error {
	err := ll.CheckRangeFromIndex(left, right)
	if err != nil {
		return err
	}
	tmpNode := NewNode(-1)
	tmpNode.Next = ll.Head
	pre := tmpNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	ll.Head = tmpNode.Next
	return nil
}

func (ll *SingleLinkedList) CheckRangeFromIndex(left, right int) error {
	if left > right {
		return errors.New("left boundary must smaller than right")
	} else if left < 1 {
		return errors.New("left boundary starts from the first node")
	} else if right > ll.length {
		return errors.New("right boundary cannot be greater than the length of the linked list")
	}
	return nil
}

// Display prints out the elements of the list.
func (ll *SingleLinkedList) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}
