// Coin Change Problem
// https://www.geeksforgeeks.org/understanding-the-coin-change-problem-with-dynamic-programming/

package coinChange

func coinChangeIterative(total int, coins []int) int {
	if len(coins) == 0 || coins[0] != 0 {
		coins = append([]int{0}, coins...)
	}

	memo := make([][]int, len(coins))
	for i := range memo {
		memo[i] = make([]int, total+1)
	}

	memo[0][0] = 1
	for i := 1; i < len(memo); i += 1 {
		coin := coins[i]
		for j := range memo[i] {
			if j == 0 {
				memo[i][j] = 1
			} else if j == coin {
				memo[i][j] = 1 + memo[i-1][j]
			} else if j > coin {
				memo[i][j] = memo[i-1][j] + memo[i][j-coin]
			} else {
				memo[i][j] = memo[i-1][j]
			}
		}
	}

	return memo[len(coins)-1][total]
}

func solveRecursive(total int, coins []int, memo *[][]int, coin int) int {
	if total == 0 {
		return 1
	}

	if coin < 0 {
		return 0
	}

	if (*memo)[coin][total] != -1 {
		return (*memo)[coin][total]
	}

	curAnswer := 0
	for i := coin; i >= 0; i -= 1 {
		if total >= coins[i] {
			curAnswer += solveRecursive(total-coins[i], coins, memo, i)
		}
	}

	(*memo)[coin][total] = curAnswer
	return curAnswer
}

func coinChangeRecursive(total int, coins []int) int {
	memo := make([][]int, len(coins))
	for i := range memo {
		memo[i] = make([]int, total+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return solveRecursive(total, coins, &memo, len(coins)-1)
}
