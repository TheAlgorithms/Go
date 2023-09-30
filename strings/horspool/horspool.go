// Implementation of the
// [Boyer–Moore–Horspool algorithm](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm)

package horspool

import "errors"

var ErrNotFound = errors.New("pattern was not found in the input string")

func Horspool(t, p string) (int, error) {
	pos := 0
	shiftMap := computeShiftMap(t, p)
	for pos <= len(t)-len(p) {
		j := len(p)
		for j > 0 && t[pos+j-1] == p[j-1] {
			j--
		}
		if j == 0 {
			return pos, nil
		}
		if pos+len(p) >= len(t) {
			break
		}
		pos = pos + shiftMap[t[pos+len(p)]]
	}

	return -1, ErrNotFound
}

func computeShiftMap(t, p string) (res map[uint8]int) {
	res = make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		res[t[i]] = len(p)
	}
	for i := 0; i < len(p); i++ {
		res[p[i]] = len(p) - i
	}
	return res
}
