// matrix chain multiplication problem
// https://en.wikipedia.org/wiki/Matrix_chain_multiplication
// www.geeksforgeeks.org/dynamic_programming-set-8-matrix-chain-multiplication/
// time complexity: O(n^3)
// space complexity: O(n^2)

package dynamic

import "github.com/TheAlgorithms/Go/math/min"

// MatrixChainRec function
func MatrixChainRec(D []int, i, j int) int {
	// d[i-1] x d[i] : dimension of matrix i
	if i == j {
		return 0
	}
	q := 1 << 32
	for k := i; k < j; k++ {
		prod := MatrixChainRec(D, i, k) + MatrixChainRec(D, k+1, j) + D[i-1]*D[k]*D[j]
		q = min.Int(prod, q)
	}
	return q
}

// MatrixChainDp function
func MatrixChainDp(D []int) int {
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
				dp[i][j] = min.Int(prod, dp[i][j])
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
