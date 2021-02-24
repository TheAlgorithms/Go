// Package hashmap provide hash table implementation with collsion resolved
// using openaddressing
package hashmap

import (
	"fmt"
	"math"
)

const (
	goldenRatio     float64 = 1.618033988749895
	radix                   = uint64('~' - ' ')
	loadFactor      float64 = 0.667
	initMValue              = 1
	deletedEntryKey         = "U+0000"
)

type entry struct {
	key   interface{}
	value interface{}
}
type table struct {
	buckets  []*entry
	size     uint64
	m        uint64
	capacity uint64 // cap = 2^m
}

// newTable is internal function for creating new hash table
func newTable() *table {
	return &table{
		buckets:  make([]*entry, 1<<initMValue),
		size:     0,
		m:        uint64(initMValue),
		capacity: uint64(1 << initMValue),
	}
}
func (t *table) Put(key interface{}, value interface{}) {
	if t.size > uint64(loadFactor*float64(t.capacity)) {
		t.resize(t.m + 1)
	}
	encoded := encode(key)
	hashed := t.hash(encoded)
	index := hashed
	var trail uint64 = 1
	for ; t.buckets[index] != nil && t.buckets[index].key != key; trail++ {
		index = (hashed + (trail * t.probe(encoded))) % t.capacity
	}
	if t.buckets[index] == nil {
		t.buckets[index] = &entry{
			key:   key,
			value: value,
		}
		t.size++
	} else {
		t.buckets[index].value = value
	}
}

func (t *table) Len() int {
	return int(t.size)
}
func (t *table) Cap() int {
	return int(t.capacity)
}

func (t *table) Get(key interface{}) interface{} {
	encoded := encode(key)
	hashed := t.hash(encoded)
	index := hashed
	var trail uint64 = 1
	for ; t.buckets[index] != nil && t.buckets[index].key != deletedEntryKey; trail++ {
		if t.buckets[index].key == key {
			return t.buckets[index].value
		}
		index = (hashed + (trail * t.probe(encoded))) % t.capacity
	}
	return nil
}
func (t *table) Delete(key interface{}) {
	encoded := encode(key)
	hashed := t.hash(encoded)
	index := hashed
	var trail uint64 = 1
	for ; t.buckets[index] != nil; trail++ {
		if t.buckets[index].key == key {
			t.buckets[index].key = deletedEntryKey
			t.buckets[index].value = nil
			t.size--
			break
		}
		index = (hashed + (trail * t.probe(encoded))) % t.capacity
	}
	if t.size <= t.capacity/4 {
		t.resize(t.m - 1)
	}
}
func (t *table) Contains(key interface{}) bool {
	return t.Get(key) != nil
}

func (t *table) resize(newM uint64) {
	t.m = newM
	t.capacity = 1 << newM
	var encoded uint64
	var hashed uint64
	var index uint64
	var trail uint64
	newBuckets := make([]*entry, int(t.capacity))
	for _, v := range t.buckets {
		if v != nil && v.key != deletedEntryKey {
			encoded = encode(v.key)
			hashed = t.hash(encoded)
			trail = 1
			index = hashed
			for ; newBuckets[index] != nil; trail++ {
				index = (hashed + (trail * t.probe(encoded))) % t.capacity
			}
			newBuckets[index] = v
		}
	}
	t.buckets = nil
	t.buckets = newBuckets
}

// linear probing
func (t *table) probe(encoded uint64) uint64 {
	return 7
}

// encode converts a key of any type to a uint64 number
func encode(key interface{}) uint64 {
	var out uint64
	for _, v := range fmt.Sprintf("%v", key) {
		out = uint64(v) + radix*out
	}
	return out
}

// Multiplication Based hashing
// M (table size) = 2^m
func (t *table) hash(encoded uint64) uint64 {
	// aK : is golden ratio multiplied with encode key
	aK := goldenRatio * float64(encoded)
	// hash based on Multiplication
	return uint64((aK - math.Floor(aK)) * float64(t.capacity))
}
