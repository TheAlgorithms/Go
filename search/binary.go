package search

// Binary search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// This function uses recursive call to itself.
// If a target is found, the index of the target is returned. Else the function return -1 and ErrNotFound.
func Binary(array []int, target int, lowIndex int, highIndex int) (int, error) {
	if highIndex < lowIndex || len(array) == 0 {
		return -1, ErrNotFound
	}
	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}

// BinaryIterative search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// Unlike Binary, this function uses iterative method and not recursive.
// If a target is found, the index of the target is returned. Else the function return -1 and ErrNotFound.
func BinaryIterative(array []int, target int, lowIndex int, highIndex int) (int, error) {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}
