package matrix

import (
	"errors"
	"fmt"
)

func CheckEqual[T comparable](matrix1, matrix2 Matrix[T]) (bool, error) {
	sameDimensions, err := matrix1.SameDimensions(matrix2)
	if err != nil {
		return false, fmt.Errorf("invalid matrices: %v", err)
	}

	if !sameDimensions {
		return false, errors.New("matrices are not compatible for comparison")
	}

	for i := range matrix1.elements {
		for j := range matrix1.elements[i] {
			if matrix1.elements[i][j] != matrix2.elements[i][j] {
				return false, nil
			}
		}
	}
	return true, nil
}
