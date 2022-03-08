package kth

// FindKthMax use quick sort algorithm to get kth large element in given array
// return -1 when not found
func FindKthMax(nums []int, k int) int {
	index := len(nums) - k
	return kthNumber(nums, index)
}

// FindKthMin use quick sort algorithm to get kth small element in given array
// return -1 when not found
func FindKthMin(nums []int, k int) int {
	index := k - 1
	return kthNumber(nums, index)
}

func kthNumber(nums []int, k int) int {
	if k < 0 || k >= len(nums) {
		return -1
	}
	start := 0
	end := len(nums)
	for start < end {
		pivot := partition(nums, start, end)
		if k == pivot {
			return nums[pivot]
		} else if k > pivot {
			start = pivot + 1
		} else {
			end = pivot
		}
	}
	return -1
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	j := start
	for i := start + 1; i < end; i++ {
		if nums[i] < pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[start], nums[j] = nums[j], nums[start]
	return j
}
