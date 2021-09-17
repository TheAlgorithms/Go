// powvialogarithm.go
// description: Powers in terms of logarithms
// details:
// implementation of exponentiation using exponent and logarithm, without using loops  - [Powers via logarithms wiki](https://en.wikipedia.org/wiki/Exponentiation#Powers_via_logarithms)
// author(s) [red_byte](https://github.com/i-redbyte)
// see powvialogarithm_test.go

package power

import (
	"math"
)

func PowUseLog(a int, b int) int {
	p := 1
	if a < 0 && b&1 != 0 {
		p = -1
	}
	log := math.Log(math.Abs(float64(a)))
	exp := math.Exp(float64(b) * log)
	result := exp * float64(p)
	return int(math.Round(result))
}
