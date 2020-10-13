package pythagoras

import (
	"math"
	"testing"
)

// TableDrivenTest for checking multiple values against our Test Function
var distanceTest = []struct {
	v1  Vector
	v2  Vector
	res float64
}{
	{Vector{2, -1, 7}, Vector{1, -3, 5}, 3.0},
	{Vector{4, 10, 9}, Vector{4, 3, 5}, 8.06},
	{Vector{8, 5, 5}, Vector{1, 1, 12}, 10.67},
	{Vector{1, 1, 1}, Vector{2, 2, 2}, 1.73},
}

//TestDistance tests the Function Distance with 2 vectors
func TestDistance(t *testing.T) {
	for _, tt := range distanceTest {
		res := Distance(tt.v1, tt.v2)           // Calculate result for
		roundRes := (math.Floor(res*100) / 100) // Round to 2 decimal places because we can't compare an infinite number of places

		if roundRes != tt.res {
			t.Errorf("Distance(%v, %v) = %f, expected %f",
				tt.v1, tt.v2, roundRes, tt.res)
		}
	}
}
