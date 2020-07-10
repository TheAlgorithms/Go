/*
This algorithm calculates the distance between two strings.
Parameters: two strings to compare and weights of insertion, substitution and deletion.
Output: distance between both strings
*/

package main

import (
	"fmt"
)

func levenshteinDistance(str1, str2 string, icost, scost, dcost int) int {
	row1 := make([]int, len(str2)+1)
	row2 := make([]int, len(str2)+1)
	var tmp []int

	for i := 1; i <= len(str2); i++ {
		row1[i] = i * icost
	}

	for i := 1; i <= len(str1); i++ {
		row2[0] = i * dcost

		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				row2[j] = row1[j-1]
			} else {
				ins := row2[j-1] + icost
				del := row1[j] + dcost
				sub := row1[j-1] + scost

				if ins < del && ins < sub {
					row2[j] = ins
				} else if del < sub {
					row2[j] = del
				} else {
					row2[j] = sub
				}
			}
		}

		tmp = row1
		row1 = row2
		row2 = tmp
	}

	return row1[len(row1)-1]
}

/*
func min(str1, str2 int) int {
	if str1 < str2 {
		return str1
	}
	return str2
}
*/

func main() {
	str1 := "stingy"
	str2 := "ring"

	// Using weight 1 for insertion, substitution and deletion
	strDistance1 := levenshteinDistance(str1, str2, 1, 1, 1)
	fmt.Printf("Distance between \"%s\" and \"%s\" is: %d.\n", str1, str2, strDistance1)
	// Output: Distance between "stingy" and "ring" is: 3.

	// Using weight 1 for insertion/substitution and weight 3 for deletion
	strDistance2 := levenshteinDistance(str1, str2, 1, 1, 3)
	fmt.Printf("Distance between \"%s\" and \"%s\" is: %d.\n", str1, str2, strDistance2)
	// Output: Distance between "stingy" and "ring" is: 7.

}
