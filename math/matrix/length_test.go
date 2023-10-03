package matrix_test

import (
	"errors"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestLen(t *testing.T) {
	tests := []matrixTest{
		{
			name:     "Float32Matrix1",
			matrix:   [][]float32{{12.3, 12.2}},
			expected: 1,
			err:      nil,
		},
		{
			name:     "Float32Matrix2",
			matrix:   [][]float32{{12.3, 12.2}, {12.3}},
			expected: 2,
			err:      errors.New("invalid matrix: rows have different numbers of columns"),
		},
		{
			name:     "EmptyMatrix",
			matrix:   [][]float32{},
			expected: 0,
			err:      nil,
		},
		{
			name:     "IntMatrix",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 2,
			err:      nil,
		},
		{
			name:     "StringMatrix",
			matrix:   [][]string{{"apple", "banana"}, {"cherry", "date"}},
			expected: 2,
			err:      nil,
		},
		{
			name:     "InvalidMatrix",
			matrix:   [][]string{{"apple", "banana"}, {"cherry"}},
			expected: 0,
			err:      errors.New("invalid matrix: rows have different numbers of columns"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch m := tt.matrix.(type) {
			case [][]float32:
				actual, err := matrix.Len(m)
				CheckResult(t, tt.name, actual, err, tt.expected, tt.err)
			case [][]int:
				actual, err := matrix.Len(m)
				CheckResult(t, tt.name, actual, err, tt.expected, tt.err)
			case [][]string:
				actual, err := matrix.Len(m)
				CheckResult(t, tt.name, actual, err, tt.expected, tt.err)
			default:
				t.Errorf("Unsupported matrix type: %T", m)
			}
		})
	}
}

