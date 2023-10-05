package matrix_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestCheckEqual(t *testing.T) {
	// Create two matrices with the same dimensions and equal values
	matrix1 := matrix.New(2, 2, 0)
	matrix2 := matrix.New(2, 2, 0)

	// Test case 1: Matrices are equal
	equal, err := matrix.CheckEqual(matrix1, matrix2)
	if err != nil {
		t.Errorf("CheckEqual(matrix1, matrix2) returned an error: %v, expected no error", err)
	}
	if !equal {
		t.Errorf("CheckEqual(matrix1, matrix2) returned false, expected true (matrices are equal)")
	}

	// Create two matrices with the same dimensions but different values
	matrix3 := matrix.New(2, 2, 1)
	matrix4 := matrix.New(2, 2, 0)

	// Test case 2: Matrices are not equal
	equal2, err2 := matrix.CheckEqual(matrix3, matrix4)
	if err2 != nil {
		t.Errorf("CheckEqual(matrix3, matrix4) returned an error: %v, expected no error", err2)
	}
	if equal2 {
		t.Errorf("CheckEqual(matrix3, matrix4) returned true, expected false (matrices are not equal)")
	}

	// Create two matrices with different dimensions
	matrix5 := matrix.New(2, 2, 0)
	matrix6 := matrix.New(2, 3, 0)

	// Test case 3: Matrices have different dimensions
	_, err3 := matrix.CheckEqual(matrix5, matrix6)
	expectedError3 := fmt.Errorf("invalid matrices: %v", errors.New("matrices have different dimensions: rows=2 vs columns=2"))
	if err3 == nil || err3.Error() != expectedError3.Error() {
		t.Errorf("CheckEqual(matrix5, matrix6) returned error: %v, expected error: %v", err3, expectedError3)
	}
}

func TestCheckEqual_ErrorHandling(t *testing.T) {
	// Create two matrices with different dimensions
	matrix1 := matrix.New(2, 2, 0)
	matrix2 := matrix.New(3, 3, 0)

	// Test case: Invalid matrices, expect an error
	_, err := matrix.CheckEqual(matrix1, matrix2)
	expectedError := fmt.Errorf("invalid matrices: %v", errors.New("matrices have different dimensions: rows=2 vs columns=2"))
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("CheckEqual(matrix1, matrix2) returned error: %v, expected error: %v", err, expectedError)
	}
}

func BenchmarkCheckEqualSmallMatrix(b *testing.B) {
	matrix1 := matrix.New(10, 10, 0) // Create a 10x10 matrix with all zeros
	matrix2 := matrix.New(10, 10, 0) // Create another 10x10 matrix with all zeros

	for i := 0; i < b.N; i++ {
		_, _ = matrix.CheckEqual(matrix1, matrix2)
	}
}

func BenchmarkCheckEqualLargeMatrix(b *testing.B) {
	size := 1000 // Choose an appropriate size for your large matrix
	matrix1 := MakeRandomMatrix[int](size, size)
	matrix2 := MakeRandomMatrix[int](size, size)

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		_, _ = matrix.CheckEqual(matrix1, matrix2)
	}
}
