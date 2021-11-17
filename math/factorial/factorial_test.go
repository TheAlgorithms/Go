// factorial_test.go
// description: Test for calculating factorial
// author(s) [red_byte](https://github.com/i-redbyte)
// see factorial.go

package factorial

import "testing"

func getTests() []struct {
	name   string
	n      int
	result int
} {
	var tests = []struct {
		name   string
		n      int
		result int
	}{
		{"5!", 5, 120},
		{"3!", 3, 6},
		{"6!", 6, 720},
		{"12!", 12, 479001600},
		{"1!", 1, 1},
	}
	return tests
}

func TestBruteForceFactorial(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := Iterative(tv.n)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}
}

func TestRecursiveFactorial(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := Recursive(tv.n)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}
}

func TestCalculateFactorialUseTree(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := UsingTree(tv.n)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.result, result)
			}
		})
	}
}

func BenchmarkBruteForceFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterative(10)
	}
}

func BenchmarkRecursiveFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursive(10)
	}
}

func BenchmarkCalculateFactorialUseTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursive(10)
	}
}
