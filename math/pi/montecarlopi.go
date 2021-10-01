// montecarlopi.go
// description: Calculating pi by the Monte Carlo method
// details:
// implementation of Monte Carlo Algorithm for the calculating of Pi - [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method)
// author(s) [red_byte](https://github.com/i-redbyte)
// see montecarlopi_test.go

package pi

import (
	"math/rand"
	"time"
)

func MonteCarloPi(randomPoints int) float64 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	inside := 0
	for i := 0; i < randomPoints; i++ {
		x := rnd.Float64()
		y := rnd.Float64()
		if x*x+y*y <= 1 {
			inside += 1
		}
	}
	pi := float64(inside) / float64(randomPoints) * 4
	return pi
}
