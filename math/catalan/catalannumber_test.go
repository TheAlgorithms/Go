// catalannumber_test.go
// description: Test for returns the Catalan number
// author(s) [red_byte](https://github.com/i-redbyte)
// see catalannumber.go

package catalan

import "testing"

func TestCatalanNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero Catalan number ", args{0}, 1},
		{"second Catalan number ", args{2}, 2},
		{"third Catalan number ", args{3}, 5},
		{"fourth Catalan number ", args{4}, 14},
		{"fifth Catalan number ", args{5}, 42},
		{"sixth Catalan number ", args{6}, 132},
		{"tenth Catalan number ", args{10}, 16796},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := CatalanNumber(test.args.n); got != test.want {
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
