// Package sort
// Patience sorting is a sorting algorithm inspired by the card game patience.
//
// For more details check out those links below here:
// GeeksForGeeks article : https://www.geeksforgeeks.org/patience-sorting/
// Wikipedia article: https://en.wikipedia.org/wiki/Patience_sorting
// authors [guuzaa](https://github.com/guuzaa)
// see patiencesort.go
package sort

import "github.com/TheAlgorithms/Go/constraints"

func Patience[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	var piles [][]T

	// Traverse the given array
	for _, card := range arr {
		// Iterate over all the piles
		// and check the top of the stack of each pile
		// is less than the current element or not.
		left, right := 0, len(piles)
		for left < right {
			mid := left + (right-left)>>1
			if piles[mid][len(piles[mid])-1] >= card {
				right = mid
			} else {
				left = mid + 1
			}
		}

		// If all the top of the stack of each pile is less than than the current element
		if left == len(piles) {
			// Create a new pile with the current element
			// and insert the pile into all the piles
			piles = append(piles, []T{card})
		} else {
			// The current element is placed on the leftmost existing pile 
			// whose top card has a value greater than or equal to the new card's value
			piles[left] = append(piles[left], card)
		}
	}

	return mergePiles(piles)
}

func mergePiles[T constraints.Ordered](piles [][]T) []T {
	var ret []T

	// Find the smallest element of top of pile
	// and remove it from the piles
	// and store it into the final array , in every iteration,
	// util all piles are empty
	for len(piles) > 0 {

		// Store the index and the value of the smallest element
		// of the top of the piles
		minId := 0
		minV := piles[minId][len(piles[minId])-1]

		// Calculate the smallest element of the top of the every pile
		for i := 1; i < len(piles); i++ {

			// if minV > the top of the current pile
			if minV > piles[i][len(piles[i])-1] {
				// Update minV
				minV = piles[i][len(piles[i])-1]
				// Update index
				minId = i
			}
		}

		// Insert the smallest element of the top of the pile
		ret = append(ret, minV)

		// Remove the top element from the current pile
		piles[minId] = piles[minId][:len(piles[minId])-1]

		// If current pile is empty
		if len(piles[minId]) == 0 {
			// Remove current pile form all piles
			piles = append(piles[:minId], piles[minId+1:]...)
		}
	}

	return ret
}
