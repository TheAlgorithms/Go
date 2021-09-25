// centerofcircle.go
// description: Center of the circle
// details:
// Find the center of the circle. - [Center](https://en.wikipedia.org/wiki/Centre_(geometry))
// author(s) [red_byte](https://github.com/i-redbyte)
// see centerofcircle_test.go

// Package geometry contains the solution of geometric problems
package geometry

func FindCenterUsEndpointsDiameter(x1 float32, y1 float32, x2 float32, y2 float32) (float32, float32) {
	return (x1 + x2) / 2, (y1 + y2) / 2
}
