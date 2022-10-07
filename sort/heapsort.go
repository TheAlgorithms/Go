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

func (h MaxHeap) heapifyUp(i int) {
	if i == 0 {
		return
	}
	p := i / 2

	if h.slice[i].More(h.slice[p]) {
		h.slice[i], h.slice[p] = h.slice[p], h.slice[i]
		h.updateidx(i)
		h.updateidx(p)
		h.heapifyUp(p)
	}
}

func (h MaxHeap) heapifyDown(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.heapSize && h.slice[l].More(h.slice[max]) {
		max = l
	}
	if r < h.heapSize && h.slice[r].More(h.slice[max]) {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.updateidx(i)
		h.updateidx(max)
		h.heapifyDown(max)
	}
}

type Comparable interface {
	Idx() int
	More(any) bool
}
type Int int

func (a Int) More(b any) bool {
	return a > b.(Int)
}
func (a Int) Idx() int {
	return int(a)
}

func heapify[T constraints.Ordered](slice []T, N, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < N && slice[largest] < slice[l] {
		largest = l
	}
	if r < N && slice[largest] < slice[r] {
		largest = r
	}

	// change root, if needed
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]

		// heapify the root
		heapify(slice, N, largest)
	}
}

func HeapSort[T constraints.Ordered](slice []T) []T {
	N := len(slice)

	// build a maxheap
	for i := N/2 - 1; i >= 0; i-- {
		heapify(slice, N, i)
	}

	for i := N - 1; i > 0; i-- {
		slice[i], slice[0] = slice[0], slice[i]
		heapify(slice, i, 0)
	}

	return slice
}
