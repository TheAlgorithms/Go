package search

// Interpolation searches for the entity in the given sortedData.
// if the entity is present, it will return the index of the entity, if not -1 will be returned.
// see: https://en.wikipedia.org/wiki/Interpolation_search
// Complexity
//
//	Worst: O(N)
//	Average: O(log(log(N))  if the elements are uniformly distributed
//	Best: O(1)
//
// Example
//
//	fmt.Println(InterpolationSearch([]int{1, 2, 9, 20, 31, 45, 63, 70, 100},100))
func Interpolation(sortedData []int, guess int) (int, error) {
	if len(sortedData) == 0 {
		return -1, ErrNotFound
	}

	var (
		low, high       = 0, len(sortedData) - 1
		lowVal, highVal = sortedData[low], sortedData[high]
	)

	for lowVal != highVal && (lowVal <= guess) && (guess <= highVal) {
		mid := low + int(float64(float64((guess-lowVal)*(high-low))/float64(highVal-lowVal)))

		// if guess is found, array can also have duplicate values, so scan backwards and find the first index
		if sortedData[mid] == guess {
			for mid > 0 && sortedData[mid-1] == guess {
				mid--
			}
			return mid, nil

		}

		// adjust our guess and continue
		if sortedData[mid] > guess {
			high, highVal = mid-1, sortedData[high]

		} else {
			low, lowVal = mid+1, sortedData[low]
		}

	}

	if guess == lowVal {
		return low, nil
	}
	return -1, ErrNotFound
}
