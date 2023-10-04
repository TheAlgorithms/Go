package matrix

// IsValid checks if the input matrix has consistent row lengths.
func IsValid[T any](elements [][]T) bool {
	if len(elements) == 0 {
		return true
	}
	columns := len(elements[0])
	for _, row := range elements {
		if len(row) != columns {
			return false
		}
	}
	return true
}
