// filename: strassenmatrixmultiply.go
// description: Implements matrix multiplication using the Strassen algorithm.
// details:
// This program takes two matrices as input and performs matrix multiplication
// using the Strassen algorithm, which is an optimized divide-and-conquer
// approach. It allows for efficient multiplication of large matrices.
// author(s): Mohit Raghav(https://github.com/mohit07raghav19)
// See strassenmatrixmultiply_test.go for test cases
package matrix

import "github.com/TheAlgorithms/Go/constraints"

// Perform matrix multiplication using Strassen's algorithm
func StrassenMatrixMultiply[T constraints.Integer](A, B [][]T) [][]T {
	n := len(A)

	// Check if matrices are 2x2 or smaller
	if n == 1 {
		result := make([][]T, 1)
		result[0] = make([]T, 1)
		result[0][0] = A[0][0] * B[0][0]
		return result
	} else {
		// Calculate the size of submatrices
		mid := n / 2

		// Create submatrices
		A11 := make([][]T, mid)
		A12 := make([][]T, mid)
		A21 := make([][]T, mid)
		A22 := make([][]T, mid)

		B11 := make([][]T, mid)
		B12 := make([][]T, mid)
		B21 := make([][]T, mid)
		B22 := make([][]T, mid)

		for i := 0; i < mid; i++ {
			A11[i] = make([]T, mid)
			A12[i] = make([]T, mid)
			A21[i] = make([]T, mid)
			A22[i] = make([]T, mid)

			B11[i] = make([]T, mid)
			B12[i] = make([]T, mid)
			B21[i] = make([]T, mid)
			B22[i] = make([]T, mid)

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

		// Calculate result addmatrices
		A1, _ := Add(A11, A22)
		A2, _ := Add(B11, B22)
		A3, _ := Add(A21, A22)
		A4, _ := Add(A11, A12)
		A5, _ := Add(B11, B12)
		A6, _ := Add(B21, B22)
		//
		S1, _ := Subtract(B12, B22)
		S2, _ := Subtract(B21, B11)
		S3, _ := Subtract(A21, A11)
		S4, _ := Subtract(A12, A22)
		// Recursive steps
		M1 := StrassenMatrixMultiply(A1, A2)
		M2 := StrassenMatrixMultiply(A3, B11)
		M3 := StrassenMatrixMultiply(A11, S1)
		M4 := StrassenMatrixMultiply(A22, S2)
		M5 := StrassenMatrixMultiply(A4, B22)
		M6 := StrassenMatrixMultiply(S3, A5)
		M7 := StrassenMatrixMultiply(S4, A6)
		//
		A7, _ := Add(M1, M4)
		A8, _ := Add(A7, M7)
		A9, _ := Add(M1, M3)
		A10, _ := Add(A9, M6)
		// Calculate result submatrices
		C11, _ := Subtract(A8, M5)
		C12, _ := Add(M3, M5)
		C21, _ := Add(M2, M4)
		C22, _ := Subtract(A10, M2)

		// Combine subMatrices into the result matrix
		C := make([][]T, n)

		for i := 0; i < mid; i++ {
			C[i] = make([]T, n)
			C[i+mid] = make([]T, n)

			for j := 0; j < mid; j++ {
				C[i][j] = C11[i][j]
				C[i][j+mid] = C12[i][j]
				C[i+mid][j] = C21[i][j]
				C[i+mid][j+mid] = C22[i][j]
			}
		}
		return C
	}
}
