package linkedlist

// demonstration of singly linked list in golang
import (
	"errors"
	"fmt"
)

// Singly structure with length of the list and its head
type Singly[T any] struct {
	length int

	// Note that Node here holds both Next and Prev Node
	// however only the Next node is used in Singly methods.
	Head *Node[T]
}

// NewSingly returns a new instance of a linked list
func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{}
}

// AddAtBeg adds a new snode with given value at the beginning of the list.
func (ll *Singly[T]) AddAtBeg(val T) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

// AddAtEnd adds a new snode with given value at the end of the list.
func (ll *Singly[T]) AddAtEnd(val T) {
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

// DelAtBeg deletes the snode at the head(beginning) of the list
// and returns its value. Returns false if the list is empty.
func (ll *Singly[T]) DelAtBeg() (T, bool) {
	if ll.Head == nil {
		var r T
		return r, false
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.Val, true
}

// DelAtEnd deletes the snode at the tail(end) of the list
// and returns its value. Returns false if the list is empty.
func (ll *Singly[T]) DelAtEnd() (T, bool) {
	if ll.Head == nil {
		var r T
		return r, false
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
	return retval, true

}

// DelByPos deletes the node at the middle based on position in the list
// and returns its value. Returns false if the list is empty or length is not more than given position
func (ll *Singly[T]) DelByPos(pos int) (T, bool) {
	switch {
	case ll.Head == nil:
		var r T
		return r, false
	case pos-1 > ll.length:
		var r T
		return r, false
	case pos-1 == 0:
		return ll.DelAtBeg()
	case pos-1 == ll.Count():
		return ll.DelAtEnd()
	}

	var prev *Node[T]
	var val T
	cur := ll.Head
	count := 0

	for count < pos-1 {
		prev = cur
		cur = cur.Next
		count++
	}

	val = cur.Val
	prev.Next = cur.Next
	ll.length--

	return val, true
}

// Count returns the current size of the list.
func (ll *Singly[T]) Count() int {
	return ll.length
}

// Reverse reverses the list.
func (ll *Singly[T]) Reverse() {
	var prev, next *Node[T]
	cur := ll.Head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	ll.Head = prev
}

// ReversePartition Reverse the linked list from the ath to the bth node
func (ll *Singly[T]) ReversePartition(left, right int) error {
	err := ll.CheckRangeFromIndex(left, right)
	if err != nil {
		return err
	}
	tmpNode := &Node[T]{}
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
func (ll *Singly[T]) CheckRangeFromIndex(left, right int) error {
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
func (ll *Singly[T]) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}
