/* Coin Change problem
1. Finding the total number of combination that can be achieved with the given number of coins to get the specified sum.
*/

package main

import "fmt"

func coinChangeCount(coins []int, sum int) int {
	if sum < 1 {
		return 0
	}

	var arr = make([]int, sum+1)

	arr[0] = 1
	
	for _, coinVal := range coins {
		for j := coinVal; j <= sum; j++ {
			arr[j] += arr[j-coinVal]
		}
	}

	return arr[sum]
}

func main() {
        fmt.Println(coinChangeCount([]int {1,2}, 4))
}
