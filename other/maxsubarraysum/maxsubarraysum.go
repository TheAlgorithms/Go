/* O(n) solution, for calculating
maximum contiguous sum in the given array. */

package maxsubarraysum

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxSubarraySum(array []int) int {
	var currentMax int = 0
	var maxTillNow int = 0
	for _, v := range array {
		currentMax = Max(v, currentMax+v)
		maxTillNow = Max(maxTillNow, currentMax)
	}
	return maxTillNow
}

// func main() {
// 	array := []int{-2, -5, 6, 0, -2, 0, -3, 1, 0, 5, -6}
// 	fmt.Println("Maximum contiguous sum: ", maxSubarraySum(array))
// }
