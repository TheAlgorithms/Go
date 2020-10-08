package main

import "testing"

func TestSurfaceAreaCube(t *testing.T) {
	res := SurfaceAreaCube(5)
	if res != 150 {
		t.Errorf("Result was incorrect, got: %f, want: %d.", res, 150)
	}
}

func TestSurfaceAreaSphere(t *testing.T) {
	res := SurfaceAreaSphere(5)
	if res != 3947.8417604357433 {
		t.Errorf("Result was incorrect, got: %f, want: %f.", res, 3947.8417604357433)
	}
}
