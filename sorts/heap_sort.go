//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

type maxHeap struct {
	slice    []int
	heapSize int
}

func buildMaxHeap(slice []int) maxHeap {
	h := maxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h maxHeap) MaxHeapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.heapSize && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.heapSize && h.slice[r] > h.slice[max] {
		max = r
	}
	//log.Printf("MaxHeapify(%v): l,r=%v,%v; max=%v\t%v\n", i, l, r, max, h.slice)
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max)
	}
}

// HeapSort main
func HeapSort(slice []int) []int {
	h := buildMaxHeap(slice)
	//log.Println(slice)
	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.MaxHeapify(0)
		/*if i == len(h.slice)-1 || i == len(h.slice)-3 || i == len(h.slice)-5 {
			element := (i - len(h.slice)) * -1
			fmt.Println("Heap after removing ", element, " elements")
			fmt.Println(h.slice)

		}*/
	}
	return h.slice
}
