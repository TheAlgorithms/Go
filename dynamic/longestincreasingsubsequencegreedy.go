package dynamic

// LongestIncreasingSubsequenceGreedy is a function to find the longest increasing
// subsequence in a given array using a greedy approach.
// The dynamic programming approach is implemented alongside this one.
// Worst Case Time Complexity: O(nlogn)
// Auxiliary Space: O(n), where n is the length of the array(slice).
// Reference: https://www.geeksforgeeks.org/construction-of-longest-monotonically-increasing-subsequence-n-log-n/
func LongestIncreasingSubsequenceGreedy(nums []int) int {
	longestIncreasingSubsequence := make([]int, 0)

	for _, num := range nums {
		// find the leftmost index in longestIncreasingSubsequence with value >= num
		leftmostIndex := lowerBound(longestIncreasingSubsequence, num)

		if leftmostIndex == len(longestIncreasingSubsequence) {
			longestIncreasingSubsequence = append(longestIncreasingSubsequence, num)
		} else {
			longestIncreasingSubsequence[leftmostIndex] = num
		}
	}

	return len(longestIncreasingSubsequence)
}

// Function to find the leftmost index in arr with value >= val, mimicking the inbuild lower_bound function in C++
// Time Complexity: O(logn)
// Auxiliary Space: O(1)
func lowerBound(arr []int, val int) int {
	searchWindowLeft, searchWindowRight := 0, len(arr)-1

	for searchWindowLeft <= searchWindowRight {
		middle := (searchWindowLeft + searchWindowRight) / 2

		if arr[middle] < val {
			searchWindowLeft = middle + 1
		} else {
			searchWindowRight = middle - 1
		}
	}

	return searchWindowRight + 1
}
