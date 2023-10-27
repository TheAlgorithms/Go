// Checks if a given string is a subsequence of another string.
// A subsequence of a given string is a string that can be derived from the given
// string by deleting some or no characters without changing the order of the
// remaining characters. (i.e., "dpr" is a subsequence of "depqr" while "drp" is not).
// Author: sanjibgirics

package strings

// Returns true if s is subsequence of t, otherwise return false.
func IsSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if s == t {
		return true
	}

	if len(s) == 0 {
		return true
	}

	sIndex := 0
	for tIndex := 0; tIndex < len(t); tIndex++ {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}

		if sIndex == len(s) {
			return true
		}
	}

	return false
}
