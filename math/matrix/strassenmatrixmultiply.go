// filename: main.go
// description: Implements matrix multiplication using the Strassen algorithm.
// details:
// This program takes two matrices as input and performs matrix multiplication
// using the Strassen algorithm, which is an optimized divide-and-conquer
// approach. It allows for efficient multiplication of large matrices.
// author(s): Mohit Raghav(https://github.com/mohit07raghav19)
// See strassenmatrixmultiply_test.go for test cases
package matrix

// Perform matrix multiplication using Strassen's algorithm
func strassenMatrixMultiply(A, B [][]int) [][]int {
	n := len(A)

	// Check if matrices are 2x2 or smaller
	if n <= 2 {
		return standardMatrixMultiply(A, B)
	}

	// Calculate the size of submatrices
	mid := n / 2

	// Create submatrices
	A11 := make([][]int, mid)
	A12 := make([][]int, mid)
	A21 := make([][]int, mid)
	A22 := make([][]int, mid)

	B11 := make([][]int, mid)
	B12 := make([][]int, mid)
	B21 := make([][]int, mid)
	B22 := make([][]int, mid)

	for i := 0; i < mid; i++ {
		A11[i] = make([]int, mid)
		A12[i] = make([]int, mid)
		A21[i] = make([]int, mid)
		A22[i] = make([]int, mid)

		B11[i] = make([]int, mid)
		B12[i] = make([]int, mid)
		B21[i] = make([]int, mid)
		B22[i] = make([]int, mid)

		for j := 0; j < mid; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+mid]
			A21[i][j] = A[i+mid][j]
			A22[i][j] = A[i+mid][j+mid]

			B11[i][j] = B[i][j]
			B12[i][j] = B[i][j+mid]
			B21[i][j] = B[i+mid][j]
			B22[i][j] = B[i+mid][j+mid]
		}
	}

	// Recursive steps
	M1 := strassenMatrixMultiply(addMatrices(A11, A22), addMatrices(B11, B22))
	M2 := strassenMatrixMultiply(addMatrices(A21, A22), B11)
	M3 := strassenMatrixMultiply(A11, subtractMatrices(B12, B22))
	M4 := strassenMatrixMultiply(A22, subtractMatrices(B21, B11))
	M5 := strassenMatrixMultiply(addMatrices(A11, A12), B22)
	M6 := strassenMatrixMultiply(subtractMatrices(A21, A11), addMatrices(B11, B12))
	M7 := strassenMatrixMultiply(subtractMatrices(A12, A22), addMatrices(B21, B22))

	// Calculate result submatrices
	C11 := subtractMatrices(addMatrices(addMatrices(M1, M4), M7), M5)
	C12 := addMatrices(M3, M5)
	C21 := addMatrices(M2, M4)
	C22 := subtractMatrices(addMatrices(addMatrices(M1, M3), M6), M2)

	// Combine submatrices into the result matrix
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i < mid && j < mid {
				C[i][j] = C11[i][j]
			} else if i < mid && j >= mid {
				C[i][j] = C12[i][j-mid]
			} else if i >= mid && j < mid {
				C[i][j] = C21[i-mid][j]
			} else {
				C[i][j] = C22[i-mid][j-mid]
			}
		}
	}

	return C
}

// Add two matrices element-wise
func addMatrices(A, B [][]int) [][]int {
	n := len(A)
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = A[i][j] + B[i][j]
		}
	}

	return result
}

// Subtract one matrix from another element-wise
func subtractMatrices(A, B [][]int) [][]int {
	n := len(A)
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = A[i][j] - B[i][j]
		}
	}

	return result
}

// Perform standard matrix multiplication
func standardMatrixMultiply(A, B [][]int) [][]int {
	n := len(A)
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return result
}
