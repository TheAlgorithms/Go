package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMatrixMatchDimensions(t *testing.T) {
	// Create two matrices with the same dimensions
	m1 := matrix.New(2, 3, 0)
	m2 := matrix.New(2, 3, 0)

	// Test case 1: Same dimensions
	if !m1.MatchDimensions(m2) {
		t.Errorf("m1.MatchDimensions(m2) returned %t, expected 1 (same dimensions)", m1.MatchDimensions(m2))
	}

	// Create two matrices with different dimensions
	m3 := matrix.New(2, 3, 0)
	m4 := matrix.New(3, 2, 0)

	// Test case 2: Different dimensions
	if m3.MatchDimensions(m4) {
		t.Errorf("m3.MatchDimensions(m4) returned : %v, expected: %v", m3.MatchDimensions(m4), false)
	}
}

// BenchmarkMatchDimensions benchmarks the MatchDimensions method.
func BenchmarkMatchDimensions(b *testing.B) {
	// Create sample matrices for benchmarking
	rows := 100
	columns := 100
	m1 := matrix.New(rows, columns, 0) // Replace with appropriate values
	m2 := matrix.New(rows, columns, 0) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_ = m1.MatchDimensions(m2)
	}
}
