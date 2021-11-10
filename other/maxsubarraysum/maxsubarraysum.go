/* O(n) solution, for calculating
maximum contiguous sum in the given array. */

// Package maxsubarraysum is a package containing a solution to a common
// problem of finding max contiguous sum within a array of ints.
package maxsubarraysum

import (
	"github.com/TheAlgorithms/Go/math/max"
)

// MaxSubarraySum returns the maximum subarray sum
func MaxSubarraySum(array []int) int {
	var currentMax int
	var maxTillNow int
	if len(array) != 0 {
		currentMax = array[0]
		maxTillNow = array[0]
	}
	for _, v := range array {
		currentMax = max.Int(v, currentMax+v)
		maxTillNow = max.Int(maxTillNow, currentMax)
	}
	return maxTillNow
}
