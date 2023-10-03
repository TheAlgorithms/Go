package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestIsValidMatrix(t *testing.T) {
	validMatrices := [][][]float32{
		{{1, 2, 3}, {4, 5, 6}},
		{{7, 8}, {9, 10}},
		{{7.2, 8.2}, {9.3, 10.1}},
		{},
		{nil},
	}

	invalidMatrices := [][][]int{
		{{1, 2, 3}, {4, 5}},
		{{7, 9}, {10}},
	}

	for _, m := range validMatrices {
		if matrix.IsValid(m) != nil {
			t.Errorf("Expected valid matrix, but got invalid: %v", m)
		}
	}

	for _, m := range invalidMatrices {
		if matrix.IsValid(m) == nil {
			t.Errorf("Expected invalid matrix, but got valid: %v", m)
		}
	}
}
