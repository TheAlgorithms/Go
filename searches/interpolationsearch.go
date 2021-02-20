package main

import "fmt"

func InterpolationSearch() {
	searchValue := 15

	arr := [10]int{1, 5, 100, 0, -100, 15, 4, 102, 30, 1000}
	fmt.Println(arr)

	// Sort the numbers
	tmp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	fmt.Println(arr)

	left := 0
	right := len(arr) - 1

	if right < left {
		fmt.Println("Not found")
		return
	}
	// We will use the following interpolation formula to find the position
	// pos = left + ((right - left) / (arr[right] - arr[left])) * (searchValue - arr[left])
	for left <= right {
		pos := left + ((right-left)/(arr[right]-arr[left]))*(searchValue-arr[left])

		if arr[pos] == searchValue {
			fmt.Println("Found at position: ", pos)
			return
		} else if arr[pos] < searchValue {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	fmt.Println("Not found")
}
