// demonstration of singly linked list in golang

package linkedlist 

import "fmt"

type singlenode struct {
	val  int
	next *singlenode
}

type singlelinkedlist struct {
	head *singlenode
}

// to avoid mistakes when using pointer vs struct for new singlenode creation
func newSingleNode(val int) *singlenode {
	return &singlenode{val, nil}
}

func (ll *singlelinkedlist) addFirst(val int) {
	n := newSingleNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *singlelinkedlist) addLast(val int) {
	n := newSingleNode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
}

func (ll *singlelinkedlist) removeFirst() int {
	if ll.head == nil {
		return -1
	}

	cur := ll.head
	ll.head = cur.next

	return cur.val
}

func (ll *singlelinkedlist) removeLast() int {
	if ll.head == nil {
		return -1
	}

	if ll.head.next == nil {
		return ll.removeFirst()
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
	var prev, next *singlenode
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

/*
func main() {
 	ll := singlelinkedlist{}

 	ll.addFirst(10)
 	ll.addLast(20)
	ll.display()

	fmt.Print("Adding item '30' at the beginning...\n") 
 	ll.addFirst(30)
	ll.display()
	
	fmt.Print("Reversing linkedlist...\n")
 	ll.reverse()
 	ll.display()

 	fmt.Print("Removing '", ll.removeFirst(), "'\n")
 	ll.display()

 	fmt.Print("Removing '", ll.removeLast(), "'\n")
 	ll.display()
}
*/