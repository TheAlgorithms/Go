package dynamic

// CoinChange finds the number of possible combinations of coins
// of different values which can get to the target amount.
func coinChange(coins []int32, amount int32) int32 {
	combination := make([]int32, amount)
	combination[0] = 1

	for _, c := range coins { // iterate combination of coin
		for i := c; i < amount; i++ {
			// find previous value
			combination[i] += combination[i-c]
		}
	}

	return combination[amount-1]
}
