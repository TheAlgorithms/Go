package hash_map

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity uint64 = 1 << 10

type node struct {
	key   interface{}
	value interface{}
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

// Get returns value associated with given key
func (hm *HashMap) Get(key interface{}) interface{} {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.value
	}

	return nil
}

// Put puts new key value in hashmap
func (hm *HashMap) Put(key interface{}, value interface{}) interface{} {
	return hm.putValue(hm.hash(key), key, value)
}

// Contains checks if given key is stored in hashmap
func (hm *HashMap) Contains(key interface{}) bool {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return true
	}

	return false
}

func (hm *HashMap) putValue(hash uint64, key interface{}, value interface{}) interface{} {
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

func newNode(key interface{}, value interface{}) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func newNodeWithNext(key interface{}, value interface{}, next *node) *node {
	return &node{
		key:   key,
		value: value,
		next:  next,
	}
}

func (hm *HashMap) hash(key interface{}) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

// func main() {
// 	HashMap := New()

// 	HashMap.Put("test-1", 10)
// 	fmt.Println(HashMap.Get("test-1"))

// 	HashMap.Put("test-1", 20)
// 	HashMap.Put("test-2", 30)
// 	HashMap.Put(1, 40)

// 	fmt.Println(HashMap.Get("test-1"))
// 	fmt.Println(HashMap.Get("test-2"))
// 	fmt.Println(HashMap.Get(1))

// 	fmt.Println(HashMap.Contains(2))
// 	fmt.Println(HashMap.Contains(1))
// 	fmt.Println(HashMap.Contains("test-1"))
// }
