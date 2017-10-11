// demonstration of singly linked list in golang

package singlelinkedlist

// package main

import "fmt"

type node struct {
	val  int
	next *node
}

type singlelinkedlist struct {
	head *node
}

// to avoid mistakes when using pointer vs struct for new node creation
func newNode(val int) *node {
	return &node{val, nil}
}

func (ll *singlelinkedlist) addAtBeg(val int) {
	n := newNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *singlelinkedlist) addAtEnd(val int) {
	n := newNode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
}

func (ll *singlelinkedlist) delAtBeg() int {
	if ll.head == nil {
		return -1
	}

	cur := ll.head
	ll.head = cur.next

	return cur.val
}

func (ll *singlelinkedlist) delAtEnd() int {
	if ll.head == nil {
		return -1
	}

	if ll.head.next == nil {
		return ll.delAtBeg()
	}

	cur := ll.head

	for ; cur.next.next != nil; cur = cur.next {
	}

	retval := cur.next.val
	cur.next = nil
	return retval

}

func (ll *singlelinkedlist) count() int {
	var ctr int = 0

	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}

	return ctr
}

func (ll *singlelinkedlist) reverse() {
	var prev, next *node
	cur := ll.head

	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	ll.head = prev
}

func (ll *singlelinkedlist) display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

// func main() {
// 	ll := singlelinkedlist{}

// 	ll.addAtBeg(10)
// 	ll.addAtEnd(20)
// 	ll.display()
// 	ll.addAtBeg(30)
// 	ll.display()
// 	ll.reverse()
// 	ll.display()

// 	fmt.Print(ll.delAtBeg(), "\n")
// 	ll.display()

// 	fmt.Print(ll.delAtEnd(), "\n")
// 	ll.display()

// }
