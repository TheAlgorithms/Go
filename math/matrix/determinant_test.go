package matrix_test

import (
	"errors"
	"math"
	"math/rand"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

// Test different matrix contents
func TestMatrixDeterminant(t *testing.T) {
	// Find Determinant of a 2 by 2 matrix.
	matrix1, err := matrix.NewFromElements([][]int{
		{3, 8},
		{4, 6},
	})
	if err != nil {
		t.Fatalf("Error creating 3 by 3 matrix: %v", err)
	}
	determinant, err := matrix1.Determinant()
	if err != nil {
		t.Fatalf("Error returned from 3 by 3 matrix: %v", err)
	}
	if determinant != -14 {
		t.Fatalf("Determinant returned for a 3 by 3 matrix was %d; wanted -14", determinant)
	}

	// Find Dertminant of a 1 by 1 matrix
	expectedValue := rand.Intn(math.MaxInt)
	matrix2, err := matrix.NewFromElements([][]int{
		{expectedValue},
	})
	if err != nil {
		t.Fatalf("Error creating 1 by 1 matrix: %v", err)
	}
	determinant, err = matrix2.Determinant()
	if err != nil {
		t.Fatalf("Error returned from 1 by 1 matrix: %v", err)
	}
	if determinant != expectedValue {
		t.Fatalf("Determinant returned for a 1 by 1 matrix was %d; wanted %d", determinant, expectedValue)
	}

}

func TestEmptyMatrix(t *testing.T) {
	emptyElements := [][]int{}
	matrix, err := matrix.NewFromElements(emptyElements)

	if err != nil {
		t.Fatalf("Error creating Matrix with empty elements: %v", err)
	}

	determinant, err := matrix.Determinant()

	if err != nil {
		t.Fatalf("Determinant returned an error for empty matrix: %v", err)
	}

	// Check that 0 is returned from an empty matrix.
	expectedValue := 0
	if determinant != expectedValue {
		t.Errorf("Determinant returned from empty matrix was %d; wanted %d", determinant, expectedValue)
	}

}

func TestNonSquareMatrix(t *testing.T) {
	// Creating non-square matrix for testing.
	initialValue := 0
	initialRows := 4
	initialCols := 2

	nonSquareMatrix := matrix.New(initialRows, initialCols, initialValue)

	determinant, err := nonSquareMatrix.Determinant()
	// Check if non square matrix returns an error.
	if err == nil {
		t.Fatalf("No error was returned for a non-square matrix")
	}

	// Check if the correct error was returned.
	expectedError := errors.New("Matrix rows and columns must equal in order to find the determinant.")

	if err.Error() != expectedError.Error() {
		t.Errorf("Error returned from non-square matrix was \n\"%v\"; \nwanted \n\"%v\"", err, expectedError)
	}

	// Check if the determinant of the non-square matrix is 0.
	if determinant != 0 {
		t.Errorf("Determinant of non-square matrix was not 0 but was %d", determinant)
	}

}

// Test matrix returned from matrix.New
func TestDefaultMatrix(t *testing.T) {
	initialValue := 0
	initialRows := 3
	initialCols := 3
	defaultMatrix := matrix.New(initialRows, initialCols, initialValue)

	determinant, err := defaultMatrix.Determinant()

	if err != nil {
		t.Fatalf("Error finding the determinant of 3 by 3 default matrix: %v.", err)
	}
	expectedValue := 0
	if determinant != expectedValue {
		t.Errorf("Determinant of the default matrix with an initial value 0 was %d; wanted %d.", initialValue, expectedValue)
	}
}

// Benchmark a 3 by 3 matrix for computational throughput
func BenchmarkSmallMatrixDeterminant(b *testing.B) {
	// Create a 3 by 3 matrix for benchmarking
	rows := 3
	columns := 3
	initialValue := 0
	matrix := matrix.New(rows, columns, initialValue)

	for i := 0; i < b.N; i++ {
		_, _ = matrix.Determinant()
	}
}

// Benchmark a 10 by 10 matrix for computational throughput.
func BenchmarkMatrixDeterminant(b *testing.B) {
	// Create a 10 by 10 matrix for benchmarking
	rows := 10
	columns := 10
	initialValue := 0
	matrix := matrix.New(rows, columns, initialValue)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = matrix.Determinant()
	}
}
