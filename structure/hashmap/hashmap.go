package hashmap

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity uint64 = 1 << 10

type node struct {
	key   any
	value any
	next  *node
}

// HashMap is a Golang implementation of a hashmap
type HashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

// DefaultNew returns a new HashMap instance with default values
func DefaultNew() *HashMap {
	return &HashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity),
	}
}

// New creates a new HashMap instance with the specified size and capacity
func New(size, capacity uint64) *HashMap {
	return &HashMap{
		size:     size,
		capacity: capacity,
		table:    make([]*node, capacity),
	}
}

// Get returns the value associated with the given key
func (hm *HashMap) Get(key any) any {
	node := hm.getNodeByKey(key)
	if node != nil {
		return node.value
	}
	return nil
}

// Put inserts a new key-value pair into the hashmap
func (hm *HashMap) Put(key, value any) {
	index := hm.hash(key)
	if hm.table[index] == nil {
		hm.table[index] = &node{key: key, value: value}
	} else {
		current := hm.table[index]
		for {
			if current.key == key {
				current.value = value
				return
			}
			if current.next == nil {
				break
			}
			current = current.next
		}
		current.next = &node{key: key, value: value}
	}
	hm.size++
	if float64(hm.size)/float64(hm.capacity) > 0.75 {
		hm.resize()
	}
}

// Contains checks if the given key is stored in the hashmap
func (hm *HashMap) Contains(key any) bool {
	return hm.getNodeByKey(key) != nil
}

// getNodeByKey finds the node associated with the given key
func (hm *HashMap) getNodeByKey(key any) *node {
	index := hm.hash(key)
	current := hm.table[index]
	for current != nil {
		if current.key == key {
			return current
		}
		current = current.next
	}
	return nil
}

// resize doubles the capacity of the hashmap and rehashes all existing entries
func (hm *HashMap) resize() {
	oldTable := hm.table
	hm.capacity <<= 1
	hm.table = make([]*node, hm.capacity)
	hm.size = 0

	for _, head := range oldTable {
		for current := head; current != nil; current = current.next {
			hm.Put(current.key, current.value)
		}
	}
}

// hash generates a hash value for the given key
func (hm *HashMap) hash(key any) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := h.Sum64()
	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}
