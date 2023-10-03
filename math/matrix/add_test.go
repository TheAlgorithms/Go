package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestAddMatricesInt(t *testing.T) {
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	matrix2 := [][]int{
		{7, 8, 9},
		{10, 11, 12},
	}

	expectedResult := [][]int{
		{8, 10, 12},
		{14, 16, 18},
	}

	result, err := matrix.Add(matrix1, matrix2)
	if err != nil {
		t.Errorf("Error: %v", err)
	} else {
		if equal, err := matrix.CheckEqual(result, expectedResult); err != nil {
			t.Errorf("Error: %v", err)
		} else if !equal {
			t.Errorf("Result matrix does not match the expected result.")
		}
	}

	// null matrix  test
	matrix3 := [][]int{{}}

	matrix4 := [][]int{{}}

	expectedResult1 := [][]int{{}}

	result1, err1 := matrix.Add(matrix3, matrix4)
	if err1 != nil {
		t.Errorf("Error: %v", err1)
	} else {
		if equal, err := matrix.CheckEqual(result1, expectedResult1); err != nil {
			t.Errorf("Error: %v", err)
		} else if !equal {
			t.Errorf("Result matrix does not match the expected result.")
		}
	}
}

func TestAddMatricesIncompatibleDimensions(t *testing.T) {
	matrix1 := [][]int{
		{1, 2},
		{3, 4},
	}

	matrix2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	result, err := matrix.Add(matrix1, matrix2)
	if err == nil {
		t.Error("Expected an error for incompatible matrix dimensions, but got nil.")
	}

	if result != nil {
		t.Error("Result matrix should be nil when dimensions are incompatible.")
	}
}
