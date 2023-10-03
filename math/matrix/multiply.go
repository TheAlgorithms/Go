package matrix

import (
	"errors"

	"github.com/TheAlgorithms/Go/constraints"
)

// MultiplyMatrices multiplies two matrices of compatible types.
func Multiply[T constraints.Integer](matrix1, matrix2 [][]T) ([][]T, error) {
	// Check if the matrices are compatible for multiplication.
	if len(matrix1[0]) == len(matrix2[0]) && len(matrix1[0]) == 0 {
		return make([][]T, len(matrix1)), nil
	}
	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("matrices are not compatible for multiplication")
	}

	// Create the result matrix with appropriate dimensions.
	result := make([][]T, len(matrix1))
	for i := range result {
		result[i] = make([]T, len(matrix2[0]))
	}

	// Perform matrix multiplication.
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix2[0]); j++ {
			var sum T
			for k := 0; k < len(matrix2); k++ {
				sum += matrix1[i][k] * matrix2[k][j]
			}
			result[i][j] = sum
		}
	}
	return result, nil
}
