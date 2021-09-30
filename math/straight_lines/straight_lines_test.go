package straight_lines

import (
	"testing"
)

func testDistance(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{3, 4}
	var wantedDistance float64 = 10
	var calculatedDistance float64 = distance(&p1, &p2)
	if calculatedDistance != wantedDistance {
		t.Fatalf("Failed to calculate Distance.")
	}
}

func testSection(t *testing.T) {
	p1 := Point{1, 0}
	p2 := Point{5, 0}
	wantedPoint := Point{3, 0}
	calculatedPoint := section(&p1, &p2, 1)
	if calculatedPoint != wantedPoint {
		t.Fatalf("Failed to calculate Section.")
	}
}

func testSlope(t *testing.T) {
	p1 := Point{2, 4}
	p2 := Point{1, 2}
	var wantedSlope float64 = 2
	var calculatedSlope float64 = slope(&p1, &p2)
	if calculatedSlope != wantedSlope {
		t.Fatalf("Failed to calculate Slope.")
	}
}

func testIntercept(t *testing.T) {
	p := Point{0, 3}
	var slope float64 = -5
	var wantedIntercept float64 = 3
	var calculatedIntercept float64 = intercept(&p, slope)
	if calculatedIntercept != wantedIntercept {
		t.Fatalf("Failed to calculate Intercept.")
	}
}

func testIsParallel(t *testing.T) {
	a := Point{1, 2}
	b := Point{2, 4}
	c := Point{25, 50}
	d := Point{50, 100}
	if !isParallel(&a, &b, &c, &d) {
		t.Fatalf("Failed to check if Parallel.")
	}
}

func testIsPerpendicular(t *testing.T) {
	a := Point{1, 2}
	b := Point{2, 4}
	c := Point{2, 2}
	d := Point{4, 1}
	if !isPerpendicular(&a, &b, &c, &d) {
		t.Fatalf("Failed to check if Perpendicular.")
	}
}

func testPointDistance(t *testing.T) {
	p := Point{1, 1}
	equation := [3]float64{4, 3, 1}
	var wantedDistance float64 = 1.6
	var calculatedDistance float64 = pointDistance(&p, equation)
	if calculatedDistance != wantedDistance {
		t.Fatalf("Failed to calculate Point Distance.")
	}
}
