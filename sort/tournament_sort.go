// tournament_sort.go
// Package sort provides Tournament Sort implementation.
//
// Tournament Sort improves upon naive Selection Sort by using a min-heap
// to find the minimum element in O(log n) instead of O(n),
// giving an overall time complexity of O(n log n).
//
// Reference: https://en.wikipedia.org/wiki/Tournament_sort

package sort

import "container/heap"

// intHeap implements heap.Interface for a min-heap of ints.
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Tournament sorts a slice of integers using Tournament Sort.
// Time complexity: O(n log n), Space complexity: O(n)
func Tournament(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	h := &intHeap{}
	heap.Init(h)
	for _, v := range arr {
		heap.Push(h, v)
	}

	result := make([]int, len(arr))
	for i := range result {
		result[i] = heap.Pop(h).(int)
	}
	return result
}
