package matrix

import "fmt"

// Print prints a 2D matrix of any type.
func (m *Matrix[T]) Print() {
	for i := range m.elements {
		for j := range m.elements[i] {
			fmt.Printf("%v ", m.elements[i][j])
		}
		fmt.Println()
	}
}
