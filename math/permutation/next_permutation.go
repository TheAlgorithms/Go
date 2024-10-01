// A practice to find lexicographically next greater permutation of the given array of integers.
// If there does not exist any greater permutation, then print the lexicographically smallest permutation of the given array.
// The implementation below, finds the next permutation in linear time and constant memory and returns in place
// time complexity: O(n)
// space complexity: O(1)
// Useful reference: https://www.geeksforgeeks.org/next-permutation/

package permutation

func NextPermutation(nums []int) {
	pivot := 0
	for pivot = len(nums) - 2; pivot >= 0; pivot-- {
		if nums[pivot] < nums[pivot+1] {
			break
		}
	}
	if pivot < 0 {
		// current permutation is the last and must be reversed totally
		for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
			nums[l], nums[r] = nums[r], nums[l]
		}
	} else {
		succ := 0
		for succ = len(nums) - 1; succ > pivot; succ = succ - 1 {
			if nums[succ] > nums[pivot] {
				break
			}
		}

		// Swap the pivot and successor
		nums[pivot], nums[succ] = nums[succ], nums[pivot]

		// Reverse the suffix part to minimize it
		for l, r := pivot+1, len(nums)-1; l < r; l, r = l+1, r-1 {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

}
