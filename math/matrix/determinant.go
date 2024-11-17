// determinant.go
// description: This method finds the determinant of a matrix.
// details: For a theoretical explanation as for what the determinant
// represents, see the [Wikipedia Article](https://en.wikipedia.org/wiki/Determinant)
// time complexity: O(n!) where n is the number of rows and columns in the matrix.
// space complexity: O(n^2) where n is the number of rows and columns in the matrix.
// author [Carter907](https://github.com/Carter907)
// see determinant_test.go

package matrix

import (
	"errors"
)

// Calculates the determinant of the matrix.
// This method only works for square matrices (e.i. matrices with equal rows and columns).
func (mat Matrix[T]) Determinant() (T, error) {

	var determinant T = 0
	var elements = mat.elements
	if mat.rows != mat.columns {

		return 0, errors.New("Matrix rows and columns must equal in order to find the determinant.")
	}

	// Specify base cases for different sized matrices.
	switch mat.rows {
	case 1:
		return elements[0][0], nil
	case 2:
		return elements[0][0]*elements[1][1] - elements[1][0]*elements[0][1], nil
	default:
		for i := 0; i < mat.rows; i++ {

			var initialValue T = 0
			minor := New(mat.rows-1, mat.columns-1, initialValue)
			// Fill the contents of minor excluding the 0th row and the ith column.
			for j, minor_i := 1, 0; j < mat.rows && minor_i < minor.rows; j, minor_i = j+1, minor_i+1 {
				for k, minor_j := 0, 0; k < mat.rows && minor_j < minor.rows; k, minor_j = k+1, minor_j+1 {
					if k != i {
						minor.elements[minor_i][minor_j] = elements[j][k]
					} else {
						minor_j-- // Decrement the column of minor to account for skipping the ith column of the matrix.
					}
				}
			}

			if i%2 == 0 {
				minor_det, _ := minor.Determinant()

				determinant += elements[0][i] * minor_det
			} else {
				minor_det, _ := minor.Determinant()

				determinant += elements[0][i] * minor_det
			}
		}
		return determinant, nil
	}
}
