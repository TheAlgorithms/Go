package matrix_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestExample(t *testing.T) {
	matrixA := [][]int8{
		{1, 2},
		{3, 4},
	}

	matrixB := [][]int8{
		{5, 6},
		{7, 8},
	}
	expected := [][]int8{
		{19, 22},
		{43, 50},
	}

	result := matrix.StrassenMatrixMultiply(matrixA, matrixB)
	// Check if the result matches the expected result
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("Mismatch at position (%d, %d). Expected %d, but got %d.", i, j, expected[i][j], result[i][j])
			}
		}
	}
}

// TestMatrixMultiplication performs a test on matrix multiplication.
func TestMatrixMultiplication(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Generate random matrices for testing
	size := 1 << (rand.Intn(8) + 1) // tests for matrix with n as power of 2
	matrixA := makeRandomMatrix(size)
	matrixB := makeRandomMatrix(size)

	// Calculate the expected result using the standard multiplication
	expected, _ := matrix.Multiply(matrixA, matrixB)

	// Calculate the result using the Strassen algorithm
	result := matrix.StrassenMatrixMultiply(matrixA, matrixB)

	// Check if the result matches the expected result
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("Mismatch at position (%d, %d). Expected %d, but got %d.", i, j, expected[i][j], result[i][j])
			}
		}
	}
}

// Generate a random matrix of the given size
func makeRandomMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = rand.Intn(1000) // Generate random integers between 0 and 1000
		}
	}
	return matrix
}

// BenchmarkStrassenMatrixMultiply benchmarks the StrassenMatrixMultiply function.
func BenchmarkStrassenMatrixMultiply(b *testing.B) {
	size := 128 // Change the size as needed for your benchmark
	matrixA := makeRandomMatrix(size)
	matrixB := makeRandomMatrix(size)

	b.ResetTimer() // Reset the benchmark timer before the actual benchmark
	for i := 0; i < b.N; i++ {
		_ = matrix.StrassenMatrixMultiply(matrixA, matrixB)
	}
}
