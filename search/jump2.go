package search

import "math"

func Jump2(arr []int, target int) (int, error) {
	step := int(math.Round(math.Sqrt(float64(len(arr)))))
	rbound := len(arr)
	for i := step; i < len(arr); i += step {
		if arr[i] > target {
			rbound = i
			break
		}
	}

	for i := rbound - step; i < rbound; i++ {
		if arr[i] == target {
			return i, nil
		}
		if arr[i] > target {
			break
		}
	}
	return -1, ErrNotFound
}
