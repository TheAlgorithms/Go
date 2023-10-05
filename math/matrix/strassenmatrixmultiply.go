// filename: strassenmatrixmultiply.go
// description: Implements matrix multiplication using the Strassen algorithm.
// details:
// This program takes two matrices as input and performs matrix multiplication
// using the Strassen algorithm, which is an optimized divide-and-conquer
// approach. It allows for efficient multiplication of large matrices.
// author(s): Mohit Raghav(https://github.com/mohit07raghav19)
// See strassenmatrixmultiply_test.go for test cases
package matrix

import (
	"github.com/TheAlgorithms/Go/constraints"
)

// Perform matrix multiplication using Strassen's algorithm
func StrassenMatrixMultiply[T constraints.Integer](A, B Matrix[T]) Matrix[T] {
	n := A.rows
	// Check if matrices are 2x2 or smaller
	if n == 1 {
		a1, _ := A.Get(0, 0)
		b1, _ := B.Get(0, 0)
		result := New(1, 1, a1*b1)
		return result
	} else {
		// Calculate the size of submatrices
		mid := n / 2

		// Create submatrices
		A11, _ := A.SubMatrix(0, 0, mid, mid)
		A12, _ := A.SubMatrix(0, mid, mid, n-mid)
		A21, _ := A.SubMatrix(mid, 0, n-mid, mid)
		A22, _ := A.SubMatrix(mid, mid, n-mid, n-mid)

		B11, _ := B.SubMatrix(0, 0, mid, mid)
		B12, _ := B.SubMatrix(0, mid, mid, n-mid)
		B21, _ := B.SubMatrix(mid, 0, n-mid, mid)
		B22, _ := B.SubMatrix(mid, mid, n-mid, n-mid)

		// Calculate result submatrices
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
		var zeroVal T
		C := New(n, n, zeroVal)

		for i := 0; i < mid; i++ {
			for j := 0; j < mid; j++ {
				val, _ := C11.Get(i, j)

				err := C.Set(i, j, val)
				if err != nil {
					panic("copyMatrix.Set error: " + err.Error())
				}

				val, _ = C12.Get(i, j)

				err1 := C.Set(i, j+mid, val)
				if err1 != nil {
					panic("copyMatrix.Set error: " + err1.Error())
				}
				val, _ = C21.Get(i, j)

				err2 := C.Set(i+mid, j, val)
				if err2 != nil {
					panic("copyMatrix.Set error: " + err2.Error())
				}
				val, _ = C22.Get(i, j)

				err3 := C.Set(i+mid, j+mid, val)
				if err3 != nil {
					panic("copyMatrix.Set error: " + err3.Error())
				}
			}
		}

		return C

	}
}
