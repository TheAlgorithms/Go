package main

import "fmt"

func exponentialrecursion(x int, p int) int {
	var temp int
	if p == 0 {
		return 1
	}
	temp = exponentialrecursion(x, p/2)
	y := temp * temp
	if p%2 == 1 {
		y = y * x
	}
	return y
}
func main() {
	fmt.Println(exponentialrecursion(5, 3))
}
