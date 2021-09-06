/*
This algorithm calculates the distance between two strings.
Parameters: two strings to compare and weights of insertion, substitution and deletion.
Output: distance between both strings
*/

package levenshtein

// Distance Function that gives Levenshtein Distance
func Distance(str1, str2 string, icost, scost, dcost int) int {
	row1 := make([]int, len(str2)+1)
	row2 := make([]int, len(str2)+1)

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
		row1, row2 = row2, row1
	}

	return row1[len(row1)-1]
}
