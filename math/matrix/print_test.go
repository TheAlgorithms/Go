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
	matrix1.Set(0, 0, 1)
	matrix1.Set(0, 1, 2)
	matrix1.Set(1, 0, 3)
	matrix1.Set(1, 1, 4)
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
	io.Copy(&buf, r)
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
	io.Copy(&buf, r)
	capturedOutput := buf.String()

	// Define the expected output
	expectedOutput := ""

	// Compare the captured output with the expected output
	if capturedOutput != expectedOutput {
		t.Errorf("Matrix.Print() produced incorrect output:\n%s\nExpected:\n%s", capturedOutput, expectedOutput)
	}
}
