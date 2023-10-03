package matrix

// Breadth returns the number of columns in a valid matrix.
func Breadth[T any](matrix [][]T) (int, error) {
	if err := IsValid(matrix); err != nil {
		return 0, err
	}

	if len(matrix) == 0 {
		return 0, nil // Return 0 if the matrix is empty and valid.
	}

	return len(matrix[0]), nil
}
