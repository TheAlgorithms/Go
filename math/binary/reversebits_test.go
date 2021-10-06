// reversebits_test.go
// description: Reverse bits
// author(s) [red_byte](https://github.com/i-redbyte)
// see reversebits.go

package binary

import "testing"

func TestReverseBits(t *testing.T) {
	var tests = []struct {
		name   string
		number uint
		result uint
	}{
		{"43261596 (00000010100101000001111010011100) to 964176192 (00111001011110000010100101000000)", 43261596, 964176192},
		{"3758096384 (11100000000000000000000000000000) to 7 (00000000000000000000000000000111)", 3758096384, 7},
		{"7 (00000000000000000000000000000111) to 3758096384 (11100000000000000000000000000000)", 7, 3758096384},
		{"2684354560 (10100000000000000000000000000000) to 5 (00000000000000000000000000000101)", 2684354560, 5},
		{"2684354561 (10100000000000000000000000000001) to 5 (10000000000000000000000000000101)", 2684354561, 2147483653},
	}
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := ReverseBits(tv.number)
			t.Log(result)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}
}

func BenchmarkReverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBits(2684354560)
	}
}
