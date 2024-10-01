// Patience sorting is a sorting algorithm inspired by the card game patience.
//
// For more details check out those links below here:
// GeeksForGeeks article : https://www.geeksforgeeks.org/patience-sorting/
// Wikipedia article: https://en.wikipedia.org/wiki/Patience_sorting
// authors [guuzaa](https://github.com/guuzaa)
// worst-case time complexity: O(n log n)
// average time complexity: O(n log n)
// space complexity: O(n)
// see patiencesort.go

package sort

import "github.com/TheAlgorithms/Go/constraints"

func Patience[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	var piles [][]T

	for _, card := range arr {
		left, right := 0, len(piles)
		for left < right {
			mid := left + (right-left)/2
			if piles[mid][len(piles[mid])-1] >= card {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == len(piles) {
			piles = append(piles, []T{card})
		} else {
			piles[left] = append(piles[left], card)
		}
	}

	return mergePiles(piles)
}

func mergePiles[T constraints.Ordered](piles [][]T) []T {
	var ret []T

	for len(piles) > 0 {
		minID := 0
		minValue := piles[minID][len(piles[minID])-1]

		for i := 1; i < len(piles); i++ {
			if minValue <= piles[i][len(piles[i])-1] {
				continue
			}

			minValue = piles[i][len(piles[i])-1]
			minID = i
		}

		ret = append(ret, minValue)

		piles[minID] = piles[minID][:len(piles[minID])-1]

		if len(piles[minID]) == 0 {
			piles = append(piles[:minID], piles[minID+1:]...)
		}
	}

	return ret
}
