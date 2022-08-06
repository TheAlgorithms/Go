// author(s) [jo3zeph](https://github.com/jo3zeph)
// description: Calculate and output round off values
// see roundoff_test.go

package roundoff

import "math"

//simple code to showcase roundoff function in GoLang
func roundoff(x float64) float64 {
	x = math.Round(x)
	return x
}
