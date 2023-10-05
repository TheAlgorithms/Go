package matrix

import (
	"errors"
	"fmt"

	"github.com/TheAlgorithms/Go/constraints"
)

func Multiply[T constraints.Integer](matrix1, matrix2 Matrix[T]) (Matrix[T], error) {
	// Check if the matrices can be multiplied.
	if matrix1.columns != matrix2.rows {
		return Matrix[T]{}, errors.New("matrices cannot be multiplied: column count of the first matrix must match row count of the second matrix")
	}

	// Create a new matrix to store the result.
	var zeroVal T
	result := New(matrix1.rows, matrix2.columns, zeroVal)

	for i := 0; i < matrix1.rows; i++ {
		for j := 0; j < matrix2.columns; j++ {
			// Compute the dot product of the row from the first matrix and the column from the second matrix.
			dotProduct := zeroVal
			for k := 0; k < matrix1.columns; k++ {
				dotProduct += matrix1.elements[i][k] * matrix2.elements[k][j]
			}
			err := result.Set(i, j, dotProduct)
			if err != nil {
				return Matrix[T]{}, fmt.Errorf("invalid matrices: %v", err)
			}
		}
	}

	return result, nil
}
