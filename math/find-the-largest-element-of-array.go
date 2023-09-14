package main

import (
	"fmt"
	"log"
)

func main() {
	var arraySize int
	var err error
	fmt.Println("Enter size of array:")
	_, err = fmt.Scan(&arraySize)
	if err != nil {
		log.Fatal(err.Error())
	}

	arrayVal := make([]int, arraySize)
	fmt.Println("Enter value of an array:")
	for i := 0; i < arraySize; i++ {
		_, err = fmt.Scan(&arrayVal[i])
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println("Largest value of an array is:", largest(arrayVal))
}

func largest(arrayVal []int) int {
	n := len(arrayVal)
	maxVal := arrayVal[0]
	for i := 0; i < n; i++ {
		if arrayVal[i] > maxVal {
			maxVal = arrayVal[i]
		}
	}
	return maxVal
}
