/* Coin Change problem
1. Finding the total number of combination that can be achived with the given number of coins to get the specified sum.
*/

package main

import "fmt"

func coin_change_count(coins []int, sum int) int {
	if sum < 1 {
		return 0
	}
	var arr = make([]int, sum+1)
	arr[0] = 1
	for _, coin_val := range coins {
		for j := coin_val; j <= sum; j++ {
			arr[j] += arr[j-coin_val]
		}
	}
	return arr[sum]
}

func main() {
        fmt.Println(coin_change_count([]int {1,2}, 4)) // Expected output will be 3
}
