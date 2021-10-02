package straightlines

import (
	"math"
)

// Defines a point with x and y coordinates.
type Point struct {
	X, Y float64
}

type Line struct {
	P1, P2 Point
}

// Calculates the shortest distance between two points.
func distance(a, b *Point) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

// Calcualtes the Point that divides a line in specific ratio.
// DO NOT specify the ratio in the form m:n, specify it as r, where r = m / n.
func section(p1, p2 *Point, r float64) Point {
	var point Point
	point.X = (r*p2.X + p1.X) / (r + 1)
	point.Y = (r*p2.Y + p1.Y) / (r + 1)
	return point
}

// Calculates the slope (gradient) of a line.
func slope(l *Line) float64 {
	return (l.P2.Y - l.P1.Y) / (l.P2.X - l.P1.X)
}

// Calculates the Y-Intercept of a line from a specific Point.
func intercept(p *Point, slope float64) float64 {
	return p.Y - (slope * p.X)
}

// Checks if two lines are parallel or not.
func isParallel(l1, l2 *Line) bool {
	return slope(l1) == slope(l2)
}

// Checks if two lines are perpendicular or not.
func isPerpendicular(l1, l2 *Line) bool {
	return slope(l1)*slope(l2) == -1
}

// Calculates the distance of a given Point from a given line.
// The slice should contain the coefficiet of x, the coefficient of y and the constant in the respective order.
func pointDistance(p *Point, equation [3]float64) float64 {
	return math.Abs(equation[0]*p.X+equation[1]*p.Y+equation[2]) / math.Sqrt(math.Pow(equation[0], 2)+math.Pow(equation[1], 2))
}
