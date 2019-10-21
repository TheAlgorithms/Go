package main

import "fmt"

func fibonacci() func() int { //closure

	i, j := 0, 1
	return func() int {
		i, j = j, i+j
		return i
	}
}

func main() {

	f := fibonacci()

	var number int
	fmt.Println("Which fibonacci number should be calculated? :")
	fmt.Scanf("%d", &number)

	fmt.Println(0)
	for i := 0; i < number; i++ {
		fmt.Println(f())
	}
}
