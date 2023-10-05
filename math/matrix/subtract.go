package matrix

import (
	"errors"
	"fmt"

	"github.com/TheAlgorithms/Go/constraints"
)

// Subtract subtracts the elements of the matrix (m1) from another matrix (m2) element-wise and returns a new matrix.
// The two matrices must have the same dimensions.
// If the matrices have different dimensions or if an error occurs during subtraction, an error is returned.
func Subtract[T constraints.Integer](matrix1, matrix2 Matrix[T]) (Matrix[T], error) {
	// Check if the matrices have the same dimensions.
	sameDimensions, err := matrix1.SameDimensions(matrix2)
	if err != nil {
		return Matrix[T]{}, fmt.Errorf("invalid matrices: %v", err)
	}

	if !sameDimensions {
		return Matrix[T]{}, errors.New("matrices are not compatible for addition")
	}

	// Create a new matrix to store the result.
	var zeroVal T // zero value for default values
	result := New(matrix1.rows, matrix1.columns, zeroVal)
	for i := range matrix1.elements {
		for j := range matrix1.elements[i] {
			// Perform addition element-wise and set the result in the new matrix.
			sum := matrix1.elements[i][j] - matrix2.elements[i][j]
			err := result.Set(i, j, sum)
			if err != nil {
				return Matrix[T]{}, fmt.Errorf("invalid matrices: %v", err)
			}
		}
	}

	return result, nil
}
