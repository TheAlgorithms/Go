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
func BinaryIterative(array []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(array) - 1
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

// LowerBound returns index to the first element in the range [0, len(array)-1] that is not less than (i.e. greater or equal to) target.
// return -1 and ErrNotFound if no such element is found.
func LowerBound(array []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(array) - 1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] < target {
			startIndex = mid + 1
		} else {
			endIndex = mid - 1
		}
	}

	//when target greater than every element in array, startIndex will out of bounds
	if startIndex >= len(array) {
		return -1, ErrNotFound
	}
	return startIndex, nil
}

// UpperBound returns index to the first element in the range [lowIndex, len(array)-1] that is greater than target.
// return -1 and ErrNotFound if no such element is found.
func UpperBound(array []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(array) - 1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else {
			startIndex = mid + 1
		}
	}

	//when target greater or equal than every element in array, startIndex will out of bounds
	if startIndex >= len(array) {
		return -1, ErrNotFound
	}
	return startIndex, nil
}
