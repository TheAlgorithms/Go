package pythagoras

import (
	"math"
)

// Vector defines a tuple with 3 values in 3d-space
type Vector struct {
	x float64
	y float64
	z float64
}

// Distance calculates the distance between to vectors with the   Pythagoras theorem
func Distance(a, b Vector) float64 {
	res := math.Pow(b.x-a.x, 2.0) + math.Pow(b.y-a.y, 2.0) + math.Pow(b.z-a.z, 2.0)
	return math.Sqrt(res)
}
