package dynamic

import "testing"

func TestWordBreak(t *testing.T) {
	testCases := map[string]struct {
		s        string
		wordDict []string
		want     bool
	}{
		"empty string and empty dict":      {"", []string{}, true},
		"empty string with non-empty dict": {"", []string{"a", "b"}, true},
		"single word match":                {"apple", []string{"apple"}, true},
		"single word no match":             {"apple", []string{"orange"}, false},
		"two words forming a string":       {"leetcode", []string{"leet", "code"}, true},
		"two words but no match":           {"applepen", []string{"apple", "pen"}, true},
		"three words forming a string":     {"applepenapple", []string{"apple", "pen"}, true},
		"complex case":                     {"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, true},
		"no possible segmentation":         {"abcdef", []string{"ab", "cd", "def", "abcd"}, false},
		"dict with overlapping words":      {"pineapplepenapple", []string{"pine", "apple", "pen"}, true},
		"single character match":           {"a", []string{"a"}, true},
		"single character no match":        {"b", []string{"a"}, false},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := WordBreak(test.s, test.wordDict); got != test.want {
				t.Errorf("wordBreak(%q, %v) = %v, want %v", test.s, test.wordDict, got, test.want)
			}
		})
	}
}
