package main

import (
	"fmt"
	"strconv"
)

func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var temp = arr[i]
		var holePosition = i

		for holePosition > 0 && arr[holePosition - 1] > temp {
			arr[holePosition] = arr[holePosition - 1]
			holePosition--
		}
            
		arr[holePosition] = temp
	}
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}
	fmt.Println("")
}

func main() {

	var arr = []int { 15, 3, 12, 6, -9, 9, 0 }

	fmt.Print("Before Sorting: ")
	printArray(arr)

	insertionSort(arr)
	fmt.Print("After Sorting: ")
	printArray(arr)
}
