/* O(n) solution, for calculating
maximum contiguous sum in the given array. */

// Package maxsubarraysum is a package containing a solution to a common
// problem of finding max contigious sum within a array of ints.
package maxsubarraysum

import "github.com/TheAlgorithms/Go/math/max"

// MaxSubarraySum returns the maximum subarray sum
func MaxSubarraySum(array []int) int {
	var currentMax int = 0
	var maxTillNow int = 0
	for _, v := range array {
		currentMax = max.Int(v, currentMax+v)
		maxTillNow = max.Int(maxTillNow, currentMax)
	}
	return maxTillNow
}

// TODO: make a test file with this
// func main() {
// 	array := []int{-2, -5, 6, 0, -2, 0, -3, 1, 0, 5, -6}
// 	fmt.Println("Maximum contiguous sum: ", maxSubarraySum(array))
// }
