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

// HashMap is golang implementation of hashmap
type HashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

// New return new HashMap instance
func New() *HashMap {
	return &HashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity),
	}
}

// Make creates a new HashMap instance with input size and capacity
func Make(size, capacity uint64) HashMap {
	return HashMap{
		size:     size,
		capacity: capacity,
		table:    make([]*node, capacity),
	}
}

// Get returns value associated with given key
func (hm *HashMap) Get(key any) any {
	node := hm.getNodeByHash(hm.hash(key))
	if node == nil {
		return nil
	}

	if node.key == key {
		return node.value
	}

	next := node.next
	for next != nil {
		if next.key == key {
			return next.value
		}
		next = next.next
	}
	return nil
}

// Put puts new key value in hashmap
func (hm *HashMap) Put(key any, value any) any {
	return hm.putValue(hm.hash(key), key, value)
}

// Contains checks if given key is stored in hashmap
func (hm *HashMap) Contains(key any) bool {
	node := hm.Get(key)
	return node != nil
}

func (hm *HashMap) putValue(hash uint64, key any, value any) any {
	if hm.capacity == 0 {
		hm.capacity = defaultCapacity
		hm.table = make([]*node, defaultCapacity)
	}

	node := hm.getNodeByHash(hash)

	if node == nil {
		hm.table[hash] = newNode(key, value)

	} else if node.key == key {
		hm.table[hash] = newNodeWithNext(key, value, node)
		return value

	} else {
		hm.resize()
		return hm.putValue(hash, key, value)
	}

	hm.size++

	return value

}

func (hm *HashMap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}

func (hm *HashMap) resize() {
	hm.capacity <<= 1

	tempTable := hm.table

	hm.table = make([]*node, hm.capacity)

	for i := 0; i < len(tempTable); i++ {
		node := tempTable[i]
		if node == nil {
			continue
		}

		hm.table[hm.hash(node.key)] = node
	}
}

func newNode(key any, value any) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func newNodeWithNext(key any, value any, next *node) *node {
	return &node{
		key:   key,
		value: value,
		next:  next,
	}
}

func (hm *HashMap) hash(key any) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}
