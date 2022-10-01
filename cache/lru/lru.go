package lru

import "container/list"

type item struct {
	key   string
	value any
}

type LRU struct {
	dl          *list.List
	capacity    int
	maxCapacity int
	storage     map[string]*list.Element
}

// New lru cache with capacity
func New(capacity int) LRU {
	return LRU{
		dl:          list.New(),
		storage:     make(map[string]*list.Element, capacity),
		capacity:    0,
		maxCapacity: capacity,
	}
}

// Get value from lru
// if not found, return nil
func (c *LRU) Get(key string) any {
	v, ok := c.storage[key]
	if ok {
		c.dl.MoveToBack(v)
		return v.Value.(item).value
	}

	return nil
}

// Put cache with key and value to lru
func (c *LRU) Put(key string, value any) {
	e, ok := c.storage[key]
	if ok {
		n := e.Value.(item)
		n.value = value
		e.Value = n
		c.dl.MoveToBack(e)
		return
	}

	if c.capacity >= c.maxCapacity {
		e := c.dl.Front()
		dk := e.Value.(item).key
		c.dl.Remove(e)
		delete(c.storage, dk)
		c.capacity--
	}

	n := item{key: key, value: value}
	ne := c.dl.PushBack(n)
	c.storage[key] = ne
	c.capacity++
}
