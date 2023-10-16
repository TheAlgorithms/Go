package matrix_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMatrixString(t *testing.T) {
	// Create a sample matrix for testing
	m1, err := matrix.NewFromElements([][]int{{1, 2}, {3, 4}})
	if err != nil {
		t.Errorf("Error creating matrix: %v", err)
	}

	// Redirect stdout to capture Stringed output
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}
	os.Stdout = w

	// Call the String method
	fmt.Print(m1)

	// Reset stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Errorf("Error copying output: %v", err)
	}
	capturedOutput := buf.String()

	// Define the expected output
	expectedOutput := "1 2 \n3 4 \n"

	// Compare the captured output with the expected output
	if capturedOutput != expectedOutput {
		t.Errorf("Matrix.Print() produced incorrect output:\n%s\nExpected:\n%s", capturedOutput, expectedOutput)
	}
}

func TestNullMatrixString(t *testing.T) {

	m1 := matrix.New(0, 0, 0)
	// Redirect stdout to capture Stringed output
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}
	os.Stdout = w

	// Call the String method
	fmt.Print(m1)

	// Reset stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Errorf("Error copying output: %v", err)
	}
	capturedOutput := buf.String()

	// Define the expected output
	expectedOutput := ""

	// Compare the captured output with the expected output
	if capturedOutput != expectedOutput {
		t.Errorf("Matrix.Print() produced incorrect output:\n%s\nExpected:\n%s", capturedOutput, expectedOutput)
	}
}

func BenchmarkString(b *testing.B) {
	// Create a sample matrix for benchmarking
	rows := 100
	columns := 100
	m := matrix.New(rows, columns, 0) // Replace with appropriate values

	for i := 0; i < b.N; i++ {
		_ = m.String()
	}
}
