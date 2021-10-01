package straightlines

import (
	"math"
)

// Defines a point with x and y coordinates.
type Point struct {
	x, y float64
}

type Line struct {
	p1, p2 Point
}

// Calculates the shortest distance between two points.
func distance(a, b *Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

// Calcualtes the Point that divides a line in specific ratio.
// DO NOT specify the ratio in the form m:n, specify it as r, where r = m / n.
func section(p1, p2 *Point, r float64) Point {
	var point Point
	point.x = (r*p2.x + p1.x) / (r + 1)
	point.y = (r*p2.y + p1.y) / (r + 1)
	return point
}

// Calculates the slope (gradient) of a line.
func slope(l *Line) float64 {
	return (l.p2.y - l.p1.y) / (l.p2.x - l.p1.x)
}

// Calculates the Y-Intercept of a line from a specific Point.
func intercept(p *Point, slope float64) float64 {
	return p.y - (slope * p.x)
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
	return math.Abs(equation[0]*p.x+equation[1]*p.y+equation[2]) / math.Sqrt(math.Pow(equation[0], 2)+math.Pow(equation[1], 2))
}
