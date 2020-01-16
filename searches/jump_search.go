package main

import (
	"fmt"
	"math"
)

//jumpSearch an algorithm to search for sorted array
func jumpSearch(array []int, query int) int {
	var n = len(array)
	var jumpStep = int(math.Sqrt(float64(n))) //block size to jump

	var prev int = 0

	//get the block in which the element may be preset
	for array[int(math.Min(float64(jumpStep), float64(n)))-1] < query {
		prev = jumpStep
		jumpStep += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	//now search for the elemnet in the block found above
	for array[prev] < query {
		prev++

		//if we go to the end or the next block it means the elemnet is not preset
		if prev == int(math.Min(float64(jumpStep), float64(n))) {
			return -1
		}
	}

	if array[prev] == query {
		return prev
	}

	return -1
}

func main() {

	fmt.Println("Jump search:")
	array := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	index := jumpSearch(array, 12)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("array[", index, "] = ", array[index])
	}
}
