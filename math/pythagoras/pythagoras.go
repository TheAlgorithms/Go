package main

import (
	"fmt"

	"math"
)

// vector defines a tuple with 3 values in 3d-space
type vector struct {
	x int
	y int
	z int
}

//Distance calculates the distance between to vectors with the   Pythagoras theorem
func Distance(a, b vector) float64 {
	abs := float64(math.Abs(float64(b.x - a.x)))
	res := float64(math.Pow(abs, 2.0) + math.Pow(float64(b.y-a.y), 2.0) + math.Pow(float64(b.z-a.z), 2))
	fmt.Println(abs, res)
	return math.Sqrt(res)
}

func main() {

}
