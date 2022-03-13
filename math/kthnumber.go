package math

import (
	"github.com/TheAlgorithms/Go/search"
	"github.com/TheAlgorithms/Go/sort"
)

// FindKthMax returns the kth large element given an integer slice
// with nil `error` if found and returns -1 with `error` `search.ErrNotFound`
// if not found. NOTE: The `nums` slice gets mutated in the process.
func FindKthMax(nums []int, k int) (int, error) {
	index := len(nums) - k
	return kthNumber(nums, index)
}

// FindKthMin returns kth small element given an integer slice
// with nil `error` if found and returns -1 with `error` `search.ErrNotFound`
// if not found. NOTE: The `nums` slice gets mutated in the process.
func FindKthMin(nums []int, k int) (int, error) {
	index := k - 1
	return kthNumber(nums, index)
}

// kthNumber use the selection algorithm (based on the partition method - the same one as used in quicksort).
func kthNumber(nums []int, k int) (int, error) {
	if k < 0 || k >= len(nums) {
		return -1, search.ErrNotFound
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		pivot := sort.Partition(nums, start, end)
		if k == pivot {
			return nums[pivot], nil
		}
		if k > pivot {
			start = pivot + 1
			continue
		}
		end = pivot - 1
	}
	return -1, search.ErrNotFound
}
