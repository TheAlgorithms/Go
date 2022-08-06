// mode.go
// author(s): [CalvinNJK] (https://github.com/CalvinNJK)
// description: Finding Mode Value In an Array
// see mode.go

package math

//package constraints

import (
	"errors"

	"github.com/TheAlgorithms/Go/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Error when array is empty
var ErrEmptySlice = errors.New("empty slice provided")

func Mode[T constraints.Number](numbers []T) (T, error) {

	// Indicate numbers of checking(looping)
	n := len(numbers)

	// Assign when there is a Mode Value
	maxCount := 0
	var maxValue T = 0

	if n == 0 {
		return 0, emptyModeArr
	} else {

		// Check for each value has any repeated value
		for i := 0; i < n; i++ {

			count := 0

			// Compare the all the values in the array with the selected value
			for k := 0; k < n; k++ {

				// If found repeated value, add into count
				if numbers[k] == numbers[i] {
					count++
				}
			}

			// Check is the selected value's count in the array,
			// is it more than previous value(MaxValue)'s count
			if count > maxCount {
				maxCount = count
				maxValue = numbers[i]
			}
		}

		return maxValue, nil

	}
}
