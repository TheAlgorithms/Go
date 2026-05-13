package dynamic

import "testing"

func TestLongestCommonSubstring(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected int
	}{
		{"basic match", "abcde", "abfce", 2},          // "ab"
		{"full match", "abcdef", "abcdef", 6},          // entire string
		{"no match", "abc", "xyz", 0},
		{"empty first", "", "abc", 0},
		{"empty second", "abc", "", 0},
		{"both empty", "", "", 0},
		{"single char match", "a", "a", 1},
		{"single char no match", "a", "b", 0},
		{"substring at end", "abcdef", "cdef", 4},      // "cdef"
		{"substring at start", "abcdef", "abcd", 4},    // "abcd"
		{"repeated chars", "aabbbcccc", "bbcc", 4},     // "bbcc"
		{"overlapping", "geeksforgeeks", "practisegeeks", 5}, // "geeks"
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LongestCommonSubstring(test.str1, test.str2)
			if result != test.expected {
				t.Errorf("LongestCommonSubstring(%q, %q) = %d, expected %d",
					test.str1, test.str2, result, test.expected)
			}
		})
	}
}
