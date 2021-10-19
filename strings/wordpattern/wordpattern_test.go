// wordpattern_test.go
// description: Word Pattern
// author(s) [red_byte](https://github.com/i-redbyte)
// see wordpattern.go

package main

import "testing"

func getTests() []struct {
	s       string
	pattern string
	result  bool
} {
	var tests = []struct {
		s       string
		pattern string
		result  bool
	}{
		{"Go kotlin kotlin Go", "xyyx", true},
		{"Go kotlin kotlin Go", "xxyy", false},
		{"C C++ Go", "abc", true},
		{"C Go Go", "abc", false},
	}
	return tests
}

func TestWordPattern(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.s+" with pattern "+tv.pattern, func(t *testing.T) {
			result := WordPattern(tv.s, tv.pattern)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", tv.result, result)
			}
		})
	}
}

func BenchmarkMeanUsingRightShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordPattern("Go kotlin kotlin Go", "xyyx")
	}
}
