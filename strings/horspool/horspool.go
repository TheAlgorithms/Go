// Implementation of the
// [Boyer–Moore–Horspool algorithm](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm)

package horspool

import "errors"

var ErrNotFound = errors.New("pattern was not found in the input string")

func Horspool(t, p string) (int, error) {
	// in order to handle multy-byte character properly
	// the input is converted into rune arrays
	return horspool([]rune(t), []rune(p))
}

func horspool(t, p []rune) (int, error) {
	shiftMap := computeShiftMap(t, p)
	pos := 0
	for pos <= len(t)-len(p) {
		if isMatch(pos, t, p) {
			return pos, nil
		}
		if pos+len(p) >= len(t) {
			// because the remaining length of the input string
			// is the same as the length of the pattern
			// and it does not match the pattern
			// it is impossible to find the pattern
			break
		}

		// because of the check above
		// t[pos+len(p)] is defined
		pos += shiftMap[t[pos+len(p)]]
	}

	return -1, ErrNotFound
}

// Checks if the array p matches the subarray of t starting at pos.
// Note that backward iteration.
// There are [other](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm#Tuning_the_comparison_loop)
// approaches possible.
func isMatch(pos int, t, p []rune) bool {
	j := len(p)
	for j > 0 && t[pos+j-1] == p[j-1] {
		j--
	}
	return j == 0
}

func computeShiftMap(t, p []rune) (res map[rune]int) {
	res = make(map[rune]int)
	for _, tCode := range t {
		res[tCode] = len(p)
	}
	for i, pCode := range p {
		res[pCode] = len(p) - i
	}
	return res
}
