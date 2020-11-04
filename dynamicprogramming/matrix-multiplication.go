// matrix chain multiplication problem
// https://en.wikipedia.org/wiki/Matrix_chain_multiplication
// www.geeksforgeeks.org/dynamicprogramming-set-8-matrix-chain-multiplication/

// package main
package matrixChainMultiplication

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func matrixChainRec(D []int, i, j int) int {
	// d[i-1] x d[i] : dimension of matrix i
	if i == j {
		return 0
	}
	q := 1 << 32
	for k := i; k < j; k++ {
		prod := matrixChainRec(D, i, k) + matrixChainRec(D, k+1, j) + D[i-1]*D[k]*D[j]
		q = min(prod, q)
	}
	return q
}

func matrixChainDp(D []int) int {
	// d[i-1] x d[i] : dimension of matrix i
	N := len(D)

	dp := make([][]int, N) // dp[i][j] = matrixChainRec(D, i, j)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 0
	}

	for l := 2; l < N; l++ {
		for i := 1; i < N-l+1; i++ {
			j := i + l - 1
			dp[i][j] = 1 << 31
			for k := i; k < j; k++ {
				prod := dp[i][k] + dp[k+1][j] + D[i-1]*D[k]*D[j]
				dp[i][j] = min(prod, dp[i][j])
			}
		}
	}

	return dp[1][N-1]
}

/*
func main() {
	D := []int{2, 2, 2, 2, 2} // 4 matrices
	fmt.Print(matrixChainRec(D, 1, 4), "\n")
	fmt.Print(matrixChainDp(D), "\n")
}
*/
