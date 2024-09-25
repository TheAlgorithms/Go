package matrix_test

import (
	"errors"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestNewMatrix(t *testing.T) {

	nullMatrix := matrix.New(0, 0, 0)
	if nullMatrix.Rows() != 0 || nullMatrix.Columns() != 0 {
		t.Errorf("matrix.New( 0, 0, 0) returned nil, expected a matrix")
	}
	// Test creating a matrix of integers
	intMatrix := matrix.New(3, 4, 0)
	if intMatrix.Rows() != 3 || intMatrix.Columns() != 4 {
		t.Errorf("matrix.New( 3, 4, 0) returned nil, expected a matrix")
	}

}

func TestNewFromElements(t *testing.T) {
	// Test case 1: Valid matrix
	validElements := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	expectedm1 := matrix.New(2, 3, 0)
	for i := 0; i < len(validElements); i++ {
		for j := 0; j < len(validElements[0]); j++ {
			err := expectedm1.Set(i, j, validElements[i][j])
			if err != nil {
				t.Errorf("copyMatrix.Set error: %s", err.Error())
			}
		}
	}

	m1, err1 := matrix.NewFromElements(validElements)
	if err1 != nil {
		t.Errorf("NewFromElements(validElements) returned an error: %v", err1)
	}
	res := m1.CheckEqual(expectedm1)
	if res != true {
		t.Errorf("NewFromElements(validElements) returned %v, expected %v", m1, expectedm1)
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
	m3, err3 := matrix.NewFromElements(emptyElements)
	if err3 != nil {
		t.Errorf("NewFromElements(emptyElements) returned an error: %v", err3)
	}
	if m3.Rows() != 0 || m3.Columns() != 0 {
		t.Errorf("NewFromElements(emptyElements) returned %v, expected nil", m3)
	}
}

func TestMatrixGet(t *testing.T) {
	// Create a sample matrix for testing
	matrix := matrix.New(3, 3, 0)
	err := matrix.Set(1, 1, 42) // Set a specific value for testing
	if err != nil {
		t.Errorf("copyMatrix.Set error: %s", err.Error())
	}
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
	val1, err := matrix.Get(1, 1)
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}
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

func TestMatrixRows(t *testing.T) {
	// Create a sample matrix
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix, err := matrix.NewFromElements(data)
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}
	// Check the number of rows
	expectedRows := len(data)
	rows := matrix.Rows()
	if rows != expectedRows {
		t.Errorf("Expected %d rows, but got %d", expectedRows, rows)
	}
}

func TestMatrixColumns(t *testing.T) {
	// Create a sample matrix
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix, err := matrix.NewFromElements(data)
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}
	// Check the number of columns
	expectedColumns := len(data[0])
	columns := matrix.Columns()
	if columns != expectedColumns {
		t.Errorf("Expected %d columns, but got %d", expectedColumns, columns)
	}
}

func TestMatrixEmptyRowsAndColumns(t *testing.T) {
	// Create an empty matrix
	emptyMatrix := matrix.New(0, 0, 0)

	// Check the number of rows and columns for an empty matrix
	rows := emptyMatrix.Rows()
	columns := emptyMatrix.Columns()

	if rows != 0 {
		t.Errorf("Expected 0 rows for an empty matrix, but got %d", rows)
	}

	if columns != 0 {
		t.Errorf("Expected 0 columns for an empty matrix, but got %d", columns)
	}
}

// BenchmarkNew benchmarks the New function.
func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = matrix.New(100, 100, 0) // Change the arguments to match your use case
	}
}

// BenchmarkNewFromElements benchmarks the NewFromElements function.
func BenchmarkNewFromElements(b *testing.B) {
	// Create a sample matrix for benchmarking
	rows := 100
	columns := 100
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make([]int, columns)
		for j := range elements[i] {
			elements[i][j] = i*columns + j // Some arbitrary values
		}
	}

	for i := 0; i < b.N; i++ {
		_, _ = matrix.NewFromElements(elements)
	}
}

// BenchmarkGet benchmarks the Get method.
func BenchmarkGet(b *testing.B) {
	// Create a sample matrix for benchmarking
	rows := 100
	columns := 100
	matrix := matrix.New(rows, columns, 0)

	for i := 0; i < b.N; i++ {
		_, _ = matrix.Get(50, 50) // Change the row and column indices as needed
	}
}

// BenchmarkSet benchmarks the Set method.
func BenchmarkSet(b *testing.B) {
	// Create a sample matrix for benchmarking
	rows := 100
	columns := 100
	matrix := matrix.New(rows, columns, 0)

	for i := 0; i < b.N; i++ {
		_ = matrix.Set(50, 50, 42) // Change the row, column, and value as needed
	}
}

// BenchmarkRows benchmarks the Rows method.
func BenchmarkRows(b *testing.B) {
	// Create a sample matrix for benchmarking
	rows := 100
	columns := 100
	matrix := matrix.New(rows, columns, 0)

	for i := 0; i < b.N; i++ {
		_ = matrix.Rows()
	}
}

// BenchmarkColumns benchmarks the Columns method.
func BenchmarkColumns(b *testing.B) {
	// Create a sample matrix for benchmarking
	rows := 100
	columns := 100
	matrix := matrix.New(rows, columns, 0)

	for i := 0; i < b.N; i++ {
		_ = matrix.Columns()
	}
}
