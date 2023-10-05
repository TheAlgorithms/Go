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
	matrixA, _ := matrix.NewFromElements(dataA)
	matrixB, _ := matrix.NewFromElements(dataB)

	// Perform matrix multiplication using Strassen's algorithm
	resultMatrix := matrix.StrassenMatrixMultiply(matrixA, matrixB)

	// Expected result
	expectedData, _ := matrix.Multiply(matrixA, matrixB)

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
			val, _ := resultMatrix.Get(i, j)
			expVal, _ := expectedData.Get(i, j)
			if val != expVal {
				t.Errorf("Expected value %d at (%d, %d) in result matrix, but got %d", expVal, i, j, val)
			}
		}
	}
}
func TestMatrixMultiplication(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Generate random matrices for testing
	size := 1 << (rand.Intn(8) + 1) // tests for matrix with n as power of 2
	matrixA := MakeRandomMatrix[int](size, size)
	matrixB := MakeRandomMatrix[int](size, size)

	// Calculate the expected result using the standard multiplication
	expected, _ := matrix.Multiply(matrixA, matrixB)

	// Calculate the result using the Strassen algorithm
	result := matrix.StrassenMatrixMultiply(matrixA, matrixB)

	// Check if the result matches the expected result
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			val, _ := result.Get(i, j)
			exp, _ := expected.Get(i, j)
			if val != exp {
				t.Errorf("Mismatch at position (%d, %d). Expected %d, but got %d.", i, j, exp, val)
			}
		}
	}
}

func MakeRandomMatrix[T constraints.Integer](rows, columns int) matrix.Matrix[T] {
	rand.Seed(time.Now().UnixNano())

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
	matrix1 := matrix.New(rows, columns, 2) // Replace with appropriate values
	matrix2 := matrix.New(rows, columns, 3) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_ = matrix.StrassenMatrixMultiply(matrix1, matrix2)
	}
}
