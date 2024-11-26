// powvialogarithm.go
// description: Powers in terms of logarithms
// details:
// implementation of exponentiation using exponent and logarithm, without using loops  - [Powers via logarithms wiki](https://en.wikipedia.org/wiki/Exponentiation#Powers_via_logarithms)
// time complexity: O(1)
// space complexity: O(1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see powvialogarithm_test.go

package power

import (
	"math"
)

func UsingLog(a float64, b float64) float64 {
	var p float64
	p = 1
	if a < 0 && int(b)&1 != 0 {
		p = -1
	}
	log := math.Log(math.Abs(a))
	exp := math.Exp(b * log)
	result := exp * p
	return math.Round(result)
}
