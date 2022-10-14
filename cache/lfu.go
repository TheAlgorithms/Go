// Package lfu (Least Frequently Used) is a type of cache algorithm used to manage memory within a computer.
// ref: https://en.m.wikipedia.org/wiki/Least_frequently_used
package cache

import (
	"fmt"

	"github.com/TheAlgorithms/Go/structure/linkedlist"
)

// LFUCache is a struct that holds the LFU cache
type LFUCache struct {
	lookUp       map[any]*linkedlist.Node[any]
	frequency    map[any]*linkedlist.Doubly[any]
	minFrequency int
	size         int
}

// LFUNode is a node for LFU cache
type LFUNode struct {
	key   any
	value any
	count int
}

// NewLFU initializes the object with the capacity of the data structure.
func NewLFU(capacity int) LFUCache {
	cache := LFUCache{}
	cache.lookUp = map[any]*linkedlist.Node[any]{}
	cache.frequency = map[any]*linkedlist.Doubly[any]{}
	cache.size = capacity
	cache.minFrequency = -1
	return cache
}

// Touch updates the frequency of the key and moves it to the front
// of the list of the new frequency bucket
func (cache *LFUCache) Touch(element *linkedlist.Node[any]) {
	key, val := element.Val.(LFUNode).key, element.Val.(LFUNode).value
	cnt := element.Val.(LFUNode).count
	node := LFUNode{key, val, cnt + 1}
	cache.frequency[cnt].Remove(element)
	if cache.minFrequency == cnt && cache.frequency[cnt].Count() == 0 {
		cache.minFrequency += 1
	}
	if _, found := cache.frequency[cnt+1]; !found {
		cache.frequency[cnt+1] = linkedlist.NewDoubly[any]()
	}
	cache.frequency[cnt+1].AddAtBeg(node)
	cache.lookUp[key] = cache.frequency[cnt+1].Front()
}

// Print prints the LFU cache
func (cache *LFUCache) Print() {
	fmt.Println("print map")
	for _, v := range cache.lookUp {
		node := v.Val.(LFUNode)
		fmt.Println(node.key, node.value, node.count)
	}
	fmt.Println("end")
}

// Get gets the value of the key if the key exists in the cache.
// Otherwise, throws an error
func (cache *LFUCache) Get(key any) (any, error) {
	if e, found := cache.lookUp[key]; found {
		cache.Touch(e)
		return e.Val.(LFUNode).value, nil
	} else {
		return nil, fmt.Errorf("key not found")
	}
}

// Put updates the value of the key if present,
// or inserts the key if not already present.
// When the cache reaches its capacity,
// it should invalidate and remove the least
// frequently used key before inserting a new item.
// In this implementation, when there is a tie
// (i.e., two or more keys with the same frequency),
// the least recently used key would be invalidated.
func (cache *LFUCache) Put(key any, value any) {
	if cache.size <= 0 {
		return
	}

	if e, found := cache.lookUp[key]; found {
		node := e.Val.(LFUNode)
		node.value = value
		e.Val = node
		cache.Touch(e)
	} else {
		if len(cache.lookUp) == cache.size {
			cache.Evict()
		}
		cnt := 1
		cache.minFrequency = 1
		node := LFUNode{key: key, value: value, count: cnt}
		if _, found := cache.frequency[cnt]; !found {
			cache.frequency[cnt] = linkedlist.NewDoubly[any]()
		}
		cache.frequency[cnt].AddAtBeg(node)
		cache.lookUp[key] = cache.frequency[cnt].Front()
	}
}

// Evict removes the least frequently used key
func (cache *LFUCache) Evict() {
	evictElement := cache.frequency[cache.minFrequency].Back()
	delete(cache.lookUp, evictElement.Val.(LFUNode).key)
	cache.frequency[cache.minFrequency].Remove(evictElement)
}
