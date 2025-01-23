// partitionproblem.go
// description: Solves the Partition Problem using dynamic programming
// reference: https://en.wikipedia.org/wiki/Partition_problem
// time complexity: O(n*sum)
// space complexity: O(n*sum)

package dynamic

// PartitionProblem checks whether the given set can be partitioned into two subsets
// such that the sum of the elements in both subsets is the same.
func PartitionProblem(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[target]
}
