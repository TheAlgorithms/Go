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

func (ll *singlelinkedlist) AddAtBeg(val interface{}) {
	n := newNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

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

func (ll *singlelinkedlist) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.Val
}

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

func (ll *singlelinkedlist) Count() int {
	return ll.length
}

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

func (ll *singlelinkedlist) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
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
