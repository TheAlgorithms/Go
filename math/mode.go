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

	countMap := make(map[T]int)

	n := len(numbers)

	if n == 0 {
		return 0, ErrEmptySlice
	}

	for i := 0; i < n; i++ {
		temp := numbers[i]
		value, check := countMap[temp]
		if check == false {
			countMap[temp] = 1
		} else {
			countMap[temp] = value + 1
		}
	}

	var mod T
	count := 0

	for k, v := range countMap {
		if v > count {
			count = v
			mod = k
		}
	}

	return mod, nil

}
