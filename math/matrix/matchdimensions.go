package matrix

// MatchDimensions checks if two matrices have the same dimensions.
func (m Matrix[T]) MatchDimensions(m1 Matrix[T]) bool {
	if m.rows == m1.rows && m.columns == m1.columns {
		return true
	}
	return false
}
