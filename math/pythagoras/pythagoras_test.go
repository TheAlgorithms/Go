package pythagoras

import (
	"math"
	"testing"
)

// TableDrivenTest for checking multiple values against our Test Function
var distanceTest = []struct {
	name string
	v1   Vector
	v2   Vector
	res  float64
}{
	{"random negative vector", Vector{2, -1, 7}, Vector{1, -3, 5}, 3.0},
	{"random wide vectors", Vector{4, 10, 9}, Vector{4, 3, 5}, 8.06},
	{"random wide vectors", Vector{8, 5, 5}, Vector{1, 1, 12}, 10.67},
	{"random short vectors", Vector{1, 1, 1}, Vector{2, 2, 2}, 1.73},
}

// TestDistance tests the Function Distance with 2 vectors
func TestDistance(t *testing.T) {
	t.Parallel() // marks TestDistance as capable of running in parallel with other tests
	for _, tt := range distanceTest {
		tt := tt // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := Distance(tt.v1, tt.v2)           // Calculate result for
			roundRes := (math.Floor(res*100) / 100) // Round to 2 decimal places because we can't compare an infinite number of places
			// Check result
			if roundRes != tt.res {
				t.Errorf("Distance(%v, %v) = %f, expected %f",
					tt.v1, tt.v2, roundRes, tt.res)
			}
		})
	}
}
