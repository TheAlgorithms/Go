package main

import (
	"fmt"
)

func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func bubbleSort(arrayzor []int) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor) - 1; i++ {
			if arrayzor[i + 1] < arrayzor[i] {
				Swap(arrayzor, i, i + 1)
				swapped = true
			}
		}
	}	
}

func Swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func main() {

	arrayzor := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arrayzor)
	bubbleSort(arrayzor)
	fmt.Println("Sorted array: ", arrayzor)
}