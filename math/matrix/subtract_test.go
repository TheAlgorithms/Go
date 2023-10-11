package matrix_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestSubtract(t *testing.T) {
	// Create two matrices with the same dimensions for Subtraction
	m1 := matrix.New(2, 2, 1)
	m2 := matrix.New(2, 2, 2)

	// Test case 1: Valid matrix Subtraction
	subMatrix, err := m1.Subtract(m2)
	if err != nil {
		t.Errorf("Add(m1, m2) returned an error: %v, expected no error", err)
	}
	expectedMatrix := matrix.New(2, 2, -1)
	res := subMatrix.CheckEqual(expectedMatrix)
	if !res {
		t.Errorf("Add(m1, m2) returned incorrect result:\n%v\nExpected:\n%v", subMatrix, expectedMatrix)
	}

	// Create two matrices with different dimensions for Subtraction
	m3 := matrix.New(2, 2, 1)
	m4 := matrix.New(2, 3, 2)

	// Test case 2: Matrices with different dimensions
	_, err2 := m3.Subtract(m4)
	expectedError2 := fmt.Errorf("matrices are not compatible for subtraction")
	if err2 == nil || err2.Error() != expectedError2.Error() {
		t.Errorf("m3.Subtract(m4) returned error: %v, expected error: %v", err2, expectedError2)
	}

}

// BenchmarkSubtract benchmarks the Subtract function.
func BenchmarkSubtract(b *testing.B) {
	// Create sample matrices for benchmarking
	rows := 100
	columns := 100
	m1 := matrix.New(rows, columns, 2) // Replace with appropriate values
	m2 := matrix.New(rows, columns, 3) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_, _ = m1.Subtract(m2)
	}
}
