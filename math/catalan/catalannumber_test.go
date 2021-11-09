// catalannumber_test.go
// description: Test for returns the Catalan number
// author(s) [red_byte](https://github.com/i-redbyte)
// see catalannumber.go

package catalan

import "testing"

func TestCatalanNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero Catalan number ", 0, 1},
		{"second Catalan number ", 2, 2},
		{"third Catalan number ", 3, 5},
		{"fourth Catalan number ", 4, 14},
		{"fifth Catalan number ", 5, 42},
		{"sixth Catalan number ", 6, 132},
		{"tenth Catalan number ", 10, 16796},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := CatalanNumber(test.n); got != test.want {
				t.Errorf("CatalanNumber() = %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkCatalanNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CatalanNumber(10)
	}
}
