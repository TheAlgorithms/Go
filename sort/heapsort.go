package sort

import "github.com/TheAlgorithms/Go/constraints"

type MaxHeap[T Comparable] struct {
	slice    []T
	heapSize int
	indices  map[int]int
}

func (h *MaxHeap[T]) Init(slice []T) {
	if slice == nil {
		slice = make([]T, 0)
	}

	h.slice = slice
	h.heapSize = len(slice)
	h.indices = make(map[int]int)
	h.Heapify()
}

func (h MaxHeap[T]) Heapify() {
	for i, v := range h.slice {
		h.indices[v.Idx()] = i
	}
	for i := h.heapSize / 2; i >= 0; i-- {
		heapifyDown(h.slice, h.Size(), i, h.more, h.swap)
	}
}

func (h *MaxHeap[T]) Pop() T {
	var zero T
	if h.heapSize == 0 {
		return zero
	}

	i := h.slice[0]
	h.heapSize--

	h.slice[0] = h.slice[h.heapSize]
	h.updateidx(0)
	heapifyDown(h.slice, h.Size(), 0, h.more, h.swap)

	h.slice = h.slice[0:h.heapSize]
	return i
}

func (h *MaxHeap[T]) Push(i T) {
	h.slice = append(h.slice, i)
	h.updateidx(h.heapSize)
	h.heapifyUp(h.heapSize)
	h.heapSize++
}

func (h MaxHeap[T]) Size() int {
	return h.heapSize
}

func (h MaxHeap[T]) Update(i T) {
	h.slice[h.indices[i.Idx()]] = i
	h.heapifyUp(h.indices[i.Idx()])
	heapifyDown(h.slice, h.Size(), i.Idx(), h.more, h.swap)
}

func (h MaxHeap[T]) updateidx(i int) {
	h.indices[h.slice[i].Idx()] = i
}

func (h *MaxHeap[T]) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
	h.updateidx(i)
	h.updateidx(j)
}

func (h MaxHeap[T]) more(i, j int) bool {
	return h.slice[i].More(h.slice[j])
}

func (h MaxHeap[T]) heapifyUp(i int) {
	if i == 0 {
		return
	}
	p := i / 2

	if h.slice[i].More(h.slice[p]) {
		h.swap(i, p)
		h.heapifyUp(p)
	}
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
