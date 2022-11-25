// hexagonalnumber.go
// description: Returns nth hexagonal number
// details:
//  Hexagonal number: The nth hexagonal number hn is the
//  number of distinct dots in a pattern of dots consisting
//  of the outlines of regular hexagons with sides up to n dots,
//  when the hexagons are overlaid so that they share one vertex.
// wikipedia: https://en.wikipedia.org/wiki/Hexagonal_number
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see hexagonalnumber_test.go

package math

// HexagonalNumber returns nth hexagonal number.
func HexagonalNumber(n int) (int, error) {
	if n < 0 {
		return 0, ErrPosArgsOnly
	}
	if n == 0 {
		return 0, ErrNonZeroArgsOnly
	}
	return n * (2*n - 1), nil
}
