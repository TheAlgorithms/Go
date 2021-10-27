package sort

type MaxHeap struct {
	slice    []Comparable
	heapSize int
	indexs   map[int]int
}

func buildMaxHeap(slice0 []int) MaxHeap {
	var slice []Comparable
	for _, i := range slice0 {
		slice = append(slice, Int(i))
	}
	h := MaxHeap{}
	h.Init(slice)
	return h
}

func (h *MaxHeap) Init(slice []Comparable) {
	if slice == nil {
		slice = make([]Comparable, 0)
	}

	h.slice = slice
	h.heapSize = len(slice)
	h.indexs = make(map[int]int)
	h.Heapify()
}

func (h MaxHeap) Size() int {
	return h.heapSize
}

func (h *MaxHeap) Heapify() {
	for i, v := range h.slice {
		h.indexs[v.Idx()] = i
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

func (h MaxHeap) Update(i Comparable) {
	h.slice[h.indexs[i.Idx()]] = i
	h.heapifyUp(h.indexs[i.Idx()])
	h.heapifyDown(h.indexs[i.Idx()])
}

func (h MaxHeap) updateidx(i int) {
	h.indexs[h.slice[i].Idx()] = i
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
	More(interface{}) bool
}
type Int int

func (a Int) More(b interface{}) bool {
	return a > b.(Int)
}
func (a Int) Idx() int {
	return int(a)
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
