package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle(array []int) []int {
	rand.Seed(int64(time.Now().UnixNano()))
	length := len(array)

	rand.Shuffle(length, func(i, j int) {
		newI := rand.Intn(length)
		newJ := rand.Intn(length)
		array[i], array[newI] = array[newI], array[i]
		array[j], array[newJ] = array[newJ], array[j]
	})

	return array
}

func isSorted(array []int) bool {
	if len(array) <= 1 {
		return true
	}

	for index := range array[:len(array)-1] {
		if array[index] > array[index+1] {
			return false
		}
	}

	return true
}

func bogoSort(array []int) {
	var iterations uint64 = 0
	startTime := time.Now()

	for !isSorted(array) {
		iterations++
		// fmt.Println(array)
		array = shuffle(array)
	}

	stopTime := time.Now()
	fmt.Printf("\nIterations: %v\nTime Taken: %v\nSorted Array: %v\n\n", iterations, stopTime.Sub(startTime), array)
}

func main() {
	array := []int{2, 0, 1, 9}
	bogoSort(array)
}
