package matrix

import (
	"errors"
	"fmt"

	"github.com/TheAlgorithms/Go/constraints"
)

// Subtract subtracts two matrices of compatible types.
func Subtract[T constraints.Integer](matrix1, matrix2 [][]T) ([][]T, error) {
	// Check if the matrices have the same dimensions.
	sameDimensions, err := SameDimensions(matrix1, matrix2)
	if err != nil {
		return nil, fmt.Errorf("invalid matrices: %v", err)
	}

	if sameDimensions==0 {
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
