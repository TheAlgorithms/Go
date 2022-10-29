package math

import (
	"testing"
)

func TestSumOfDigits(t *testing.T) {
	tests := getSumOfDigitsTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := SumOfDigits(test.n); got != test.want {
				t.Errorf("SumOfDigits() = %v, want %v", got, test.want)
			}
		})
	}
}

func getSumOfDigitsTests() []struct {
	name string
	n    int64
	want int64
} {
	tests := []struct {
		name string
		n    int64
		want int64
	}{
		{"positive integer", 124294, 22},
		{"negative integer", -23057, 17},
		{"zero ", 0, 0},
		{"large integer ", 1231231231231211121, 33},
		{"max allowed int64", 9223372036854775807, 88},
	}
	return tests
}
