// filename: strassenmatrixmultiply.go
// description: Implements matrix multiplication using the Strassen algorithm.
// details:
// This program takes two matrices as input and performs matrix multiplication
// using the Strassen algorithm, which is an optimized divide-and-conquer
// approach. It allows for efficient multiplication of large matrices.
// author(s): Mohit Raghav(https://github.com/mohit07raghav19)
// See strassenmatrixmultiply_test.go for test cases
package matrix

// Perform matrix multiplication using Strassen's algorithm
func (A Matrix[T]) StrassenMatrixMultiply(B Matrix[T]) Matrix[T] {
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
		A1, _ := A11.Add(A22)
		A2, _ := B11.Add(B22)
		A3, _ := A21.Add(A22)
		A4, _ := A11.Add(A12)
		A5, _ := B11.Add(B12)
		A6, _ := B21.Add(B22)
		//
		S1, _ := B12.Subtract(B22)
		S2, _ := B21.Subtract(B11)
		S3, _ := A21.Subtract(A11)
		S4, _ := A12.Subtract(A22)
		// Recursive steps
		M1 := A1.StrassenMatrixMultiply(A2)
		M2 := A3.StrassenMatrixMultiply(B11)
		M3 := A11.StrassenMatrixMultiply(S1)
		M4 := A22.StrassenMatrixMultiply(S2)
		M5 := A4.StrassenMatrixMultiply(B22)
		M6 := S3.StrassenMatrixMultiply(A5)
		M7 := S4.StrassenMatrixMultiply(A6)
		//
		A7, _ := M1.Add(M4)
		A8, _ := A7.Add(M7)
		A9, _ := M1.Add(M3)
		A10, _ := A9.Add(M6)
		// Calculate result submatrices
		C11, _ := A8.Subtract(M5)
		C12, _ := M3.Add(M5)
		C21, _ := M2.Add(M4)
		C22, _ := A10.Subtract(M2)

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
