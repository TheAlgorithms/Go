package heap

/*
This is an implementation of a Binary heap (or simply heap), is parameterized with:
T: The type of elements
I: The type of the cardinality of the binary heap as a vector

P.s. If you want use a primitive type you have to wrap it into a custom type and then implement the Prioritizable interface

Usage example:

	type MyInt32 int32

	func (a MyInt32) PriorTo(o MyInt32) bool {
		return a < o
	}
	func main() {
		h := heap.NewBinaryHeap[MyInt32, int32]()
		h.Push(1)
		fmt.Println(h.Pop())
	}

*/

import (
	"github.com/TheAlgorithms/Go/constraints"
)

type Prioritizable[T any] interface {
	PriorTo(x T) bool // If you take this as a relation: if (a,b) in R => (b,a) not in R
}

type BinaryHeap[T Prioritizable[T], I constraints.Signed] struct {
	s []T
}

func NewBinaryHeapFromSlice[T Prioritizable[T], I constraints.Signed](s []T) (h *BinaryHeap[T, I]) {
	h = &BinaryHeap[T, I]{
		s: s,
	}
	for i := (h.Len() - 2) / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
	return
}
func NewBinaryHeap[T Prioritizable[T], I constraints.Signed]() *BinaryHeap[T, I] {
	return &BinaryHeap[T, I]{s: make([]T, 0)}
}
func (h *BinaryHeap[T, I]) Len() I {
	return I(len(h.s))
}
func (h *BinaryHeap[T, I]) Push(value T) {
	h.s = append(h.s, value)
	h.heapifyUp(h.Len() - 1)
}
func (h *BinaryHeap[T, I]) Pop() (res T) {
	res = h.s[0]
	h.s[0] = h.s[h.Len()-1]
	h.s = h.s[:h.Len()-1]
	h.heapifyDown(0)
	return
}

func (h *BinaryHeap[T, I]) heapifyDown(index I) bool {
	origin := index
	for {
		j := index*2 + 2
		if j < h.Len() {
			if h.s[j-1].PriorTo(h.s[j]) {
				j--
			}
		} else {
			if j <= h.Len() && !h.s[j].PriorTo(h.s[index]) {
				h.s[j], h.s[index] = h.s[index], h.s[j]
			}
			break
		}
		if h.s[j].PriorTo(h.s[index]) {
			h.s[j], h.s[index] = h.s[index], h.s[j]
			index = j
		} else {
			break
		}
	}
	return origin != index
}
func (h *BinaryHeap[T, I]) heapifyUp(index I) {
	for {
		parent := (index - 1) / 2
		if parent == index || h.s[parent].PriorTo(h.s[index]) {
			break
		}
		h.s[index], h.s[parent] = h.s[parent], h.s[index]
		index = parent
	}
}
func (h *BinaryHeap[T, I]) Preview() T {
	return h.s[0]
}
