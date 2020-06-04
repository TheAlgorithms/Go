/*
	Ternary search implementation in Go

	    Worst-case performance	Θ(log3(N))
 		Best-case performance	Θ(1)
 		Average performance	    Θ(log3(N))
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Ternary Search of 1 in primes should return: " + strconv.Itoa(-1) + ", and it returns: " + strconv.Itoa(ternarySearch(primes, 1)))
	fmt.Println("Ternary Search of 3 in primes should return: " + strconv.Itoa(1) + ", and it returns: " + strconv.Itoa(ternarySearch(primes, 3)))
	fmt.Println("Ternary Search of 13 in primes should return: " + strconv.Itoa(5) + ", and it returns: " + strconv.Itoa(ternarySearch(primes, 13)))
}

func ternarySearch(array []int, target int) int {
	return ternarySearchAux(array, target, 0, len(array)-1)
}

func ternarySearchAux(array []int, target int, lowIndex int, highIndex int) int {

	// Value not found
	if highIndex < lowIndex {
		return -1
	}

	// First boundary: between lowIndex and first 1/3 of array
	mid1 := int(lowIndex + (highIndex-lowIndex)/3)

	// Second boundary: between first 1/3 and second 1/3 of array
	mid2 := int(lowIndex + 2*(highIndex-lowIndex)/3)

	if target == array[mid1] {
		// Value found in first boundary position
		return mid1
	} else if target == array[mid1] {
		// Value found in second boundary position
		return mid2
	} else if target < array[mid1] {
		// Value to be found between lowIndex and first boundary
		return ternarySearchAux(array, target, lowIndex, mid1-1)
	} else if target > array[mid2] {
		// Value to be found between second boundary and highIndex
		return ternarySearchAux(array, target, mid2+1, highIndex)
	} else {
		// Value to be found between boundaries
		return ternarySearchAux(array, target, mid1, mid2)
	}
}
