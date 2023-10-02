package matrix_test

import (
	"reflect"
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

	result, err := matrix.AddMatrices(matrix1, matrix2)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Result matrix does not match the expected result.")
	}
	// null matrix  test
	matrix3 := [][]int{{}}

	matrix4 := [][]int{{}}

	expectedResult1 := [][]int{{}}

	result1, err1 := matrix.AddMatrices(matrix3, matrix4)
	if err != nil {
		t.Errorf("Error: %v", err1)
	}

	if !reflect.DeepEqual(result1, expectedResult1) {
		t.Errorf("Result matrix does not match the expected result.")
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

	result, err := matrix.AddMatrices(matrix1, matrix2)
	if err == nil {
		t.Error("Expected an error for incompatible matrix dimensions, but got nil.")
	}

	if result != nil {
		t.Error("Result matrix should be nil when dimensions are incompatible.")
	}
}
