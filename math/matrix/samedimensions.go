package matrix

import "fmt"

// SameDimensions checks if two matrices have the same dimensions.
func SameDimensions[T any](matrix1, matrix2 [][]T) (int, error) {
	len1, err1 := Len(matrix1)
	len2, err2 := Len(matrix2)
	if err1 != nil {
		return 0, fmt.Errorf("%v", err1)
	} else if err2 != nil {
		return 0, fmt.Errorf("%v", err2)
	}

	breadth1, err3 := Breadth(matrix1)
	breadth2, err4 := Breadth(matrix2)
	if err3 != nil {
		return 0, fmt.Errorf("%v", err3)
	} else if err4 != nil {
		return 0, fmt.Errorf("%v", err4)
	}

	if (len1 == len2) && (breadth1 == breadth2) {
		return 1, nil
	} else {
		return 0, fmt.Errorf("matrices have different dimensions: %v", err3)
	}

}

// func main() {
// 	matrix1 := [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 	}
// 	matrix2 := [][]int{
// 		{7, 8, 9},
// 	}
// 	SameDimensions(matrix1, matrix2)
// }
