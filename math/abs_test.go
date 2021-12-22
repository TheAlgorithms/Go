package math

import (
	"github.com/TheAlgorithms/Go/math/binary"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Abs(test.n); got != test.want {
				t.Errorf("Abs() = %v, want %v", got, test.want)
			}
		})
	}
}

func getTests() []struct {
	name string
	n    int
	want int
} {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"-1 = |1| ", -1, 1},
		{"-255 = |255| ", -255, 255},
		{"0 = |0| ", 0, 0},
		{"5 = |5| ", 5, 5},
		{"-5 = |5| ", -5, 5},
		{"-98368972 = |98368972| ", -98368972, 98368972},
	}
	return tests
}

func BenchmarkSimpleAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-1024)
	}
}
func BenchmarkBinaryAbs32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.Abs(32, -1024)
	}
}

func BenchmarkBinaryAbs64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.Abs(64, -1024)
	}
}

func BenchmarkStdLibAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Abs(-1024)
	}
}
