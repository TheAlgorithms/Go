package math

func Lerp(v0, v1, t float64) float64 {
	// see: https://en.wikipedia.org/wiki/Linear_interpolation
	return (1-t)*v0 + t*v1
}
