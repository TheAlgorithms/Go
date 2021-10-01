// arithmeticmean_test.go
// description: Test for Arithmetic mean
// author(s) [red_byte](https://github.com/i-redbyte)
// see arithmeticmean.go

package binary

import "testing"

func TestMeanUsingAndXor(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := MeanUsingAndXor(tv.a, tv.b)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}

}

func TestMeanUsingRightShift(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := MeanUsingRightShift(tv.a, tv.b)
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
		{"Average of 2 and 4", 2, 4, 3},
		{"Average of 5 and 5", 5, 5, 5},
		{"Average of 1000 and 1002", 1000, 1002, 1001},
		{"Average of 80 and 40", 80, 40, 60},
		{"Average of 7 and 4", 7, 4, 5},
	}
	return tests
}
func BenchmarkMeanUsingAndXor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MeanUsingAndXor(222, 888)
	}
}

func BenchmarkMeanUsingRightShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MeanUsingRightShift(222, 888)
	}
}
