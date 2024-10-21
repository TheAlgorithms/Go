// heapsort.go
// description: Implementation of heap sort algorithm
// worst-case time complexity: O(n log n)
// average-case time complexity: O(n log n)
// space complexity: O(1)

package sort

import "github.com/TheAlgorithms/Go/constraints"

type MaxHeap struct {
	slice    []Comparable
	heapSize int
	indices  map[int]int
}

func (h *MaxHeap) Init(slice []Comparable) {
	if slice == nil {
		slice = make([]Comparable, 0)
	}

	h.slice = slice
	h.heapSize = len(slice)
	h.indices = make(map[int]int)
	h.Heapify()
}

func (h MaxHeap) Heapify() {
	for i, v := range h.slice {
		h.indices[v.Idx()] = i
	}
	for i := h.heapSize / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MaxHeap) Pop() Comparable {
	if h.heapSize == 0 {
		return nil
	}

	i := h.slice[0]
	h.heapSize--

	h.slice[0] = h.slice[h.heapSize]
	h.updateidx(0)
	h.heapifyDown(0)

	h.slice = h.slice[0:h.heapSize]
	return i
}

func (h *MaxHeap) Push(i Comparable) {
	h.slice = append(h.slice, i)
	h.updateidx(h.heapSize)
	h.heapifyUp(h.heapSize)
	h.heapSize++
}

func (h MaxHeap) Size() int {
	return h.heapSize
}

func (h MaxHeap) Update(i Comparable) {
	h.slice[h.indices[i.Idx()]] = i
	h.heapifyUp(h.indices[i.Idx()])
	h.heapifyDown(h.indices[i.Idx()])
}

func (h MaxHeap) updateidx(i int) {
	h.indices[h.slice[i].Idx()] = i
}

func (h *MaxHeap) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
	h.updateidx(i)
	h.updateidx(j)
}

func (h MaxHeap) more(i, j int) bool {
	return h.slice[i].More(h.slice[j])
}

func (h MaxHeap) heapifyUp(i int) {
	if i == 0 {
		return
	}
	p := i / 2

	if h.slice[i].More(h.slice[p]) {
		h.swap(i, p)
		h.heapifyUp(p)
	}
}

func (h MaxHeap) heapifyDown(i int) {
	heapifyDown(h.slice, h.heapSize, i, h.more, h.swap)
}

func heapifyDown[T any](slice []T, N, i int, moreFunc func(i, j int) bool, swapFunc func(i, j int)) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < N && moreFunc(l, max) {
		max = l
	}
	if r < N && moreFunc(r, max) {
		max = r
	}
	if max != i {
		swapFunc(i, max)

		heapifyDown(slice, N, max, moreFunc, swapFunc)
	}
}

type Comparable interface {
	Idx() int
	More(any) bool
}

func HeapSort[T constraints.Ordered](slice []T) []T {
	N := len(slice)

	moreFunc := func(i, j int) bool {
		return slice[i] > slice[j]
	}
	swapFunc := func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	}

	// build a maxheap
	for i := N/2 - 1; i >= 0; i-- {
		heapifyDown(slice, N, i, moreFunc, swapFunc)
	}

	for i := N - 1; i > 0; i-- {
		slice[i], slice[0] = slice[0], slice[i]
		heapifyDown(slice, i, 0, moreFunc, swapFunc)
	}

	return slice
}
