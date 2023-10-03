package matrix

// Len returns the number of rows in a valid matrix.
func Len[T any](matrix [][]T) (int, error) {
	if err := IsValid(matrix); err != nil {
		return 0, err
	}
	return len(matrix), nil
}