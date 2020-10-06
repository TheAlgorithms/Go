// Function to find number of combinations
// with or without repetition
// with or without use of factorial

package main

import (
	"errors"
	"fmt"
)

// factorial of an unsigned integer
// 0 factorial equals 1
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// number of combinations of n things taken k at a time without repetition
func combination(n, k uint) uint {
	if k > n {
		return 0
	}
	return factorial(n) / (factorial(k) * factorial(n-k))
}

// calculate combinations without factorial - Pascal's triangle
func combinationPascal(n, k uint) uint {
	if k == n || k == 0 {
		return 1
	}
	return combinationPascal(n-1, k-1) * n / k
	//return combinationPascal(n-1, k-1) + combinationPascal(n-1, k)
}

// number of combinations of n things taken k at a time WITH repetition
func multicombination(n, k uint) (uint, error) {
	if n > 0 {
		return combinationPascal(n+k-1, k), nil
	}
	return 0, errors.New("Undefined: n is 0") // error handling - n should never be 0
}

func main() {
	fmt.Println("Factorial of 4: ", factorial(4))
	fmt.Println("5 choose 3 - combinations: ", combination(5, 3))
	fmt.Println("5 choose 3 - combinations (no factorial): ", combinationPascal(5, 3))
	for i := 0; i < 5; i++ {
		if r, e := multicombination(uint(i), 3); e != nil {
			fmt.Println(i, "choose 3 -", e)
		} else {
			fmt.Println(i, "choose 3 - combinations with repetition: ", r)
		}
	}
}
