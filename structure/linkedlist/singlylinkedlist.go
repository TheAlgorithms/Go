package linkedlist

// demonstration of singly linked list in golang
import "fmt"

// Singly structure with length of the list and its head
type Singly struct {
	length int

	// Note that Node here holds both Next and Prev Node
	// however only the Next node is used in Singly methods.
	Head *Node
}

// NewSingly returns a new instance of a linked list
func NewSingly() *Singly {
	return &Singly{}
}

// AddAtBeg adds a new snode with given value at the beginning of the list.
func (ll *Singly) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

// AddAtEnd adds a new snode with given value at the end of the list.
func (ll *Singly) AddAtEnd(val int) {
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
func (ll *Singly) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.Val
}

// DelAtEnd deletes the snode at the tail(end) of the list and returns its value. Returns -1 if the list is empty.
func (ll *Singly) DelAtEnd() interface{} {
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
func (ll *Singly) Count() int {
	return ll.length
}

// Reverse reverses the list.
func (ll *Singly) Reverse() {
	var prev, Next *Node
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
func (ll *Singly) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}
