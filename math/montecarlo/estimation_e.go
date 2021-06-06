package montecarlo

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	e = math.E
)

func estimation_e() {
	fmt.Printf("Estimating value of 'e' with %d experiment\n", INTERVAL)
	rand.Seed(time.Now().UTC().UnixNano())

	var exp int
	for i := 0; i < INTERVAL; i++ {
		var (
			sum         float64
			totalSumExp int
		)

		for sum <= 1 {
			n := rand.Float64()
			sum += n
			totalSumExp++
		}
		exp += totalSumExp
	}

	expected := float64(exp) / float64(INTERVAL)
	err := 100.0 * math.Abs(expected-e) / e
	fmt.Printf("Expected value of e : %9f\n", expected)
	fmt.Printf("e: %9f\n", e)
	fmt.Printf("Error %9f%%\n", err)
}

// func main() {
// 	estimation_e()
// }
