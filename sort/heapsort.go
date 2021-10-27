package sort

type MaxHeap struct {
	slice    []Comparable
	heapSize int
}

func buildMaxHeap(slice0 []int) MaxHeap {
	var slice []Comparable
	for _, i := range slice0 {
		slice = append(slice, Int(i))
	}
	h := MaxHeap{slice: slice, heapSize: len(slice)}
	h.Heapify()
	return h
}

func (h MaxHeap) Heapify() {
	for i := h.heapSize / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h MaxHeap) heapifyDown(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.heapSize && h.slice[l].more(h.slice[max]) {
		max = l
	}
	if r < h.heapSize && h.slice[r].more(h.slice[max]) {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.heapifyDown(max)
	}
}

func (h MaxHeap) heapifyUp(i int) {
	if i == 0 {
		return
	}
	p := i / 2

	if h.slice[i].more(h.slice[p]) {
		h.slice[i], h.slice[p] = h.slice[p], h.slice[i]
		h.heapifyUp(p)
	}
}

func (h MaxHeap) Pop() Comparable {
	if h.heapSize == 0 {
		return nil
	}

	i := h.slice[0]
	h.heapSize--

	h.slice[0] = h.slice[h.heapSize]
	h.heapifyDown(0)

	h.slice = h.slice[0:h.heapSize]
	return i
}

func (h MaxHeap) Push(i Comparable) {
	h.slice = append(h.slice, i)
	h.heapifyUp(h.heapSize)
	h.heapSize++
}

type Comparable interface {
	more(interface{}) bool
}
type Int int

func (a Int) more(b interface{}) bool {
	return a > b.(Int)
}

func HeapSort(slice []int) []int {
	h := buildMaxHeap(slice)
	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.heapifyDown(0)
	}

	res := []int{}
	for _, i := range h.slice {
		res = append(res, int(i.(Int)))
	}
	return res
}
