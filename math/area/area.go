package main

import (
	"fmt"
	"math"
)

//SurfaceAreaCube returns the Surface Area of a Cube with 6 sides
func SurfaceAreaCube(sideLength float64) float64 {
	if sideLength < 0 {
		fmt.Println("SurfaceCube() only accepts non-negative values")
		return 0
	}
	return 6 * math.Pow(sideLength, 2)
}

// SurfaceAreaSphereCalculate the Surface Area of a Sphere. Wikipedia reference: https://en.wikipedia.org/wiki/Sphere
func SurfaceAreaSphere(radius float64) float64 {
	if radius < 0 {
		fmt.Println("SurfaceAreaSphere() only accepts non-negative values")
		return 0
	}
	return math.Pow(4*math.Pi*radius, 2)
}
