// heaps.go
// description: Implementation of Heap's Algorithm for generating all permutations of n objects
// time complexity: O(n!)
// space complexity: O(n)

package permutation

import (
	"strings"
)

// Heap's Algorithm for generating all permutations of n objects
func Heaps(out chan []string, n int) {
	elementSetCh := make(chan []string)
	go GenerateElementSet(elementSetCh, n)
	elementSet := <-elementSetCh

	var recursiveGenerate func([]string, int, []string)
	var permutations []string
	recursiveGenerate = func(previousIteration []string, n int, elements []string) {
		if n == 1 {
			permutations = append(permutations, strings.Join(elements, ""))
		} else {
			for i := 0; i < n; i++ {
				recursiveGenerate(previousIteration, n-1, elements)
				if n%2 == 1 {
					tmp := elements[i]
					elements[i] = elements[n-1]
					elements[n-1] = tmp
				} else {
					tmp := elements[0]
					elements[0] = elements[n-1]
					elements[n-1] = tmp
				}
			}
		}
	}
	recursiveGenerate(permutations, n, elementSet)
	out <- permutations
}

func GenerateElementSet(out chan []string, n int) {
	elementSet := make([]string, n)
	for i := range elementSet {
		elementSet[i] = string(rune(i + 49)) // Adjust this if you want to change your charset
	}
	out <- elementSet
}
