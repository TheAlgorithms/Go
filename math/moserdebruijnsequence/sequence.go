// The Moser-de Bruijn sequence is the sequence obtained by
// adding up the distinct powers of the number 4 (For example 1, 4, 16, 64, etc).

package main

import "fmt"

func main() {
	number := 5
	sequence := []int{}

	for i := 0; i < number; i++ {
		res := getSequenceNumber(i)
		sequence = append(sequence, res)
	}

	fmt.Println(sequence)
	// [0 1 4 5 16]
}

func getSequenceNumber(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	//number is even
	if num%2 == 0 {
		return 4 * getSequenceNumber(num/2)
	}

	//number is odd
	if num%2 != 0 {
		return 4*getSequenceNumber(num/2) + 1
	}

	return 0
}
