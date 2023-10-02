package matrix

import (
	"errors"
)

// SubtractMatrices subtracts two matrices of compatible types.
func SubtractMatrices[T Number](matrix1, matrix2 [][]T) ([][]T, error) {
	// Check if the matrices have the same dimensions.
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		return nil, errors.New("matrices are not compatible for subtraction")
	}

	result := make([][]T, len(matrix1))
	for i := range result {
		result[i] = make([]T, len(matrix1[i]))
		for j := range result[i] {
			result[i][j] = matrix1[i][j] - matrix2[i][j]
		}
	}
	return result, nil
}
