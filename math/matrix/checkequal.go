package matrix

// CheckEqual checks if the current matrix is equal to another matrix (m2).
// Two matrices are considered equal if they have the same dimensions and
// all their elements are equal.
func (m1 Matrix[T]) CheckEqual(m2 Matrix[T]) bool {
	if !m1.MatchDimensions(m2) {
		return false
	}

	for i := range m1.elements {
		for j := range m1.elements[i] {
			if m1.elements[i][j] != m2.elements[i][j] {
				return false
			}
		}
	}
	return true
}
