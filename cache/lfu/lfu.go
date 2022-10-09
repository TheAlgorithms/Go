// Package lfu (Least Frequently Used) is a type of cache algorithm used to manage memory within a computer.
// ref: https://en.m.wikipedia.org/wiki/Least_frequently_used

package lfu

import (
	"container/heap"
	"errors"
)

type LFUCache struct {
	bucket    map[string]*Item
	capacity  int
	pq        PriorityQueue
	timeStamp int
}

type Item struct {
	counter   int
	key       string
	value     any
	index     int
	timeStamp int
}

// NewLFU initializes the object with the capacity of the data structure.
func NewLFU(capacity int) LFUCache {
	priorityQueue := make(PriorityQueue, 0, capacity)
	heap.Init(&priorityQueue)
	return LFUCache{
		bucket:   make(map[string]*Item, capacity),
		capacity: capacity,
		pq:       priorityQueue,
	}
}

// Get gets the value of the key if the key exists in the cache. Otherwise, throws an error
func (c *LFUCache) Get(key string) (any, error) {
	item, exists := c.bucket[key]
	if !exists {
		return nil, errors.New("key does not exist")
	}

	item.counter++
	c.timeStamp++
	item.timeStamp = c.timeStamp
	// we need to update the item in the PQ, so it can be reordered
	c.pq.Update(item)

	return item.value, nil
}

// Put updates the value of the key if present,
// or inserts the key if not already present.
// When the cache reaches its capacity,
// it should invalidate and remove the least
// frequently used key before inserting a new item.
// In this implementation, when there is a tie
// (i.e., two or more keys with the same frequency),
// the least recently used key would be invalidated.
func (c *LFUCache) Put(key string, value any) {
	if c.capacity == 0 {
		return
	}
	c.timeStamp++
	item, exists := c.bucket[key]
	if !exists {
		item = &Item{
			counter:   1,
			value:     value,
			timeStamp: c.timeStamp,
			key:       key,
		}
		//We need to make more room for the new item
		if len(c.bucket) == c.capacity {
			//Remove the least frequently used or the least recently used item
			removedItem := heap.Pop(&c.pq).(*Item)
			delete(c.bucket, removedItem.key)
		}
		c.bucket[key] = item
		heap.Push(&c.pq, item)
		return
	}

	item.counter++
	item.value = value
	item.timeStamp = c.timeStamp
	c.pq.Update(item)
}

type PriorityQueue []*Item

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less returns true if the priority of the first item is less than the second.
func (pq PriorityQueue) Less(i, j int) bool {

	//it is a tie, so we need to use the least recently used rule
	if pq[i].counter == pq[j].counter {
		return pq[i].timeStamp < pq[j].timeStamp
	}
	// We want Pop to give us the lowest, not highest.
	return pq[i].counter < pq[j].counter
}

// Swap swaps the elements with indexes i and j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push pushes the element x onto the heap.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes the minimum element (according to Less) from the heap and returns it.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item) {
	heap.Fix(pq, item.index)
}
