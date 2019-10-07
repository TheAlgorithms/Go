// Given a value N, if we want to make change for N cents, and we have infinite supply of each of S = { S1, S2, .. , Sm} valued coins
// how many ways can we make the change?
// https://www.geeksforgeeks.org/coin-change-dp-7/

//package main

package coinChange

//for storing the value
var memoization [4][100000]int

//item price
var price int

//coin type
var coins []int

func calculateWays(index int, currentPrice int) int {

	if index >= len(coins) || currentPrice < 0 {
		return 0
	}

	//return memoization value if already exists
	if memoization[index][currentPrice] != 0 {
		return memoization[index][currentPrice]
	}

	if currentPrice == 0 {
		return 1
	}

	//storing the value before returning
	memoization[index][currentPrice] = calculateWays(index, currentPrice-coins[index]) + calculateWays(index+1, currentPrice-coins[index])

	return memoization[index][currentPrice]
}

/*
	Using normal recursive is possible to solve this
	but will hugely affect performance as shown below (for result where price is 2000)

	Recursive took 33.744592936s
	Dynamic Programming took 206.323785ms
*/

/*
func main() {

	// add coins
	coins = append(coins, 1, 2, 5, 10)

	//price
	price = 2000

	fmt.Println(calculateWays(0, price))
}*/
