// horspool_test.go
// description: Tests for horspool
// see horspool.go

package horspool

import "testing"
import "fmt"

func TestLHorspool(t *testing.T) {
	testCases := []struct {
		input    string
		pattern  string
		expected int
	}{
		{"aaaaXaaa", "X", 4},
		{"aaaaXXaa", "XX", 4},
		{"Xaaab", "X", 0},
		{"XYaab", "XY", 0},
		{"abcefghXYZ", "XYZ", 7},
		{"abcefgh€YZ⌘", "€YZ", 7},
		{"⌘bcefgh€YZ⌘", "€YZ", 7},
		{"abc", "abc", 0},
		{"", "", 0},
		{"a", "", 0},
		{"a", "a", 0},
		{"aa", "a", 0},
		{"aa", "aa", 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint("test with ", tc.input, " ", tc.pattern), func(t *testing.T) {
			result, curError := Horspool(tc.input, tc.pattern)
			if curError != nil {
				t.Fatalf("Got unexpected error")
			}
			if tc.expected != result {
				t.Fatalf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func TestLHorspoolNotExisintPattern(t *testing.T) {
	testCases := []struct {
		input   string
		pattern string
	}{
		{"", "X"},
		{"X", "Y"},
		{"X", "XX"},
		{"aaaaaaaXaXaaaa", "XXX"},
		{"aaaaaaaXaX", "XXX"},
		{"XaX", "XXX"},
		{"XaX", "XXX"},
		{"\xe2\x8c\x98", "\x98"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint("test with ", tc.input, " ", tc.pattern), func(t *testing.T) {
			result, curError := Horspool(tc.input, tc.pattern)
			if curError != ErrNotFound {
				t.Fatalf("Got unexpected error")
			}
			if result != -1 {
				t.Fatalf("expected -1, got %d", result)
			}
		})
	}
}
