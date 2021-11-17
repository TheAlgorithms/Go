// bitcounter_test.go
// description: Test for counts the number of set bits in a number
// author(s) [red_byte](https://github.com/i-redbyte)
// see bitcounter.go

package binary

import "testing"

func TestBitCounter(t *testing.T) {
	tests := []struct {
		name   string
		number uint
		want   int
	}{
		{"Number of bits in a number 0", 0, 0},
		{"Number of bits in a number 1", 1, 1},
		{"Number of bits in a number 255", 255, 8},
		{"Number of bits in a number 7", 7, 3},
		{"Number of bits in a number 8", 8, 1},
		{"Number of bits in a number 9223372036854775807", 9223372036854775807, 63},
		{"Number of bits in a number 2147483647", 2147483647, 31},
		{"Number of bits in a number 15", 15, 4},
		{"Number of bits in a number 16", 16, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := BitCounter(test.number); got != test.want {
				t.Errorf("BitCounter() = %v, want %v", got, test.want)
			}
		})
	}
}
