//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func bubbleSort(nums []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1] < nums[i] {
				nums[i+1], nums[i] = nums[i], nums[i+1]
				swapped = true
			}
		}
	}
	return nums
}
