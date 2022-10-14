// Heap data structure
// description: based on `geeksforgeeks` heap is a specialized tree-based data structure which is essentially an almost complete tree that satisfies the heap property.
//
//	In a heap, the highest (or lowest) priority element is always stored at the root.
//	Heap data structure has the following two main operations:
//		- Insert: Inserts an element into the heap.
//		- Extract Max: Removes the root element from the heap.
//
// details:
//
//	Heap Data Structure : https://www.geeksforgeeks.org/heap-data-structure/
//	Heap (data structure) : https://en.wikipedia.org/wiki/Heap_(data_structure)
package heap

type SortInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// HeapInterface is an interface that must be implemented by types that are used as a heap.
type HeapInterface interface {
	SortInterface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

// newHeap turns a slice of any into a heap, in place.
// The complexity is O(n) where n = h.Len().
func newHeap(h HeapInterface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h HeapInterface, x any) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h HeapInterface) any {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove(h HeapInterface, i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix modifies the heap h with the changed element at index i.
// The complexity is O(log n) where n = h.Len().
func Fix(h HeapInterface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

// Up modifies the heap with the changed element at index i (parent = (j - 1) / 2 )
func up(h HeapInterface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// Down modifies the heap with the changed element at index i (left child = 2*i + 1 and right child = 2*i + 2 )
func down(h HeapInterface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
