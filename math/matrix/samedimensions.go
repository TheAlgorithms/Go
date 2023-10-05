package matrix

import "fmt"

// SameDimensions checks if two matrices have the same dimensions.
func (m *Matrix[T]) SameDimensions(matrix2 Matrix[T]) (bool, error) {
	rows1 := m.rows
	rows2 := matrix2.rows

	columns1 := m.columns
	columns2 := matrix2.columns

	if rows1 == rows2 && columns1 == columns2 {
		return true, nil
	} else {
		return false, fmt.Errorf("matrices have different dimensions: rows=%d vs columns=%d", rows1, columns1)
	}
}
