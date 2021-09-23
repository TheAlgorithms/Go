// arithmeticmean_test.go
// description: Test for Arithmetic mean
// author(s) [red_byte](https://github.com/i-redbyte)
// see arithmeticmean.go
package binarymath

import "testing"

func TestArithmeticMean1(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := ArithmeticMean1(tv.a, tv.b)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}

}

func TestArithmeticMean2(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := ArithmeticMean2(tv.a, tv.b)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}

}
func getTests() []struct {
	name   string
	a      int
	b      int
	result int
} {
	var tests = []struct {
		name   string
		a      int
		b      int
		result int
	}{
		{"(2+4)/2=3", 2, 4, 3},
		{"(5+5)/2=5", 5, 5, 5},
		{"(1000+1002)/2=1000", 1000, 1002, 1001},
		{"(80+40)/2=60", 80, 40, 60},
		{"(7+4)/2=5", 7, 4, 5},
	}
	return tests
}
func BenchmarkArithmeticMean1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArithmeticMean1(222, 888)
	}
}

func BenchmarkArithmeticMean2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArithmeticMean2(222, 888)
	}
}
