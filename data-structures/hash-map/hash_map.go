package data-structures

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

type hashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

func newHashMap() *hashMap {
	return &hashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity),
	}
}

func (hm *hashMap) get(key interface{}) interface{} {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.value
	}

	return nil
}

func (hm *hashMap) put(key interface{}, value interface{}) interface{} {
	return hm.putValue(hm.hash(key), key, value)
}

func (hm *hashMap) putValue(hash uint64, key interface{}, value interface{}) interface{} {
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

func (hm *hashMap) contains(key interface{}) bool {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return true
	}

	return false
}

func (hm *hashMap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}

func (hm *hashMap) resize() {
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

func (hm *hashMap) hash(key interface{}) uint64 {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

// func main() {
// 	hashMap := newHashMap()

// 	hashMap.put("test-1", 10)
// 	fmt.Println(hashMap.get("test-1"))

// 	hashMap.put("test-1", 20)
// 	hashMap.put("test-2", 30)
// 	hashMap.put(1, 40)

// 	fmt.Println(hashMap.get("test-1"))
// 	fmt.Println(hashMap.get("test-2"))
// 	fmt.Println(hashMap.get(1))

// 	fmt.Println(hashMap.contains(2))
// 	fmt.Println(hashMap.contains(1))
// 	fmt.Println(hashMap.contains("test-1"))
// }
