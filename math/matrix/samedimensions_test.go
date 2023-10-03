package matrix_test

import (
	"errors"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestSameDimensions(t *testing.T) {
	type matrixTest struct {
		name     string
		matrix1  interface{}
		matrix2  interface{}
		expected int
		err      error
	}
	tests := []matrixTest{
		{
			name: "EqualDimensionsInt",
			matrix1: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: [][]int{
				{7, 8, 9},
				{10, 11, 12},
			},
			expected: 1,
			err:      nil,
		},
		{
			name: "DifferentRowDimensionsInt",
			matrix1: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: [][]int{
				{7, 8, 9},
			},
			expected: 0,
			err:      errors.New("matrices have different dimensions: <nil>"),
		},
		{
			name: "DifferentColumnDimensionsInt",
			matrix1: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			matrix2: [][]int{
				{7, 8},
				{9, 10},
			},
			expected: 0,
			err:      errors.New("matrices have different dimensions: <nil>"),
		},
		{
			name:     "BothEmptyMatricesInt",
			matrix1:  [][]int{},
			matrix2:  [][]int{},
			expected: 1,
			err:      nil,
		},
		{
			name:    "EmptyMatrixAndNonEmptyMatrixInt",
			matrix1: [][]int{},
			matrix2: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: 0,
			err:      errors.New("matrices have different dimensions: <nil>"),
		},
		{
			name: "EqualDimensionsString",
			matrix1: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
			matrix2: [][]string{
				{"dog", "elephant"},
				{"fox", "goat"},
			},
			expected: 1,
			err:      nil,
		},
		{
			name: "DifferentRowDimensionsString",
			matrix1: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
			matrix2: [][]string{
				{"dog", "elephant"},
			},
			expected: 0,
			err:      errors.New("matrices have different dimensions: <nil>"),
		},
		{
			name: "DifferentColumnDimensionsString",
			matrix1: [][]string{
				{"apple", "banana"},
				{"cherry", "date"},
			},
			matrix2: [][]string{
				{"dog"},
				{"elephant", "fox"},
			},
			expected: 0,
			err:      errors.New("invalid matrix: rows have different numbers of columns"),
		},
		{
			name:    "EmptyMatrixAndNonEmptyMatrixString",
			matrix1: [][]string{},
			matrix2: [][]string{
				{"dog", "elephant"},
				{"fox", "goat"},
			},
			expected: 0,
			err:      errors.New("matrices have different dimensions: <nil>"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch m1 := tt.matrix1.(type) {
			case [][]int:
				switch m2 := tt.matrix2.(type) {
				case [][]int:
					result, err := matrix.SameDimensions(m1, m2)
					CheckResult(t, tt.name, result, err, tt.expected, tt.err)
				default:
					t.Errorf("Unsupported matrix2 type: %T", m2)
				}
			case [][]string:
				switch m2 := tt.matrix2.(type) {
				case [][]string:
					result, err := matrix.SameDimensions(m1, m2)
					CheckResult(t, tt.name, result, err, tt.expected, tt.err)
				default:
					t.Errorf("Unsupported matrix2 type: %T", m2)
				}
			default:
				t.Errorf("Unsupported matrix1 type: %T", m1)
			}
		})
	}
}
