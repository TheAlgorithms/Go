// longestpalindrome_test.go
// description: Manacher's algorithm (Longest palindromic substring)
// author(s) [red_byte](https://github.com/i-redbyte)
// see longestpalindrome.go

package manacher

import "testing"

func getTests() []struct {
	s      string
	result string
} {
	var tests = []struct {
		s      string
		result string
	}{
		{"olokazakabba", "kazak"},
		{"abaacakkkkk", "kkkkk"},
		{"qqqq C++ groovy mom pooop", "pooop"},
		{"CCCPCCC @@@ hello", "CCCPCCC"},
		{"goog gogogogogog -go- gogogogo ", "gogogogogog"},
	}
	return tests
}

func TestLongestPalindrome(t *testing.T) {
	tests := getTests()
	for _, tv := range tests {
		t.Run(tv.s, func(t *testing.T) {
			result := LongestPalindrome(tv.s)
			t.Log(tv.s, " ", result)
			if result != tv.result {
				t.Errorf("Wrong result! Expected:%s, returned:%s ", tv.result, result)
			}
		})
	}
}
