package abs

import (
	"github.com/TheAlgorithms/Go/math/binary"
	"math"
	"testing"
)

func TestABS(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := ABS(test.n); got != test.want {
				t.Errorf("ABS() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestBinaryABS(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := binary.ABS(64, test.n); got != test.want {
				t.Errorf("ABS() = %v, want %v", got, test.want)
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
		{"-98368972 = |5| ", -98368972, 98368972},
	}
	return tests
}

func BenchmarkSimpleABS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ABS(-1024)
	}
}
func BenchmarkBinaryABS32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.ABS(32, -1024)
	}
}

func BenchmarkBinaryABS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.ABS(64, -1024)
	}
}

func BenchmarkStdLibABS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Abs(-1024)
	}
}
