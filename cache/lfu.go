// lfu.go
// description: a type of cache algorithm used to manage memory within a computer.
// details:
// The standard characteristics of this method involve the system keeping track of the number of times a block is referenced in memory.
// When the cache is full and requires more room the system will purge the item with the lowest reference frequency.
// ref: (https://en.wikipedia.org/wiki/Least_frequently_used)
// time complexity: O(N)
// space complexity: O(1)
// author: [CocaineCong](https://github.com/CocaineCong)

package cache

import (
	"container/list"
	"math"
)

// LFU the Least Frequently Used (LFU) page-replacement algorithm
type LFU struct {
	len     int // length
	cap     int // capacity
	minFreq int // The element that operates least frequently in LFU

	// key: key of element, value: value of element
	itemMap map[string]*list.Element

	// key: frequency of possible occurrences of all elements in the itemMap
	// value: elements with the same frequency
	freqMap map[int]*list.List
}

// NewLFU init the LFU cache with capacity
func NewLFU(capacity int) LFU {
	return LFU{
		len:     0,
		cap:     capacity,
		minFreq: math.MaxInt,
		itemMap: make(map[string]*list.Element),
		freqMap: make(map[int]*list.List),
	}
}

// initItem to init item for LFU
func initItem(k string, v any, f int) item {
	return item{
		key:   k,
		value: v,
		freq:  f,
	}
}

// Get the key in cache by LFU
func (c *LFU) Get(key string) any {
	// if existed, will return value
	if e, ok := c.itemMap[key]; ok {
		// the frequency of e +1 and change freqMap
		c.increaseFreq(e)
		obj := e.Value.(item)
		return obj.value
	}

	// if not existed, return nil
	return nil
}

// Put the key in LFU cache
func (c *LFU) Put(key string, value any) {
	if e, ok := c.itemMap[key]; ok {
		// if key existed, update the value
		obj := e.Value.(item)
		obj.value = value
		c.increaseFreq(e)
	} else {
		// if key not existed
		obj := initItem(key, value, 1)
		// if the length of item gets to the top line
		// remove the least frequently operated element
		if c.len == c.cap {
			c.eliminate()
			c.len--
		}
		// insert in freqMap and itemMap
		c.insertMap(obj)
		// change minFreq to 1 because insert the newest one
		c.minFreq = 1
		// length++
		c.len++
	}
}

// increaseFreq increase the frequency if element
func (c *LFU) increaseFreq(e *list.Element) {
	obj := e.Value.(item)
	// remove from low frequency first
	oldLost := c.freqMap[obj.freq]
	oldLost.Remove(e)
	// change the value of minFreq
	if c.minFreq == obj.freq && oldLost.Len() == 0 {
		// if it is the last node of the minimum frequency that is removed
		c.minFreq++
	}
	// add to high frequency list
	c.insertMap(obj)
}

// insertMap insert item in map
func (c *LFU) insertMap(obj item) {
	// add in freqMap
	l, ok := c.freqMap[obj.freq]
	if !ok {
		l = list.New()
		c.freqMap[obj.freq] = l
	}
	e := l.PushFront(obj)
	// update or add the value of itemMap key to e
	c.itemMap[obj.key] = e
}

// eliminate clear the least frequently operated element
func (c *LFU) eliminate() {
	l := c.freqMap[c.minFreq]
	e := l.Back()
	obj := e.Value.(item)
	l.Remove(e)

	delete(c.itemMap, obj.key)
}
