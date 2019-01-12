package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	waitGroup sync.WaitGroup
)

func sleepAndReturn(val int, valueChannel chan int) {
	waitGroup.Wait()
	time.Sleep(time.Duration(val) * time.Millisecond)
	valueChannel <- val
}

func main() {
	array := []int{2, 0, 1, 9}
	valueChannel := make(chan int, len(array))
	resultChannel := make(chan []int, 1)

	go func() {
		sortedArray := []int{}
		for _ = range array {
			sortedArray = append(sortedArray, <-valueChannel)
		}
		resultChannel <- sortedArray
	}()

	waitGroup.Add(1)
	for _, val := range array {
		go sleepAndReturn(val, valueChannel)
	}
	waitGroup.Done()

	sortedArray := <-resultChannel

	fmt.Println(sortedArray)
}
