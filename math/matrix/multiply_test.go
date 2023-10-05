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
		result, err := matrix.Multiply(mat1, mat2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		} else {
			if equal, err := matrix.CheckEqual(result, expected); err != nil {
				t.Errorf("Error: %v", err)
			} else if !equal {
				t.Errorf("Result matrix does not match the expected result.")
			}
		}

	})
	// Test case with compatible matrices
	t.Run("Compatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2, 3}, {4, 5, 6}}
		mat2 := [][]int{{7, 8}, {9, 10}, {11, 12}}
		matrix1, _ := matrix.NewFromElements(mat1)
		matrix2, _ := matrix.NewFromElements(mat2)
		exp := [][]int{{58, 64}, {139, 154}}
		expected, _ := matrix.NewFromElements(exp)
		result, err := matrix.Multiply(matrix1, matrix2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		} else {
			if equal, err := matrix.CheckEqual(result, expected); err != nil {
				t.Errorf("Error: %v", err)
			} else if !equal {
				t.Errorf("Result matrix does not match the expected result.")
			}
		}

	})

}

func TestMultiplyIncompatibleMatrix(t *testing.T) {
	// Test case with incompatible matrices
	t.Run("Incompatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2, 3}, {4, 5, 6}}
		mat2 := [][]int{{7, 8}, {9, 10}}
		matrix1, _ := matrix.NewFromElements(mat1)
		matrix2, _ := matrix.NewFromElements(mat2)

		_, err := matrix.Multiply(matrix1, matrix2)
		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})

	t.Run("Incompatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2}}
		mat2 := [][]int{{}}
		matrix1, _ := matrix.NewFromElements(mat1)
		matrix2, _ := matrix.NewFromElements(mat2)
		_, err := matrix.Multiply(matrix1, matrix2)

		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})
}

// BenchmarkMultiply benchmarks the Multiply function.
func BenchmarkMultiply(b *testing.B) {
	// Create sample matrices for benchmarking
	rows := 100
	columns := 100
	matrix1 := matrix.New(rows, columns, 2) // Replace with appropriate values
	matrix2 := matrix.New(rows, columns, 3) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_, _ = matrix.Multiply(matrix1, matrix2)
	}
}
