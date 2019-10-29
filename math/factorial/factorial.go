package factorial

// package main

func factorial(number uint64) uint64 {
	if number < 2 {
		return 1
	}
	return factorial(number-1) * number
}

// func main() {
// 	fmt.Println(factorial(4))
// }
