package heap

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	// insert three elements into the heap, then remove them
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	newHeap(&pq)
	if pq.Len() != 3 {
		fmt.Println("expected 3 items in the heap but got", pq.Len())
	}
	item := Pop(&pq).(*Item)
	if item.value != "pear" {
		fmt.Println("expected pear but got", item.value)
	}
	if item.priority != 4 {
		fmt.Println("expected 4 but got", item.priority)
	}

	Push(&pq, item)

	// Insert a new item and then modify its priority.
	item = &Item{
		value:    "orange",
		priority: 1,
	}
	Push(&pq, item)
	pq.update(item, item.value, 5)

	if pq.Len() != 4 {
		fmt.Println("expected 4 items in the heap but got", pq.Len())
	}
	item = Pop(&pq).(*Item)
	if item.value != "orange" {
		fmt.Println("expected orange but got", item.value)
	}
	if item.priority != 5 {
		fmt.Println("expected 5 but got", item.priority)
	}
	item = Pop(&pq).(*Item)
	if item.value != "pear" {
		fmt.Println("expected pear but got", item.value)
	}
	if item.priority != 4 {
		fmt.Println("expected 4 but got", item.priority)
	}
	item = Pop(&pq).(*Item)
	if item.value != "banana" {
		fmt.Println("expected banana but got", item.value)
	}
	if item.priority != 3 {
		fmt.Println("expected 3 but got", item.priority)
	}
	item = Pop(&pq).(*Item)
	if item.value != "apple" {
		fmt.Println("expected apple but got", item.value)
	}
	if item.priority != 2 {
		fmt.Println("expected 2 but got", item.priority)
	}
	if pq.Len() != 0 {
		fmt.Println("expected 0 items in the heap but got", pq.Len())
	}
}
