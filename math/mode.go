// mode.go
// author(s): [CalvinNJK] (https://github.com/CalvinNJK)
// description: Finding Mode Value In an Array
// see mode.go

package math

import (
	"errors"

	"github.com/TheAlgorithms/Go/constraints"
)

// ErrEmptySlice is the error returned by functions in math package when
// an empty slice is provided to it as argument when the function expects
// a non-empty slice.
var ErrEmptySlice = errors.New("empty slice provided")

func Mode[T constraints.Number](numbers []T) (T, error) {

	n := len(numbers)

	if n == 0 {
		return 0, ErrEmptySlice
	}

	maxCount := 0
	var maxValue T = 0

	for i := 0; i < n; i++ {

		count := 0

		for k := 0; k < n; k++ {

			if numbers[k] == numbers[i] {
				count++
			}
		}

		if count > maxCount {
			maxCount = count
			maxValue = numbers[i]
		}
	}

	return maxValue, nil

}
