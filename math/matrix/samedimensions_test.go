package matrix_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMatrixSameDimensions(t *testing.T) {
	// Create two matrices with the same dimensions
	matrix1 := matrix.New(2, 3, 0)
	matrix2 := matrix.New(2, 3, 0)

	// Test case 1: Same dimensions
	result1, err1 := matrix1.SameDimensions(matrix2)
	if err1 != nil {
		t.Errorf("matrix1.SameDimensions(matrix2) returned an error: %v, expected no error", err1)
	}
	if !result1 {
		t.Errorf("matrix1.SameDimensions(matrix2) returned %t, expected 1 (same dimensions)", result1)
	}

	// Create two matrices with different dimensions
	matrix3 := matrix.New(2, 3, 0)
	matrix4 := matrix.New(3, 2, 0)

	// Test case 2: Different dimensions
	result2, err2 := matrix3.SameDimensions(matrix4)
	expectedError2 := fmt.Errorf("matrices have different dimensions: rows=%d vs columns=%d", 2, 3)
	if err2 == nil || err2.Error() != expectedError2.Error() {
		t.Errorf("matrix3.SameDimensions(matrix4) returned error: %v, expected error: %v", err2, expectedError2)
	}
	if result2 {
		t.Errorf("matrix3.SameDimensions(matrix4) returned %t, expected 0 (different dimensions)", result2)
	}
}

// BenchmarkSameDimensions benchmarks the SameDimensions method.
func BenchmarkSameDimensions(b *testing.B) {
	// Create sample matrices for benchmarking
	rows := 100
	columns := 100
	matrix1 := matrix.New(rows, columns, 0) // Replace with appropriate values
	matrix2 := matrix.New(rows, columns, 0) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_, _ = matrix1.SameDimensions(matrix2)
	}
}
