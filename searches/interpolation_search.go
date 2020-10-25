package main

import (
	"fmt"
)

func interpolationSearch(arr []int, x int) int {
	var first = 0
	var last = len(arr) - 1

	for first <= last && x >= arr[first] && x <= arr[last] {
		if first == last {
			if arr[first] == x {
				return first
			}
			return -1
		}

		tmp := first + int(((last-first)/(arr[last]-arr[first]))*(x-arr[first]))

		if arr[tmp] == x {
			return tmp
		}

		if arr[tmp] < x {
			first = tmp + 1
		} else {
			last = tmp - 1
		}
	}

	return -1
}

func main() {

	arr := []int{10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47}
	var x = 47

	index := interpolationSearch(arr, x)

	if index != -1 {
		fmt.Println("Index:", index)
	} else {
		fmt.Println("Not found")
	}
}
