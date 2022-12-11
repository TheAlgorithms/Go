package heap

import (
	"fmt"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	var binaryHeapTestData = []struct {
		array       []int
	}{
		{
			array:       []int{1},
		},
		{
			array:       []int{1, 2},
		},
		{
			array:       []int{2, 1},
		},
		{
			array:       []int{3, 2, 1},
		},
		{
			array:       []int{3, 1, 2},
		},
		{
			array:       []int{2, 3, 1, 5, 4},
		},
		{
			array:       []int{11, 9, 6, 2, 1, 3, 4, 7, 8, 5, 10},
		},
		{
			array:       []int{10, 10, 6, 1, 1, 3, 3, 8, 8, 6, 10},
		},
	}
	for _, test := range binaryHeapTestData {
		t.Run(fmt.Sprint("Testing ", test.array), func(t *testing.T) {
			fromSlice := NewBinaryHeapFromSlice[int, uint64](test.array,
				func(a, b int) bool { return a < b },
			)
			notFromSlice := NewBinaryHeap[int, uint64](func(a, b int) bool { return a < b })
			for i := 0; i < len(test.array); i++ {
				notFromSlice.Push(test.array[i])
			}
			if notFromSlice.Len() != fromSlice.Len() {
				t.Fatalf("Length is different while pushing elements")
			}
			check := func(h *BinaryHeap[int, uint64]) {
				last := 0
				for h.Len() > 0 {
					next := h.Pop()
					if next < last {
						t.Fatalf(fmt.Sprint("Popped wrong element"))
					}
					last = next
				}
			}
			check(fromSlice)
			check(notFromSlice)
		})
	}
}
