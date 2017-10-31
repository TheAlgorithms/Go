package main

import "fmt"
import "math/rand"

func main() {
	arr := RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")

	PancakeSort(arr)

	fmt.Println("Sorted array is: ", arr)
}

func RandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func PancakeSort(arr []int) {
	for uns := len(arr) - 1; uns > 0; uns-- {
		lx, lg := 0, arr[0]
		for i := 1; i <= uns; i++ {
			if arr[i] > lg {
				lx, lg = i, arr[i]
			}
		}
		pancakeFlip(arr, lx)
		pancakeFlip(arr, uns)
	}
}

func pancakeFlip(arr []int, r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
}
