// This is a simple recursive implementation of [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance).
// It is not an efficient implementation, so memoisation was used for better performance.

package levenshtein

import (
	"math"
	"strings"
)

var lookUpTable map[string]int

func init() {
	lookUpTable = make(map[string]int)
}

func min(a, b, c int) int {
	return int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
}

// Distance calculates the levenshtein distance
func Distance(a, b string, caseInsensitive bool) int {
	var cost int

	lenA := len(a)
	lenB := len(b)

	if lenA == 0 {
		return lenB
	}
	if lenB == 0 {
		return lenA
	}

	if d, ok := lookUpTable[a+b]; ok {
		return d
	}

	if caseInsensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	if a[lenA-1] != b[lenB-1] {
		cost = 1
	}

	d := min(Distance(a[:lenA-1], b, caseInsensitive)+1,
		Distance(a, b[:lenB-1], caseInsensitive)+1,
		Distance(a[:lenA-1], b[:lenB-1], caseInsensitive)+cost)

	lookUpTable[a+b] = d

	return d
}
