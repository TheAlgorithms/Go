package main

import (
	"fmt"
)

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func cocktailSort(arrayzor []int) {
	var i int
	for !isSorted(arrayzor) {
		for i = 0; i < len(arrayzor)-2; i++ {
			if arrayzor[i] > arrayzor[i+1] {
				arrayzor[i], arrayzor[i+1] = arrayzor[i+1], arrayzor[i]
			}
		}
		for ; i > 0; i-- {
			if arrayzor[i] > arrayzor[i+1] {
				arrayzor[i], arrayzor[i+1] = arrayzor[i+1], arrayzor[i]
			}
		}
	}
}

func main() {
	arrayzor := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arrayzor)
	cocktailSort(arrayzor)
	fmt.Println("Sorted array: ", arrayzor)
}
