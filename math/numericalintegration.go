package math

//TrapezoidalInteg function implements the Trapezoidal rule of numerical integration to find the approximate area under
//the curve of a function f between the limits a and b. The last parameter h denotes the step size. Smaller the value of h, more is the accuracy of the result, but more will be the time taken to evaluate the integral.
//https://en.wikipedia.org/wiki/Trapezoidal_rule
//Sample usage: area:=TrapezoidalInteg(func(x float64) float64 { return x * x }, 0, 1, 200)

func TrapezoidalInteg(f func(float64) float64, a, b float64, h float64) float64 {
	sum := 0.0
	for i := a; i < b; i += h {
		sum += h * (f(i) + f(i+h)) / 2
	}
	return sum
}

//SimpsonsOneThirdInteg function implements Simpson's 1/3rd rule of integration to find the approximate area under the
//curve of a function f between teh limits a and b. The last aprameter n is the number of intervals to use.
//Higher the value of n, more is the accuracy of the result, but more will be the time taken to evaluate the area.
//n has to be even. If odd number is provided, it will be incremented.
//https://en.wikipedia.org/wiki/Simpson%27s_rule
//Sample usage: area:=SimpsonsOneThirdInteg(func(x float64) float64 { return x * x }, 0, 1, 200)
func SimpsonsOneThirdInteg(f func(float64) float64, a, b float64, n int) float64 {
	//n has to be even
	if n%2 != 0 {
		n++
	}
	h := (b - a) / float64(n)
	sum := 0.0
	for i := a; i < b; i += h {
		sum += h * (f(i) + 4*f(i+h/2) + f(i+h)) / 6
	}
	return sum
}
