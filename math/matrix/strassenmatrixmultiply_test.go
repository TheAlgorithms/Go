package matrix

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

// Example usage of the matrix multiplication function
func Example() {
	matrixA := [][]int{
		{1, 2},
		{3, 4},
	}

	matrixB := [][]int{
		{5, 6},
		{7, 8},
	}

	result := strassenMatrixMultiply(matrixA, matrixB)
	printMatrix(result)

	// Output:
	// [19 22]
	// [43 50]
}

// TestMatrixMultiplication performs a test on matrix multiplication.
func TestMatrixMultiplication(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Generate random matrices for testing
	size := 4
	matrixA := makeRandomMatrix(size)
	matrixB := makeRandomMatrix(size)

	// Calculate the expected result using the standard multiplication
	expected := standardMatrixMultiply(matrixA, matrixB)

	// Calculate the result using the Strassen algorithm
	result := strassenMatrixMultiply(matrixA, matrixB)

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
			matrix[i][j] = rand.Intn(10) // Generate random integers between 0 and 9
		}
	}
	return matrix
}

func TestMain(m *testing.M) {
	// Run the tests
	code := m.Run()
	os.Exit(code)
}

// BenchmarkStrassenMatrixMultiply benchmarks the strassenMatrixMultiply function.
func BenchmarkStrassenMatrixMultiply(b *testing.B) {
	size := 128 // Change the size as needed for your benchmark
	matrixA := makeRandomMatrix(size)
	matrixB := makeRandomMatrix(size)

	b.ResetTimer() // Reset the benchmark timer before the actual benchmark
	for i := 0; i < b.N; i++ {
		_ = strassenMatrixMultiply(matrixA, matrixB)
	}
}

// Print a matrix
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
