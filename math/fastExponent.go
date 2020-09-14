package main

func fastExponentiation(x int, y int) float64 {
	res := 1 // res is the final result
	for y > 0 {
		// If y is odd, multiply x with result
		if (y & 1) != 0 {
			res = res * x
		}
		// After AND operation with 1 y will be even now
		y = y >> 1 //  Right Bit Shift(y=y/2)
		x = x * x  // Change x to x^2
	}
	return float64(res)
}

func main() {
	print(fastExponentiation(3, 0))
}
