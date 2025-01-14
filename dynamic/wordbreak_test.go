package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseWordBreak struct {
	s        string
	wordDict []string
	expected bool
}

func getWordBreakTestCases() []testCaseWordBreak {
	return []testCaseWordBreak{
		{"leetcode", []string{"leet", "code"}, true},                        // "leetcode" can be segmented into "leet" and "code"
		{"applepenapple", []string{"apple", "pen"}, true},                   // "applepenapple" can be segmented into "apple", "pen", "apple"
		{"catsanddog", []string{"cats", "dog", "sand", "and", "cat"}, true}, // "catsanddog" can be segmented into "cats", "and", "dog"
		{"bb", []string{"a", "b", "bbb", "aaaa", "aaa"}, true},              // "bb" can be segmented into "b" and "b"
		{"", []string{"cat", "dog", "sand", "and"}, true},                   // Empty string can always be segmented (empty words)
		{"applepie", []string{"apple", "pie"}, true},                        // "applepie" can be segmented into "apple" and "pie"
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false}, // "catsandog" cannot be segmented
		{"ilovecoding", []string{"i", "love", "coding"}, true},              // "ilovecoding" can be segmented into "i", "love", "coding"
		{"cars", []string{"car", "ca", "rs"}, true},                         // "cars" can be segmented into "car" and "s"
		{"pen", []string{"pen", "pencil"}, true},                            // "pen" is a direct match
		{"apple", []string{"orange", "banana"}, false},                      // "apple" is not in the word dictionary
	}
}

func TestWordBreak(t *testing.T) {
	t.Run("Word Break test cases", func(t *testing.T) {
		for _, tc := range getWordBreakTestCases() {
			actual := dynamic.WordBreak(tc.s, tc.wordDict)
			if actual != tc.expected {
				t.Errorf("WordBreak(%q, %v) = %v; expected %v", tc.s, tc.wordDict, actual, tc.expected)
			}
		}
	})
}
