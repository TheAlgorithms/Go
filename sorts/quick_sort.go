package main

import "fmt"
import "math/rand"

func quick_sort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	// Pick a pivot
	pivot := rand.Int() % len(arr)

	left, right := 0, len(arr)-1

	// Move the pivot to the right
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Pile elements smaller than the pivot on the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	//Do this recursively
	quick_sort(arr[:left])
	quick_sort(arr[left+1:])

	return arr
}

func main() {
	arr := RandomArray(10)

	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("Sorted array is: ", quick_sort(arr))
}

func RandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
