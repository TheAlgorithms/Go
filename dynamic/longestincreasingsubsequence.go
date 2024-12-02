// longestincreasingsubsequence.go
// description: Implementation of the Longest Increasing Subsequence using dynamic programming
// reference: https://en.wikipedia.org/wiki/Longest_increasing_subsequence
// time complexity: O(n^2)
// space complexity: O(n)

package dynamic

import (
	"github.com/TheAlgorithms/Go/math/max"
)

// LongestIncreasingSubsequence returns the longest increasing subsequence
// where all elements of the subsequence are sorted in increasing order
func LongestIncreasingSubsequence(elements []int) int {
	n := len(elements)
	lis := make([]int, n)
	for i := range lis {
		lis[i] = 1
	}
	for i := range lis {
		for j := 0; j < i; j++ {
			if elements[i] > elements[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}
	res := 0
	for _, value := range lis {
		res = max.Int(res, value)
	}
	return res
}
