package heap

/*
This is an implementation of a Binary heap (or simply heap), is parameterized with:
T: The type of elements
I: The type of the cardinality of the binary heap as a vector

P.s. If you want use a primitive type you have to wrap it into a custom type and then implement the Prioritizable interface

Usage example:

	type MyInt32 int32

	func (a MyInt32) PriorTo(b MyInt32) bool {
		return a < b
	}
	func main() {
		h := heap.NewBinaryHeapFromSlice[MyInt32, int32]([]MyInt32{4, 2})
		h.Push(3)
		fmt.Println(h.Pop())
	}
*/

import (
	"github.com/TheAlgorithms/Go/constraints"
)

type BinaryHeap[T any, I constraints.Unsigned] struct {
	s []T
	prior func(T, T) bool 
}

// Note for function Prior: It's a strict order relation
func NewBinaryHeapFromSlice[T any, I constraints.Unsigned](
	s []T,
	Prior func(T, T) bool) (h *BinaryHeap[T, I]) {
	h = &BinaryHeap[T, I]{
		s: s,
		prior: Prior,
	}
	if h.Len() > 1 {
		i := (h.Len() - 2) / 2
		for {
			h.heapifyDown(i)
			if i == 0 {
				break
			}
			i--
		}
	}
	return
}
func NewBinaryHeap[T any, I constraints.Unsigned](Prior func(T, T) bool) *BinaryHeap[T, I] {
	return &BinaryHeap[T, I]{s: make([]T, 0), prior: Prior}
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
			if h.prior(h.s[j-1], h.s[j]) {
				j--
			}
		} else {
			j--
			if j >= h.Len() {
				break
			}
		}
		if h.prior(h.s[j], h.s[index]) {
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
		if index == 0 {
			break
		}
		parent := (index - 1) / 2
		if h.prior(h.s[parent], h.s[index]) {
			break
		}
		h.s[index], h.s[parent] = h.s[parent], h.s[index]
		index = parent
	}
}
func (h *BinaryHeap[T, I]) Preview() T {
	return h.s[0]
}
func (h BinaryHeap[T, I]) String() string {
	return "" // #TODO
}