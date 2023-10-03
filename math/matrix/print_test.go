package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestPrint(t *testing.T) {
	matrix1 := [][]int{
		nil,
	}

	matrix2 := [][]int{
		{7, 8, 9},
		{10, 11, 12},
	}
	matrix3 := [][]bool{
		{true, false},
		{false, false},
		{true, false},
	}
	matrix4 := [][]string{
		{"apple", "banana", "cherry"},
		{"dog", "elephant", "fox"},
	}

	matrix5 := [][]float32{
		{12.3, 12.2},
	}

	matrix.Print(matrix1)
	matrix.Print(matrix2)
	matrix.Print(matrix3)
	matrix.Print(matrix4)
	matrix.Print(matrix5)
}
