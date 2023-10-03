package matrix

import (
	"errors"
	"fmt"
)

func CheckEqual[T comparable](mat1, mat2 [][]T) (bool, error) {
	sameDimensions, err := SameDimensions(mat1, mat2)
	if err != nil {
		return false, fmt.Errorf("invalid matrices: %v", err)
	}

	if sameDimensions == 0 {
		return false, errors.New("matrices are not compatible for comparison")
	}

	for i := range mat1 {
		for j := range mat1[i] {
			if mat1[i][j] != mat2[i][j] {
				return false, nil
			}
		}
	}
	return true, nil
}
