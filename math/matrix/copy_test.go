package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMatrixCopy(t *testing.T) {
	// Create a sample matrix
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// Ensure that the copy is not the same as the original
	matrix, err := matrix.NewFromElements(data)
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}
	copyMatrix, err := matrix.Copy()
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}

	// Ensure that the copy is not the same as the original
	if &matrix == &copyMatrix {
		t.Errorf("Copy did not create a new matrix.")
	}

	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Columns(); j++ {
			val1, err := matrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			val2, err := copyMatrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			if val1 != val2 {
				t.Errorf("Copy did not correctly copy element (%d, %d).", i, j)
			}
		}
	}

}

func TestMatrixCopyEmpty(t *testing.T) {
	// Create an empty matrix
	emptyMatrix := matrix.New(0, 0, 0)

	// Make a copy of the empty matrix
	copyMatrix, err := emptyMatrix.Copy()
	if err != nil { // as empty matrix
		t.Fatalf("Failed to copy matrix: %v", err)
	}

	// Ensure that the copy is not the same as the original by comparing their addresses
	if &emptyMatrix == &copyMatrix {
		t.Errorf("Copy did not create a new matrix for an empty matrix.")
	}

	// Check if the copy is also empty
	if copyMatrix.Rows() != 0 || copyMatrix.Columns() != 0 {
		t.Errorf("Copy of an empty matrix should also be empty.")
	}
}

func TestMatrixCopyWithDefaultValues(t *testing.T) {
	// Create a matrix with default values (zeroes)
	rows, columns := 3, 3
	defaultValue := 0
	defaultMatrix := matrix.New(rows, columns, defaultValue)

	// Make a copy of the matrix
	copyMatrix, err := defaultMatrix.Copy()
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}

	// Ensure that the copy is not the same as the original by comparing their addresses
	if &defaultMatrix == &copyMatrix {
		t.Errorf("Copy did not create a new matrix for default values.")
	}

	// Check if the copy has the same values as the original (all zeroes)
	for i := 0; i < defaultMatrix.Rows(); i++ {
		for j := 0; j < defaultMatrix.Columns(); j++ {
			val1, err := defaultMatrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			val2, err := copyMatrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			if val1 != val2 || val1 != defaultValue || val2 != defaultValue {
				t.Errorf("Copy did not preserve default values at row %d, column %d. Expected %v, got %v", i, j, defaultValue, val2)
			}
		}
	}
}

func BenchmarkCopyMatrix(b *testing.B) {
	// Create a matrix for benchmarking
	rows := 100
	columns := 100
	initialValue := 0
	matrix := matrix.New(rows, columns, initialValue)

	// Reset the benchmark timer
	b.ResetTimer()

	// Run the benchmarks
	for i := 0; i < b.N; i++ {
		_, _ = matrix.Copy()
	}
}
