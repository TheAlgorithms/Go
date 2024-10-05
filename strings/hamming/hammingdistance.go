/*
This algorithm calculates the hamming distance between two equal length strings.
The Hamming distance between two equal-length strings of symbols is the number of positions
at which the corresponding symbols are different:
https://en.wikipedia.org/wiki/Hamming_distance

Note that we didn't consider strings as an array of bytes, therefore, we didn't use the XOR operator.
In this case, we used a simple loop to compare each character of the strings, and if they are different,
we increment the hamming distance by 1.

Parameters: two strings to compare
Output: distance between both strings */

package hamming

import "errors"

func Distance(str1, str2 string) (int, error) {
	if len(str1) != len(str2) {
		return -1, errors.New("strings must have a same length")
	}

	hammingDistance := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
