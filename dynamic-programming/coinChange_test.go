package coinChange

import "testing"

type coinChangeFunction func(int, []int) int

func testLog(desired int, gotten int, t *testing.T) {
	if desired != gotten {
		t.Fatal("Got", gotten, ", wanted", desired)
	}
}

func testCoinChange(solver coinChangeFunction, t *testing.T) {
	testLog(5, solver(5, []int{1, 2, 3}), t)
	testLog(4, solver(4, []int{1, 2, 3}), t)
	testLog(3, solver(3, []int{1, 2, 3}), t)
	testLog(2, solver(2, []int{1, 2, 3}), t)
	testLog(1, solver(1, []int{1, 2, 3}), t)
	testLog(1, solver(0, []int{1, 2, 3}), t)
	testLog(3, solver(12, []int{3, 5, 6, 10}), t)
	testLog(0, solver(2, []int{3, 5, 6, 10}), t)
	testLog(0, solver(4, []int{3, 5, 6, 10}), t)
	testLog(63, solver(52, []int{1, 5, 10, 25, 26}), t)
	testLog(20, solver(32, []int{1, 5, 10, 25, 26}), t)
	testLog(0, solver(5, []int{}), t)
	testLog(1, solver(0, []int{}), t)
}

func TestCoinChange(t *testing.T) {
	testCoinChange(coinChangeIterative, t)
	testCoinChange(coinChangeRecursive, t)
}
