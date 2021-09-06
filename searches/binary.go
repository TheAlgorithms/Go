package search

// Binary Binary Search
func Binary(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex || len(array) == 0 {
		return -1
	}
	mid := (highIndex + lowIndex) / 2
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

// BinaryIterative Iterative Binary Search
func BinaryIterative(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	size := len(array)
	if size == 0 || highIndex > size || lowIndex < 0 {
		return -1
	}
	for startIndex <= endIndex {
		mid = (endIndex + startIndex) / 2
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
