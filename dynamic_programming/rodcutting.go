// Solution to Rod cutting problem
// https://en.wikipedia.org/wiki/Cutting_stock_problem
// http://www.geeksforgeeks.org/dynamic-programming-set-13-cutting-a-rod/

package dynamic_programming

// CutRodRec solve the problem recursively: initial approach
func CutRodRec(price []int, length int) int {
	if length == 0 {
		return 0
	}

	q := -1
	for i := 1; i <= length; i++ {
		q = Max(q, price[i]+CutRodRec(price, length-i))
	}
	return q
}

// CutRodDp solve the same problem using dynamic programming
func CutRodDp(price []int, length int) int {
	r := make([]int, length+1) // a.k.a the memoization array
	r[0] = 0                   // cost of 0 length rod is 0

	for j := 1; j <= length; j++ { // for each length (subproblem)
		q := -1
		for i := 1; i <= j; i++ {
			q = Max(q, price[i]+r[j-i]) // avoiding recursive call
		}
		r[j] = q
	}

	return r[length]
}
