/*
	interpolation search is a search algorithm that works best for sorted
	uniformly distributed arrays.
	best case time complexity is O(log(log(n))
	worst case time complexity is O(n)
*/

package main

import "fmt"

//find target element in sorted uniformly distributed array
func interpolationSearch(array []int, target int) int {
	low := 0
	high := len(array) - 1

	for low <= high && target >= array[low] && target <= array[high] {
		if low == high {
			if array[low] == target {
				return low
			}
			return -1
		}

		// Position estimate considering uniform distribution
		position := low + int((float64(high-low)/float64(array[high]-array[low]))*float64(target-array[low]))

		if array[position] == target {
			return position
		}

		if array[position] < target {
			low = position + 1
		} else {
			high = position - 1
		}
	}
	return -1
}

// recursively find target element in sorted uniformly distributed array
func interpolationSearchRecursive(array []int, low int, high int, target int) int {
	if low > high || target < array[low] || target > array[high] {
		return -1
	}

	// Position estimate considering uniform distribution
	position := low + int((float64(high-low)/float64(array[high]-array[low]))*float64(target-array[low]))

	// If target is larger, target is in right sub array
	if array[position] < target {
		return interpolationSearchRecursive(array, position+1, high, target)
	}

	// If target is smaller, target is in left sub array
	if array[position] > target {
		return interpolationSearchRecursive(array, low, position-1, target)
	}

	// if array[position] == target
	return position
}

func main() {
	array := []int{1, 2, 3, 7, 11, 15, 23, 24, 27, 28, 33, 34, 37, 40, 50, 55, 70}

	targets := []int{23, 12, 70, 1, 0, 90, 33}
	for _, target := range targets {
		index1 := interpolationSearch(array, target)
		index2 := interpolationSearchRecursive(array, 0, len(array)-1, target)

		if index1 != -1 {
			fmt.Println("Iterative: Element ", target, " found at index ", index1)
		} else {
			fmt.Println("Iterative: Element not found.")
		}

		if index2 != -1 {
			fmt.Println("Recursive: Element ", target, " found at index ", index2)
		} else {
			fmt.Println("Recursive: Element not found.")
		}
	}
}
