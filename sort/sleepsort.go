// sleepsort.go
// Implementation of Sleep Sort algorithm (mostly as a fun experiment)
// Note: This is not a practical sorting algorithm and relies on goroutines and time.Sleep
// Reference: https://en.wikipedia.org/wiki/Sorting_algorithm#Sleep_sort

package sort

import (
	"sync"
	"time"
)

// SleepSort performs the sleep sort algorithm on the given array.
// It uses goroutines to "sort" by sleeping for a duration proportional to each element's value.
// WARNING: Not suitable for production or large numbers, as it is inefficient and can cause timing issues.
func SleepSort(arr []int) []int {
	if len(arr) == 0 { // handle empty array
		return arr
	}

	var wg sync.WaitGroup
	resultChan := make(chan int, len(arr))

	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond) // sleep for n milliseconds
			resultChan <- n
		}(num)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect sorted results
	var result []int
	for num := range resultChan {
		result = append(result, num)
	}

	return result
}

// Example usage:
// func main() {
//     arr := []int{3, 1, 4, 2}
//     sortedArr := SleepSort(arr)
//     fmt.Println(sortedArr) // Output may vary based on timing accuracy
// }
