package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
  
  // Floyd’s Cycle-Finding Algorithm 
  
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}

	}
	return false
}

func main() {
	head := ListNode{1,nil}
	head.Next = &ListNode{2,nil}
	head.Next.Next = &ListNode{3,nil}
	head.Next.Next.Next = &ListNode{4,nil}
	
	// {1} -> {2} -> {3} -> {4} -> nil
	fmt.Println(hasCycle(&head)) // false
	
	head.Next.Next.Next.Next = head.Next.Next // loop
	
	
	// {1} -> {2} -> {3} -> {4}
	//         ^-------------'
	fmt.Println(hasCycle(&head)) // true
}
