// coinchange.go
// description: Implementation of the coin change problem using dynamic programming
// details: The coin change problem is a problem that asks for the number of ways to make change for a given amount of money using a given set of coins. The problem can be solved using dynamic programming.
// time complexity: O(n*m) where n is the number of coins and m is the amount of money
// space complexity: O(m) where m is the amount of money

package dynamic

// CoinChange finds the number of possible combinations of coins
// of different values which can get to the target amount.
func CoinChange(coins []int32, amount int32) int32 {
	combination := make([]int32, amount)
	combination[0] = 1

	for _, c := range coins {
		for i := c; i < amount; i++ {

			combination[i] += combination[i-c]
		}
	}

	return combination[amount-1]
}
