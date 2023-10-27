package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMultiplyMatrix(t *testing.T) {
	// Test case with compatible NULL matrices
	t.Run("NULL Matrices", func(t *testing.T) {
		mat1 := matrix.New(0, 0, 0)
		mat2 := matrix.New(0, 0, 0)

		expected := matrix.New(0, 0, 0)
		result, err := mat1.Multiply(mat2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		} else if !result.CheckEqual(expected) {
			t.Errorf("Result matrix does not match the expected result.")
		}

	})
	// Test case with compatible matrices
	t.Run("Compatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2, 3}, {4, 5, 6}}
		mat2 := [][]int{{7, 8}, {9, 10}, {11, 12}}
		m1, err := matrix.NewFromElements(mat1)
		if err != nil {
			t.Fatalf("Failed to copy matrix: %v", err)
		}
		m2, err := matrix.NewFromElements(mat2)
		if err != nil {
			t.Fatalf("Failed to copy matrix: %v", err)
		}
		exp := [][]int{{58, 64}, {139, 154}}
		expected, err := matrix.NewFromElements(exp)
		if err != nil {
			t.Fatalf("Failed to copy matrix: %v", err)
		}
		result, err := m1.Multiply(m2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		} else if !result.CheckEqual(expected) {
			t.Errorf("Result matrix does not match the expected result.")
		}

	})

}

func TestMultiplyIncompatibleMatrix(t *testing.T) {
	// Test case with incompatible matrices
	t.Run("Incompatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2, 3}, {4, 5, 6}}
		mat2 := [][]int{{7, 8}, {9, 10}}
		m1, err := matrix.NewFromElements(mat1)
		if err != nil {
			t.Fatalf("Failed to copy matrix: %v", err)
		}
		m2, err := matrix.NewFromElements(mat2)
		if err != nil {
			t.Fatalf("Failed to copy matrix: %v", err)
		}

		_, err = m1.Multiply(m2)
		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})

	t.Run("Incompatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2}}
		mat2 := [][]int{{}}
		m1, err := matrix.NewFromElements(mat1)
		if err != nil {
			t.Fatalf("Failed to copy matrix: %v", err)
		}
		m2, err := matrix.NewFromElements(mat2)
		if err != nil {
			t.Fatalf("Failed to copy matrix: %v", err)
		}
		_, err = m1.Multiply(m2)

		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})
}

func BenchmarkMatrixMultiply(b *testing.B) {
	// Create sample matrices for benchmarking
	rows := 10
	columns := 10
	m1 := matrix.New(rows, columns, 2) // Replace with appropriate values
	m2 := matrix.New(rows, columns, 3) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_, _ = m1.Multiply(m2)
	}
}
