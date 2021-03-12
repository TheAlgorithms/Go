// Package heap : max heap implemnation
// https://en.wikipedia.org/wiki/Heap_(data_structure)
package heap

// Heap represents a max heap
type Heap struct {
	bucket []int
	Size   int
}

// New : create a new max heap with given capacity
func New(capacity int) *Heap {
	return &Heap{
		bucket: make([]int, capacity),
		Size:   0,
	}
}

// Heapify : convert a given array into a valid max heap
func Heapify(arr []int) *Heap {
	h := Heap{
		bucket: arr,
		Size:   len(arr),
	}
	for i := (len(arr) / 2) - 1; i >= 0; i-- {
		h.shiftDown(i)
	}
	return &h
}

// Insert : will insert given element into heap
func (h *Heap) Insert(elm int) {
	h.bucket[h.Size] = elm
	h.shiftUp(h.Size)
	h.Size++
}

// Max : returns max element from the heap
func (h *Heap) Max() int {
	return h.bucket[0]
}

// ExtractMax : remove max element from the heap and return its value
func (h *Heap) ExtractMax() int {
	out := h.bucket[0]
	h.Size--
	h.bucket[0] = h.bucket[h.Size]
	h.shiftDown(0)
	return out
}

func (h *Heap) shiftDown(i int) {
	for i < h.Size {
		/// tmp : index of max child
		var tmp int
		if 2*i+2 < h.Size {
			tmp = 2*i + 1
			if h.bucket[2*i+2] > h.bucket[2*i+1] {
				tmp = 2*i + 2
			}
		} else if 2*i+1 < h.Size {
			tmp = 2*i + 1
		} else {
			break
		}
		if h.bucket[tmp] > h.bucket[i] {
			h.bucket[i], h.bucket[tmp] = h.bucket[tmp], h.bucket[i]
			i = tmp
		} else {
			break
		}
	}
}

func (h *Heap) shiftUp(i int) {
	for i > 0 && h.bucket[i] > h.bucket[(i-1)/2] {
		h.bucket[i], h.bucket[(i-1)/2] = h.bucket[(i-1)/2], h.bucket[i]
		i = (i - 1) / 2
	}
}
