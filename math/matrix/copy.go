package matrix

func (m *Matrix[T]) Copy() *Matrix[T] {
	rows := m.rows
	columns := m.columns
	zeroVal, _ := m.Get(0, 0) // Get the zero value of the element type

	copyMatrix := New(rows, columns, zeroVal)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			val, _ := m.Get(i, j)
			copyMatrix.Set(i, j, val)
		}
	}

	return copyMatrix
}
