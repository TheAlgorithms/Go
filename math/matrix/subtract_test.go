package matrix_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestSubtract(t *testing.T) {
	// Create two matrices with the same dimensions for addition
	matrix1 := matrix.New(2, 2, 1)
	matrix2 := matrix.New(2, 2, 2)

	// Test case 1: Valid matrix addition
	subMatrix, err := matrix.Subtract(matrix1, matrix2)
	if err != nil {
		t.Errorf("Add(matrix1, matrix2) returned an error: %v, expected no error", err)
	}
	expectedMatrix := matrix.New(2, 2, -1)
	res, _ := matrix.CheckEqual(subMatrix, expectedMatrix)
	if !res {
		t.Errorf("Add(matrix1, matrix2) returned incorrect result:\n%v\nExpected:\n%v", subMatrix, expectedMatrix)
	}

	// Create two matrices with different dimensions for addition
	matrix3 := matrix.New(2, 2, 1)
	matrix4 := matrix.New(2, 3, 2)

	// Test case 2: Matrices with different dimensions
	_, err2 := matrix.Subtract(matrix3, matrix4)
	expectedError2 := fmt.Errorf("invalid matrices: %v", errors.New("matrices have different dimensions: rows=2 vs columns=2"))
	if err2 == nil || err2.Error() != expectedError2.Error() {
		t.Errorf("Add(matrix3, matrix4) returned error: %v, expected error: %v", err2, expectedError2)
	}

}

// BenchmarkSubtract benchmarks the Subtract function.
func BenchmarkSubtract(b *testing.B) {
	// Create sample matrices for benchmarking
	rows := 100
	columns := 100
	matrix1 := matrix.New(rows, columns, 2) // Replace with appropriate values
	matrix2 := matrix.New(rows, columns, 3) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_, _ = matrix.Subtract(matrix1, matrix2)
	}
}
