package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestCheckEqual(t *testing.T) {
	// Test case 1: Equal matrices
	mat1 := [][]int{{1, 2}, {3, 4}}
	mat2 := [][]int{{1, 2}, {3, 4}}
	expected1 := true
	result1, err1 := matrix.CheckEqual(mat1, mat2)
	if err1 != nil {
		t.Errorf("Unexpected error for test case 1: %v", err1)
	}
	if result1 != expected1 {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
	}

	// Test case 2: Matrices with different values
	mat3 := [][]int{{1, 2}, {3, 4}}
	mat4 := [][]int{{1, 2}, {3, 5}}
	expected2 := false
	result2, err2 := matrix.CheckEqual(mat3, mat4)
	if err2 != nil {
		t.Errorf("Unexpected error for test case 2: %v", err2)
	}
	if result2 != expected2 {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, result2)
	}

	// Test case 3: Matrices with different dimensions
	mat5 := [][]int{{1, 2, 3}, {4, 5, 6}}
	mat6 := [][]int{{1, 2}, {3, 4}}
	expected3 := false
	result3, err3 := matrix.CheckEqual(mat5, mat6)
	if err3 == nil {
		t.Errorf("Expected error for test case 3, but got none")
	}
	if result3 != expected3 {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected3, result3)
	}
}
