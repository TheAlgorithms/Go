package matrix

import "errors"

// IsValid checks if a matrix is valid, meaning all rows have the same number of columns.
func IsValid[T any](matrix [][]T) error {
	if len(matrix) == 0 {
		return nil // an empty matrix is considered valid
	}

	columns := len(matrix[0]) // Number of columns in the first row.

	for _, row := range matrix {
		if len(row) != columns {
			return errors.New("invalid matrix: rows have different numbers of columns")
		}
	}

	return nil
}
