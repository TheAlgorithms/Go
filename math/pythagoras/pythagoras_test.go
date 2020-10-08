package main

import "testing"

//TestDistance tests the Function Distance with 2 vectors
func TestDistance(t *testing.T) {
	v1 := vector{2, -1, 7}
	v2 := vector{1, -3, 5}
	dist := Distance(v1, v2)
	if dist != 3.0 {
		t.Errorf("Distance was incorrect, got: %f, want: %f.", dist, 3.0)
	}
}
