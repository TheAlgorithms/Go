//The Catalan numbers are a sequence of positive integers that appear in many counting
//problems in combinatorics.
//reference: https://brilliant.org/wiki/catalan-numbers/

package dynamic

// NthCatalan returns the n-th Catalan Number
func NthCatalanNumber(n int) int64 {
	if n < 0 {
		//doesn't accept negative number
		return -1
	}

	var catalanNumberList []int64
	catalanNumberList = append(catalanNumberList, 1) //first value is 1

	for i := 1; i <= n; i++ {
		catalanNumberList = append(catalanNumberList, 0) //append 0 and calculate

		for j := 0; j < i; j++ {
			catalanNumberList[i] += catalanNumberList[j] * catalanNumberList[i-j-1]
		}
	}

	return catalanNumberList[n]
}
