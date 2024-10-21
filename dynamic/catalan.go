//The Catalan numbers are a sequence of positive integers that appear in many counting
// problems in combinatorics.
// time complexity: O(n²)
// space complexity: O(n)
//reference: https://brilliant.org/wiki/catalan-numbers/

package dynamic

import "fmt"

var errCatalan = fmt.Errorf("can't have a negative n-th catalan number")

// NthCatalan returns the n-th Catalan Number
// Complexity: O(n²)
func NthCatalanNumber(n int) (int64, error) {
	if n < 0 {
		//doesn't accept negative number
		return 0, errCatalan
	}

	var catalanNumberList []int64
	catalanNumberList = append(catalanNumberList, 1) //first value is 1

	for i := 1; i <= n; i++ {
		catalanNumberList = append(catalanNumberList, 0) //append 0 and calculate

		for j := 0; j < i; j++ {
			catalanNumberList[i] += catalanNumberList[j] * catalanNumberList[i-j-1]
		}
	}

	return catalanNumberList[n], nil
}
