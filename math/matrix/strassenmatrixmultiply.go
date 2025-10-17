// filename: strassenmatrixmultiply.go
// description: Implements matrix multiplication using the Strassen algorithm.
// details:
// This program takes two matrices as input and performs matrix multiplication
// using the Strassen algorithm, which is an optimized divide-and-conquer
// approach. It allows for efficient multiplication of large matrices.
//
// Example:
// | 1  2 | * | 5  6 | = | (1*5 + 2*7)  (1*6 + 2*8) | = | 19  22 |
// | 3  4 |   | 7  8 |   | (3*5 + 4*7)  (3*6 + 4*8) |   | 43  50 |
//
// Strassen's algorithm achieves a better time complexity by using
// 7 recursive multiplications instead of the 8 used in a standard
// divide-and-conquer approach.
//
// time complexity: O(n^log2(7)) ≈ O(n^2.81)
// space complexity: O(n^2) for storing submatrices
// author(s): Mohit Raghav(https://github.com/mohit07raghav19)
// See strassenmatrixmultiply_test.go for test cases
package matrix

// StrassenMatrixMultiply performs matrix multiplication using Strassen's
// divide-and-conquer algorithm.
//
// NOTE: This implementation expects the matrices to be square (n x n).
// The base case for the recursion is a 1x1 matrix.
func (A Matrix[T]) StrassenMatrixMultiply(B Matrix[T]) (Matrix[T], error) {
	n := A.rows
	// Base case for the recursion:
	// If the matrix is 1x1, perform a single multiplication.
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
		// --- 1. DIVIDE ---
		// Calculate the size of submatrices (quadrants)
		mid := n / 2

		// Create submatrices (A11, A12, A21, A22 and B11, B12, B21, B22)
		// A = | A11  A12 |
		//     | A21  A22 |
		// B = | B11  B12 |
		//     | B21  B22 |
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

		// --- 2. CALCULATE INTERMEDIATE TERMS ---
		// These are the 10 additions/subtractions required to
		// set up the 7 recursive multiplication steps.
		A1, err := A11.Add(A22) // A1 = A11 + A22
		if err != nil {
			return Matrix[T]{}, err
		}

		A2, err := B11.Add(B22) // A2 = B11 + B22
		if err != nil {
			return Matrix[T]{}, err
		}

		A3, err := A21.Add(A22) // A3 = A21 + A22
		if err != nil {
			return Matrix[T]{}, err
		}

		A4, err := A11.Add(A12) // A4 = A11 + A12
		if err != nil {
			return Matrix[T]{}, err
		}

		A5, err := B11.Add(B12) // A5 = B11 + B12
		if err != nil {
			return Matrix[T]{}, err
		}

		A6, err := B21.Add(B22) // A6 = B21 + B22
		if err != nil {
			return Matrix[T]{}, err
		}
		//
		S1, err := B12.Subtract(B22) // S1 = B12 - B22
		if err != nil {
			return Matrix[T]{}, err
		}
		S2, err := B21.Subtract(B11) // S2 = B21 - B11
		if err != nil {
			return Matrix[T]{}, err
		}
		S3, err := A21.Subtract(A11) // S3 = A21 - A11
		if err != nil {
			return Matrix[T]{}, err
		}
		S4, err := A12.Subtract(A22) // S4 = A12 - A22
		if err != nil {
			return Matrix[T]{}, err
		}

		// --- 3. CONQUER ---
		// Recursive steps: Calculate the 7 Strassen products (M1 to M7).
		// M1 = (A11 + A22) * (B11 + B22)
		M1, err := A1.StrassenMatrixMultiply(A2)
		if err != nil {
			return Matrix[T]{}, err
		}
		// M2 = (A21 + A22) * B11
		M2, err := A3.StrassenMatrixMultiply(B11)
		if err != nil {
			return Matrix[T]{}, err
		}
		// M3 = A11 * (B12 - B22)
		M3, err := A11.StrassenMatrixMultiply(S1)
		if err != nil {
			return Matrix[T]{}, err
		}
		// M4 = A22 * (B21 - B11)
		M4, err := A22.StrassenMatrixMultiply(S2)
		if err != nil {
			return Matrix[T]{}, err
		}
		// M5 = (A11 + A12) * B22
		M5, err := A4.StrassenMatrixMultiply(B22)
		if err != nil {
			return Matrix[T]{}, err
		}
		// M6 = (A21 - A11) * (B11 + B12)
		M6, err := S3.StrassenMatrixMultiply(A5)
		if err != nil {
			return Matrix[T]{}, err
		}
		// M7 = (A12 - A22) * (B21 + B22)
		M7, err := S4.StrassenMatrixMultiply(A6)

		if err != nil {
			return Matrix[T]{}, err
		} //
		// (Temporary combinations for calculating C submatrices)
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

		// --- 4. COMBINE ---
		// Calculate result submatrices (C11, C12, C21, C22)
		// C11 = M1 + M4 - M5 + M7
		C11, err := A8.Subtract(M5)
		if err != nil {
			return Matrix[T]{}, err
		}
		// C12 = M3 + M5
		C12, err := M3.Add(M5)
		if err != nil {
			return Matrix[T]{}, err
		}
		// C21 = M2 + M4
		C21, err := M2.Add(M4)
		if err != nil {
			return Matrix[T]{}, err
		}
		// C22 = M1 - M2 + M3 + M6
		C22, err := A10.Subtract(M2)
		if err != nil {
			return Matrix[T]{}, err
		}

		// Combine subMatrices into the final result matrix C
		// C = | C11  C12 |
		//     | C21  C22 |
		var zeroVal T
		C := New(n, n, zeroVal)

		for i := 0; i < mid; i++ {
			for j := 0; j < mid; j++ {
				// Set C11 (top-left quadrant)
				val, err := C11.Get(i, j)
				if err != nil {
					return Matrix[T]{}, err
				}

				err = C.Set(i, j, val)
				if err != nil {
					return Matrix[T]{}, err
				}

				// Set C12 (top-right quadrant)
				val, err = C12.Get(i, j)
				if err != nil {
					return Matrix[T]{}, err
				}

				err1 := C.Set(i, j+mid, val)
				if err1 != nil {
					return Matrix[T]{}, err1
				}
				
				// Set C21 (bottom-left quadrant)
				val, err = C21.Get(i, j)
				if err != nil {
					return Matrix[T]{}, err
				}

				err2 := C.Set(i+mid, j, val)
				if err2 != nil {
					return Matrix[T]{}, err2
				}
				
				// Set C22 (bottom-right quadrant)
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
