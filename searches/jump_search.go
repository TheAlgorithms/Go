package main
import (
	"fmt"
	"math"
)
func jumpSearch(arr []int, x int) int {
	var left, right int
	length := len(arr)
	jump := int(math.Sqrt(float64(length)))
	for left < length && arr[left] >= x {
		right = min(length-1, left+jump)
		if arr[left] <= x && arr[right] >= x {
			break
		}
		left += jump
	}
	if left >= length || arr[left] > x {
		return -1
	}
	right = min(length-1, right)
	tempIndex := left
	for tempIndex <= right && arr[tempIndex] <= x {
		if arr[tempIndex] == x {
			return tempIndex
		}
		tempIndex++
	}
	return -1
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func resultPrint(element, index int) {
	if index < 0 {
		fmt.Printf("The element %d was not found\n", element)
	} else {
		fmt.Printf("The element %d is at index %d \n", element, index)
	}
}
func main() {
	values := []int{1, 2, 3, 4, 5, 12, 30, 35, 46, 84}
	resultPrint(5, jumpSearch(values, 5))
	resultPrint(7, jumpSearch(values, 7))
}
  