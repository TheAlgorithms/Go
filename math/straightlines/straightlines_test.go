package straightlines

import (
	"testing"
)

func TestDistance(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{3, 4}
	var wantedDistance float64 = 5
	var calculatedDistance float64 = distance(&p1, &p2)
	if calculatedDistance != wantedDistance {
		t.Fatalf("Failed to calculate Distance.")
	}
}

func TestSection(t *testing.T) {
	p1 := Point{1, 0}
	p2 := Point{5, 0}
	wantedPoint := Point{3, 0}
	calculatedPoint := section(&p1, &p2, 1)
	if calculatedPoint != wantedPoint {
		t.Fatalf("Failed to calculate Section.")
	}
}

func TestSlope(t *testing.T) {
	line := Line{P1: Point{1, 2}, P2: Point{2, 4}}
	var wantedSlope float64 = 2
	var calculatedSlope float64 = slope(&line)
	if calculatedSlope != wantedSlope {
		t.Fatalf("Failed to calculate Slope.")
	}
}

func TestIntercept(t *testing.T) {
	p := Point{0, 3}
	var slope float64 = -5
	var wantedIntercept float64 = 3
	var calculatedIntercept float64 = intercept(&p, slope)
	if calculatedIntercept != wantedIntercept {
		t.Fatalf("Failed to calculate Intercept.")
	}
}

func TestIsParallel(t *testing.T) {
	l1 := Line{P1: Point{1, 2}, P2: Point{2, 4}}
	l2 := Line{P1: Point{25, 50}, P2: Point{50, 100}}
	if !isParallel(&l1, &l2) {
		t.Fatalf("Failed to check if Parallel.")
	}
}

func TestIsPerpendicular(t *testing.T) {
	l1 := Line{P1: Point{1, 2}, P2: Point{2, 4}}
	l2 := Line{P1: Point{2, 2}, P2: Point{4, 1}}
	if !isPerpendicular(&l1, &l2) {
		t.Fatalf("Failed to check if Perpendicular.")
	}
}

func TestPointDistance(t *testing.T) {
	p := Point{1, 1}
	equation := [3]float64{4, 3, 1}
	var wantedDistance float64 = 1.6
	var calculatedDistance float64 = pointDistance(&p, equation)
	if calculatedDistance != wantedDistance {
		t.Fatalf("Failed to calculate Point Distance.")
	}
}
