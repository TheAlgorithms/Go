package matrix_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMatrixPrint(t *testing.T) {
	// Create a sample matrix for testing
	matrix1 := matrix.New(2, 2, 0)
	err := matrix1.Set(0, 0, 1)
	if err != nil {
		panic("copyMatrix.Set error: " + err.Error())
	}
	err1 := matrix1.Set(0, 1, 2)
	if err1 != nil {
		panic("copyMatrix.Set error: " + err1.Error())
	}
	err2 := matrix1.Set(1, 0, 3)
	if err2 != nil {
		panic("copyMatrix.Set error: " + err2.Error())
	}
	err3 := matrix1.Set(1, 1, 4)
	if err3 != nil {
		panic("copyMatrix.Set error: " + err3.Error())
	}
	// matrix1 := matrix.New(0, 0, 0)
	// Redirect stdout to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the Print method
	matrix1.Print()

	// Reset stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	_, err4 := io.Copy(&buf, r)
	if err4 != nil {
		panic("copyMatrix.Set error: " + err4.Error())
	}
	capturedOutput := buf.String()

	// Define the expected output
	expectedOutput := "1 2 \n3 4 \n"
	// expectedOutput := ""

	// Compare the captured output with the expected output
	if capturedOutput != expectedOutput {
		t.Errorf("Matrix.Print() produced incorrect output:\n%s\nExpected:\n%s", capturedOutput, expectedOutput)
	}
}

func TestNullMatrixPrint(t *testing.T) {

	matrix1 := matrix.New(0, 0, 0)
	// Redirect stdout to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the Print method
	matrix1.Print()

	// Reset stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		panic("copyMatrix.Set error: " + err.Error())
	}
	capturedOutput := buf.String()

	// Define the expected output
	expectedOutput := ""

	// Compare the captured output with the expected output
	if capturedOutput != expectedOutput {
		t.Errorf("Matrix.Print() produced incorrect output:\n%s\nExpected:\n%s", capturedOutput, expectedOutput)
	}
}

// BenchmarkPrint benchmarks the Print method.
func BenchmarkPrint(b *testing.B) {
	// Create a sample matrix for benchmarking
	rows := 100
	columns := 100
	matrix := matrix.New(rows, columns, 0) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		matrix.Print()
	}
}
