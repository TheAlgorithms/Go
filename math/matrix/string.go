package matrix

import "fmt"

// String implements the fmt.Stringer interface for Matrix.
func (m Matrix[T]) String() string {
	var result string
	for i := range m.elements {
		for j := range m.elements[i] {
			result += fmt.Sprintf("%v ", m.elements[i][j])
		}
		result += "\n"
	}
	return result
}
