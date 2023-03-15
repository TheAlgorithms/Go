// Implementation of Dutch National Flag Sort algorithm
// Reference: https://en.wikipedia.org/wiki/Dutch_national_flag_problem
// Input: [0, 1, 2, 0, 1, 2]
// Output: [0, 0, 1, 1, 2, 2]

package sort

import "github.com/TheAlgorithms/Go/constraints"

// DutchNationalFlag sorts an array of integers with the Dutch National Flag algorithm
func DutchNationalFlag[T constraints.Signed](arr []T) []T {
	low, mid, high := 0, 0, len(arr)-1
	for mid <= high {
		switch arr[mid] {
		case 0:
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
	return arr
}
