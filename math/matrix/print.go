package matrix

import (
	"fmt"
)

// Print prints a 2D matrix of any type.
func Print[T any](matrix [][]T) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}
