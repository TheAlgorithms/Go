package matrix_test

import (
	"errors"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestNewMatrix(t *testing.T) {

	nullMatrix := matrix.New(0, 0, 0)
	if nullMatrix == nil {
		t.Errorf("matrix.New( 0, 0, 0) returned nil, expected a matrix")
	}
	// Test creating a matrix of integers
	intMatrix := matrix.New(3, 4, 0)
	if intMatrix == nil {
		t.Errorf("matrix.New( 3, 4, 0) returned nil, expected a matrix")
	}

	// Test creating a matrix of floats
	floatMatrix := matrix.New(2, 2, 1.0)
	if floatMatrix == nil {
		t.Errorf("matrix.New(t64, 2, 2, 1.0) returned nil, expected a matrix")
	}

	// Test creating a matrix of strings
	stringMatrix := matrix.New(2, 3, "hello")
	if stringMatrix == nil {
		t.Errorf("matrix.New(ng, 2, 3, 'hello') returned nil, expected a matrix")
	}

	// Test creating a matrix of boolean values
	boolMatrix := matrix.New(2, 1, false)
	if boolMatrix == nil {
		t.Errorf("matrix.New(ng, 2, 3, 'hello') returned nil, expected a matrix")
	}

}

func TestNewFromElements(t *testing.T) {
	// Test case 1: Valid matrix
	validElements := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	expectedMatrix1 := matrix.New(2, 3, 0)
	for i := 0; i < len(validElements); i++ {
		for j := 0; j < len(validElements[0]); j++ {
			expectedMatrix1.Set(i, j, validElements[i][j])
		}
	}

	matrix1, err1 := matrix.NewFromElements(validElements)
	if err1 != nil {
		t.Errorf("NewFromElements(validElements) returned an error: %v", err1)
	}
	res, err := matrix.CheckEqual(matrix1, expectedMatrix1)
	if err != nil || res != true {
		t.Errorf("NewFromElements(validElements) returned %v, expected %v", matrix1, expectedMatrix1)
	}

	// Test case 2: Invalid matrix with different column counts
	invalidElements := [][]int{
		{1, 2, 3},
		{4, 5},
	}
	_, err2 := matrix.NewFromElements(invalidElements)
	expectedError2 := errors.New("rows have different numbers of columns")
	if err2 == nil || err2.Error() != expectedError2.Error() {
		t.Errorf("NewFromElements(invalidElements) returned error: %v, expected error: %v", err2, expectedError2)
	}

	// Test case 3: Empty matrix
	emptyElements := [][]int{}
	matrix3, err3 := matrix.NewFromElements(emptyElements)
	if err3 != nil {
		t.Errorf("NewFromElements(emptyElements) returned an error: %v", err3)
	}
	if matrix3 != nil {
		t.Errorf("NewFromElements(emptyElements) returned %v, expected nil", matrix3)
	}
}

func TestMatrixGet(t *testing.T) {
	// Create a sample matrix for testing
	matrix := matrix.New(3, 3, 0)
	matrix.Set(1, 1, 42) // Set a specific value for testing

	// Test case 1: Valid Get
	val1, err1 := matrix.Get(1, 1)
	if err1 != nil {
		t.Errorf("matrix.Get(1, 1) returned an error: %v, expected no error", err1)
	}
	if val1 != 42 {
		t.Errorf("matrix.Get(1, 1) returned %v, expected 42", val1)
	}

	// Test case 2: Get with invalid indices
	_, err2 := matrix.Get(10, 10)
	expectedError2 := errors.New("index out of range")
	if err2 == nil || err2.Error() != expectedError2.Error() {
		t.Errorf("matrix.Get(10, 10) returned error: %v, expected error: %v", err2, expectedError2)
	}
	// Test case 3: Get with invalid indices
	_, err3 := matrix.Get(-1, -3)
	expectedError3 := errors.New("index out of range")
	if err3 == nil || err3.Error() != expectedError3.Error() {
		t.Errorf("matrix.Get(10, 10) returned error: %v, expected error: %v", err3, expectedError3)
	}
}

func TestMatrixSet(t *testing.T) {
	// Create a sample matrix for testing
	matrix := matrix.New(3, 3, 0)

	// Test case 1: Valid Set
	err1 := matrix.Set(1, 1, 42)
	if err1 != nil {
		t.Errorf("matrix.Set(1, 1, 42) returned an error: %v, expected no error", err1)
	}
	val1, _ := matrix.Get(1, 1)
	if val1 != 42 {
		t.Errorf("matrix.Set(1, 1, 42) did not set the value correctly, expected 42, got %v", val1)
	}

	// Test case 2: Set with invalid indices
	err2 := matrix.Set(10, 10, 100)
	expectedError2 := errors.New("index out of bounds")
	if err2 == nil || err2.Error() != expectedError2.Error() {
		t.Errorf("matrix.Set(10, 10, 100) returned error: %v, expected error: %v", err2, expectedError2)
	}
	// Test case 3: Get with invalid indices
	err3 := matrix.Set(-13, -1, 100)
	expectedError3 := errors.New("index out of bounds")
	if err3 == nil || err3.Error() != expectedError3.Error() {
		t.Errorf("matrix.Get(10, 10) returned error: %v, expected error: %v", err3, expectedError3)
	}
}
