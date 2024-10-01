// filename: strassenmatrixmultiply.go
// description: Implements matrix multiplication using the Strassen algorithm.
// details:
// This program takes two matrices as input and performs matrix multiplication
// using the Strassen algorithm, which is an optimized divide-and-conquer
// approach. It allows for efficient multiplication of large matrices.
// time complexity: O(n^2.81)
// space complexity: O(n^2)
// author(s): Mohit Raghav(https://github.com/mohit07raghav19)
// See strassenmatrixmultiply_test.go for test cases
package matrix

// Perform matrix multiplication using Strassen's algorithm
func (A Matrix[T]) StrassenMatrixMultiply(B Matrix[T]) (Matrix[T], error) {
	n := A.rows
	// Check if matrices are 2x2 or smaller
	if n == 1 {
		a1, err := A.Get(0, 0)
		if err != nil {
			return Matrix[T]{}, err
		}
		b1, err := B.Get(0, 0)
		if err != nil {
			return Matrix[T]{}, err
		}
		result := New(1, 1, a1*b1)
		return result, nil
	} else {
		// Calculate the size of submatrices
		mid := n / 2

		// Create submatrices
		A11, err := A.SubMatrix(0, 0, mid, mid)
		if err != nil {
			return Matrix[T]{}, err
		}
		A12, err := A.SubMatrix(0, mid, mid, n-mid)
		if err != nil {
			return Matrix[T]{}, err
		}
		A21, err := A.SubMatrix(mid, 0, n-mid, mid)
		if err != nil {
			return Matrix[T]{}, err
		}
		A22, err := A.SubMatrix(mid, mid, n-mid, n-mid)
		if err != nil {
			return Matrix[T]{}, err
		}

		B11, err := B.SubMatrix(0, 0, mid, mid)
		if err != nil {
			return Matrix[T]{}, err
		}
		B12, err := B.SubMatrix(0, mid, mid, n-mid)
		if err != nil {
			return Matrix[T]{}, err
		}
		B21, err := B.SubMatrix(mid, 0, n-mid, mid)
		if err != nil {
			return Matrix[T]{}, err
		}
		B22, err := B.SubMatrix(mid, mid, n-mid, n-mid)
		if err != nil {
			return Matrix[T]{}, err
		}

		// Calculate result submatrices
		A1, err := A11.Add(A22)
		if err != nil {
			return Matrix[T]{}, err
		}

		A2, err := B11.Add(B22)
		if err != nil {
			return Matrix[T]{}, err
		}

		A3, err := A21.Add(A22)
		if err != nil {
			return Matrix[T]{}, err
		}

		A4, err := A11.Add(A12)
		if err != nil {
			return Matrix[T]{}, err
		}

		A5, err := B11.Add(B12)
		if err != nil {
			return Matrix[T]{}, err
		}

		A6, err := B21.Add(B22)
		if err != nil {
			return Matrix[T]{}, err
		}
		//
		S1, err := B12.Subtract(B22)
		if err != nil {
			return Matrix[T]{}, err
		}
		S2, err := B21.Subtract(B11)
		if err != nil {
			return Matrix[T]{}, err
		}
		S3, err := A21.Subtract(A11)
		if err != nil {
			return Matrix[T]{}, err
		}
		S4, err := A12.Subtract(A22)
		if err != nil {
			return Matrix[T]{}, err
		}
		// Recursive steps
		M1, err := A1.StrassenMatrixMultiply(A2)
		if err != nil {
			return Matrix[T]{}, err
		}
		M2, err := A3.StrassenMatrixMultiply(B11)
		if err != nil {
			return Matrix[T]{}, err
		}
		M3, err := A11.StrassenMatrixMultiply(S1)
		if err != nil {
			return Matrix[T]{}, err
		}
		M4, err := A22.StrassenMatrixMultiply(S2)
		if err != nil {
			return Matrix[T]{}, err
		}
		M5, err := A4.StrassenMatrixMultiply(B22)
		if err != nil {
			return Matrix[T]{}, err
		}
		M6, err := S3.StrassenMatrixMultiply(A5)
		if err != nil {
			return Matrix[T]{}, err
		}
		M7, err := S4.StrassenMatrixMultiply(A6)

		if err != nil {
			return Matrix[T]{}, err
		} //
		A7, err := M1.Add(M4)

		if err != nil {
			return Matrix[T]{}, err
		}
		A8, err := A7.Add(M7)

		if err != nil {
			return Matrix[T]{}, err
		}
		A9, err := M1.Add(M3)

		if err != nil {
			return Matrix[T]{}, err
		}
		A10, err := A9.Add(M6)

		if err != nil {
			return Matrix[T]{}, err
		}
		// Calculate result submatrices
		C11, err := A8.Subtract(M5)
		if err != nil {
			return Matrix[T]{}, err
		}
		C12, err := M3.Add(M5)
		if err != nil {
			return Matrix[T]{}, err
		}
		C21, err := M2.Add(M4)
		if err != nil {
			return Matrix[T]{}, err
		}
		C22, err := A10.Subtract(M2)
		if err != nil {
			return Matrix[T]{}, err
		}

		// Combine subMatrices into the result matrix
		var zeroVal T
		C := New(n, n, zeroVal)

		for i := 0; i < mid; i++ {
			for j := 0; j < mid; j++ {
				val, err := C11.Get(i, j)
				if err != nil {
					return Matrix[T]{}, err
				}

				err = C.Set(i, j, val)
				if err != nil {
					return Matrix[T]{}, err
				}

				val, err = C12.Get(i, j)
				if err != nil {
					return Matrix[T]{}, err
				}

				err1 := C.Set(i, j+mid, val)
				if err1 != nil {
					return Matrix[T]{}, err1
				}
				val, err = C21.Get(i, j)
				if err != nil {
					return Matrix[T]{}, err
				}

				err2 := C.Set(i+mid, j, val)
				if err2 != nil {
					return Matrix[T]{}, err2
				}
				val, err = C22.Get(i, j)
				if err != nil {
					return Matrix[T]{}, err
				}

				err3 := C.Set(i+mid, j+mid, val)
				if err3 != nil {
					return Matrix[T]{}, err3
				}
			}
		}
		return C, nil
	}
}
