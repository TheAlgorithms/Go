package matrix

import (
	"fmt"
	"reflect"
)

// PrintMatrix prints a 2D matrix of any type.
func PrintMatrix(matrix interface{}) {
	value := reflect.ValueOf(matrix)

	if value.Kind() != reflect.Slice {
		fmt.Println("Input is not a matrix.")
		return
	}

	for i := 0; i < value.Len(); i++ {
		row := value.Index(i)
		if row.Kind() != reflect.Slice {
			fmt.Println("Invalid matrix structure.")
			return
		}

		for j := 0; j < row.Len(); j++ {
			element := row.Index(j)
			fmt.Printf("%v ", element.Interface())
		}
		fmt.Println()
	}
}
