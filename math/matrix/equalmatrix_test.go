package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMatrixEqual(t *testing.T) {
	// Test case with equal matrices
	t.Run("Equal Null Matrices", func(t *testing.T) {
		mat1 := [][]int{{}}
		mat2 := [][]int{{}}

		if !matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to be equal, but they are not")
		}
	})
	// Test case with equal matrices
	t.Run("Equal Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2}, {3, 4}}
		mat2 := [][]int{{1, 2}, {3, 4}}

		if !matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to be equal, but they are not")
		}
	})

	// Test case with matrices of different dimensions
	t.Run("Different Dimensions", func(t *testing.T) {
		mat1 := [][]int{{1, 2}, {3, 4}}
		mat2 := [][]int{{1, 2, 3}, {4, 5, 6}}

		if matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to have different dimensions, but they are equal")
		}
	})

	// Test case with matrices that are not equal
	t.Run("Different Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2}, {3, 4}}
		mat2 := [][]int{{1, 2}, {5, 6}}

		if matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to be different, but they are equal")
		}
	})
	// Test case with equal matrices of type bool
	t.Run("Equal Bool Matrices", func(t *testing.T) {
		mat1 := [][]bool{{true, false}, {false, true}}
		mat2 := [][]bool{{true, false}, {false, true}}

		if !matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to be equal, but they are not")
		}
	})

	// Test case with equal matrices of type string
	t.Run("Equal String Matrices", func(t *testing.T) {
		mat1 := [][]string{{"apple", "banana"}, {"cherry", "date"}}
		mat2 := [][]string{{"apple", "banana"}, {"cherry", "date"}}

		if !matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to be equal, but they are not")
		}
	})
	// Test case with unequal matrices of type string
	t.Run("Different string Matrices", func(t *testing.T) {
		mat1 := [][]string{{"apple", "banana"}, {"cherry", "date"}}
		mat2 := [][]string{{"Apple", "BanAna"}, {"cHerry", "dAte"}}

		if matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to be different, but they are equal")
		}
	})
	// Test case with equal matrices of type float64
	t.Run("Equal Float Matrices", func(t *testing.T) {
		mat1 := [][]float64{{1.1, 2.2}, {3.3, 4.4}}
		mat2 := [][]float64{{1.1, 2.2}, {3.3, 4.4}}

		if !matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to be equal, but they are not")
		}
	})
	// Test case with unequal matrices of type float64
	t.Run("Different Float Matrices", func(t *testing.T) {
		mat1 := [][]float64{{1.1, 2.2}, {3.3, 4.4}}
		mat2 := [][]float64{{3.3, 2.2, 3.3}, {3.3, 4.4, 5.5}}

		if matrix.MatrixEqual(mat1, mat2) {
			t.Errorf("Expected matrices to be different, but they are equal")
		}
	})
}
