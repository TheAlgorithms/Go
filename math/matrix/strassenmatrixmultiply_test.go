package matrix_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestStrassenMatrixMultiply(t *testing.T) {
	// Create two sample matrices
	dataA := [][]int{{1, 2}, {4, 5}}
	dataB := [][]int{{9, 8}, {6, 5}}
	matrixA, err := matrix.NewFromElements(dataA)
	if err != nil {
		t.Error("copyMatrix.Set error: " + err.Error())
	}
	matrixB, err := matrix.NewFromElements(dataB)
	if err != nil {
		t.Error("copyMatrix.Set error: " + err.Error())
	}

	// Perform matrix multiplication using Strassen's algorithm
	resultMatrix, err := matrixA.StrassenMatrixMultiply(matrixB)
	if err != nil {
		t.Error("copyMatrix.Set error: " + err.Error())
	}

	// Expected result
	expectedData, err := matrixA.Multiply(matrixB)
	if err != nil {
		t.Error("copyMatrix.Set error: " + err.Error())
	}

	// Check the dimensions of the result matrix
	expectedRows := expectedData.Rows()
	expectedColumns := expectedData.Columns()
	rows := resultMatrix.Rows()
	columns := resultMatrix.Columns()

	if rows != expectedRows {
		t.Errorf("Expected %d rows in result matrix, but got %d", expectedRows, rows)
	}

	if columns != expectedColumns {
		t.Errorf("Expected %d columns in result matrix, but got %d", expectedColumns, columns)
	}

	// Check the values in the result matrix
	for i := 0; i < expectedRows; i++ {
		for j := 0; j < expectedColumns; j++ {
			val, err := resultMatrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			expVal, err := expectedData.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			if val != expVal {
				t.Errorf("Expected value %d at (%d, %d) in result matrix, but got %d", expVal, i, j, val)
			}
		}
	}
}
func TestMatrixMultiplication(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate random matrices for testing
	size := 1 << (rand.Intn(8) + 1) // tests for matrix with n as power of 2
	matrixA := MakeRandomMatrix[int](size, size)
	matrixB := MakeRandomMatrix[int](size, size)

	// Calculate the expected result using the standard multiplication
	expected, err := matrixA.Multiply(matrixB)
	if err != nil {
		t.Error("copyMatrix.Set error: " + err.Error())
	}
	// Calculate the result using the Strassen algorithm
	result, err := matrixA.StrassenMatrixMultiply(matrixB)
	if err != nil {
		t.Error("copyMatrix.Set error: " + err.Error())
	}

	// Check if the result matches the expected result
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			val, err := result.Get(i, j)
			if err != nil {
				t.Error("copyMatrix.Set error: " + err.Error())
			}
			exp, err := expected.Get(i, j)
			if err != nil {
				t.Error("copyMatrix.Set error: " + err.Error())
			}
			if val != exp {
				t.Errorf("Mismatch at position (%d, %d). Expected %d, but got %d.", i, j, exp, val)
			}
		}
	}
}

func MakeRandomMatrix[T constraints.Integer](rows, columns int) matrix.Matrix[T] {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	matrixData := make([][]T, rows)
	for i := 0; i < rows; i++ {
		matrixData[i] = make([]T, columns)
		for j := 0; j < columns; j++ {
			matrixData[i][j] = T(rand.Intn(1000)) // Generate random integers between 0 and 1000
		}
	}

	randomMatrix, _ := matrix.NewFromElements(matrixData)
	return randomMatrix
}

// BenchmarkStrassenMatrixMultiply benchmarks the StrassenMatrixMultiply function.
func BenchmarkStrassenMatrixMultiply(b *testing.B) {
	// Create sample matrices for benchmarking
	rows := 64 // it is large enough for multiplication
	columns := 64
	m1 := matrix.New(rows, columns, 2) // Replace with appropriate values
	m2 := matrix.New(rows, columns, 3) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_, _ = m1.StrassenMatrixMultiply(m2)
	}
}
