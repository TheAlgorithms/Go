package matrix_test

import (
	"errors"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

type matrixTest struct {
	name     string
	matrix   interface{} //use interface{} for various matrix types for testing
	expected int
	err      error
}

func TestBreadth(t *testing.T) {
	tests := []matrixTest{
		{
			name:     "Float32Matrix1",
			matrix:   [][]float32{{12.3, 12.2}},
			expected: 2,
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
			expected: 3,
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
				actual, err := matrix.Breadth(m)
				CheckResult(t, tt.name, actual, err, tt.expected, tt.err)
			case [][]int:
				actual, err := matrix.Breadth(m)
				CheckResult(t, tt.name, actual, err, tt.expected, tt.err)
			case [][]string:
				actual, err := matrix.Breadth(m)
				CheckResult(t, tt.name, actual, err, tt.expected, tt.err)
			default:
				t.Errorf("Unsupported matrix type: %T", m)
			}
		})
	}
}

func CheckResult(t *testing.T, name string, actual int, err error, expected int, expectedErr error) {
	if err != nil {
		if expectedErr == nil {
			t.Errorf("%s: Unexpected error: %v", name, err)
		} else if err.Error() != expectedErr.Error() {
			t.Errorf("%s: Expected error: %v, but got error: %v", name, expectedErr, err)
		}
	} else {
		if expectedErr != nil {
			t.Errorf("%s: Expected error: %v, but got no error", name, expectedErr)
		}
		if actual != expected {
			t.Errorf("%s: Expected result %d, but got %d", name, expected, actual)
		}
	}
}
