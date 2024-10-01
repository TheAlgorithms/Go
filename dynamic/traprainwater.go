// filename: traprainwater.go
// description: Provides a function to calculate the amount of trapped rainwater between bars represented by an elevation map using dynamic programming.
// details:
// The TrapRainWater function calculates the amount of trapped rainwater between the bars represented by the given elevation map.
// It uses dynamic programming to precompute the maximum height of bars to the left and right of each position.
// Then, it iterates through the array to calculate the amount of trapped rainwater at each position based on the minimum of the left and right maximum heights.
// Finally, it sums up the trapped rainwater for all positions and returns the total amount.
// time complexity: O(n)
// space complexity: O(n)
// author(s) [TruongNhanNguyen (SOZEL)](https://github.com/TruongNhanNguyen)
package dynamic

import "math"

// TrapRainWater calculates the amount of trapped rainwater between the bars represented by the given elevation map.
// It uses dynamic programming to precompute the maximum height of bars to the left and right of each position.
// Then, it iterates through the array to calculate the amount of trapped rainwater at each position based on the minimum of the left and right maximum heights.
// Finally, it sums up the trapped rainwater for all positions and returns the total amount.
func TrapRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(height[i])))
	}

	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = int(math.Max(float64(rightMax[i+1]), float64(height[i])))
	}

	trappedWater := 0
	for i := 0; i < len(height); i++ {
		trappedWater += int(math.Min(float64(leftMax[i]), float64(rightMax[i]))) - height[i]
	}

	return trappedWater
}
