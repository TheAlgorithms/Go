// montecarlopi.go
// description: Calculating pi by the Monte Carlo method
// details:
// implementations of Monte Carlo Algorithm for the calculating of Pi - [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method)
// time complexity: O(n)
// space complexity: O(1)
// author(s): [red_byte](https://github.com/i-redbyte), [Paul Leydier] (https://github.com/paul-leydier)
// see montecarlopi_test.go

package pi

import (
	"fmt"       // Used for error formatting
	"math/rand" // Used for random number generation in Monte Carlo method
	"runtime"   // Used to get information on available CPUs
	"time"      // Used for seeding the random number generation
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

// MonteCarloPiConcurrent approximates the value of pi using the Monte Carlo method.
// Unlike the MonteCarloPi function (first version), this implementation uses
// goroutines and channels to parallelize the computation.
// More details on the Monte Carlo method available at https://en.wikipedia.org/wiki/Monte_Carlo_method.
// More details on goroutines parallelization available at https://go.dev/doc/effective_go#parallel.
func MonteCarloPiConcurrent(n int) (float64, error) {
	numCPU := runtime.GOMAXPROCS(0)
	c := make(chan int, numCPU)
	pointsToDraw, err := splitInt(n, numCPU) // split the task in sub-tasks of approximately equal sizes
	if err != nil {
		return 0, err
	}

	// launch numCPU parallel tasks
	for _, p := range pointsToDraw {
		go drawPoints(p, c)
	}

	// collect the tasks results
	inside := 0
	for i := 0; i < numCPU; i++ {
		inside += <-c
	}
	return float64(inside) / float64(n) * 4, nil
}

// drawPoints draws n random two-dimensional points in the interval [0, 1), [0, 1) and sends through c
// the number of points which where within the circle of center 0 and radius 1 (unit circle)
func drawPoints(n int, c chan<- int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	inside := 0
	for i := 0; i < n; i++ {
		x, y := rnd.Float64(), rnd.Float64()
		if x*x+y*y <= 1 {
			inside++
		}
	}
	c <- inside
}

// splitInt takes an integer x and splits it within an integer slice of length n in the most uniform
// way possible.
// For example, splitInt(10, 3) will return []int{4, 3, 3}, nil
func splitInt(x int, n int) ([]int, error) {
	if x < n {
		return nil, fmt.Errorf("x must be < n - given values are x=%d, n=%d", x, n)
	}
	split := make([]int, n)
	if x%n == 0 {
		for i := 0; i < n; i++ {
			split[i] = x / n
		}
	} else {
		limit := x % n
		for i := 0; i < limit; i++ {
			split[i] = x/n + 1
		}
		for i := limit; i < n; i++ {
			split[i] = x / n
		}
	}
	return split, nil
}
