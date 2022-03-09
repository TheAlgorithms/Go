package math

import (
	"github.com/TheAlgorithms/Go/search"
	"github.com/TheAlgorithms/Go/sort"
)

// FindKthMax use quick sort algorithm to get kth large element in given array
// return -1 when not found
func FindKthMax(nums []int, k int) (int, error) {
	index := len(nums) - k
	return kthNumber(nums, index)
}

// FindKthMin use quick sort algorithm to get kth small element in given array
// return -1 when not found
func FindKthMin(nums []int, k int) (int, error) {
	index := k - 1
	return kthNumber(nums, index)
}

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
