package binarysearch

// BinarySearch Binary Search
func BinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex || len(array) == 0 {
		return -1
	}
	mid := (highIndex + lowIndex) / 2
	if array[mid] > target {
		return BinarySearch(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return BinarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

// IterBinarySearch Iter Binary Search
func IterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
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
