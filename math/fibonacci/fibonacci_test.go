package fibonacci

import (
	"github.com/TheAlgorithms/Go/dynamic"
	"testing"
)

func getTests() []struct {
	name string
	n    uint
	want uint
} {
	tests := []struct {
		name string
		n    uint
		want uint
	}{
		{"Fibonacci 0-th number == 0", 0, 0},
		{"Fibonacci 1-th number == 1", 1, 1},
		{"Fibonacci 2-th number == 1", 2, 1},
		{"Fibonacci 3-th number == 2", 3, 2},
		{"Fibonacci 4-th number == 3", 4, 3},
		{"Fibonacci 5-th number == 5", 5, 5},
		{"Fibonacci 6-th number == 8", 6, 8},
		{"Fibonacci 7-th number == 13", 7, 13},
		{"Fibonacci 8-th number == 21", 8, 21},
		{"Fibonacci 9-th number == 34", 9, 34},
		{"Fibonacci 10-th number == 55", 10, 55},
		{"Fibonacci 90-th number == 2880067194370816120", 90, 2880067194370816120},
	}
	return tests
}

func TestMatrix(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Matrix(test.n); got != test.want {
				t.Errorf("Return value = %v, want %v", got, test.want)
			}
		})
	}
}

func TestNthFibonacci(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := dynamic.NthFibonacci(test.n); got != test.want {
				t.Errorf("Return value = %v, want %v", got, test.want)
			}
		})
	}
}

func TestFormula(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Formula(test.n); test.n <= 10 && got != test.want {
				t.Errorf("Return value = %v, want %v", got, test.want)
			}
		})
	}
}

func TestRecursive(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		if test.n <= 10 {
			t.Run(test.name, func(t *testing.T) {
				if got := Recursive(test.n); got != test.want {
					t.Errorf("Return value = %v, want %v", got, test.want)
				}
			})
		}
	}
}

func BenchmarkNthFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dynamic.NthFibonacci(90)
	}
}

func BenchmarkMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Matrix(90)
	}
}

func BenchmarkFormula(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Formula(90)
	}
}
