package matrix

func (m Matrix[T]) Copy() Matrix[T] {
	rows := m.Rows()
	columns := m.Columns()
	zeroVal, _ := m.Get(0, 0) // Get the zero value of the element type

	copyMatrix := New(rows, columns, zeroVal)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			val, _ := m.Get(i, j)
			err := copyMatrix.Set(i, j, val)
			if err != nil {
				panic("copyMatrix.Set error: " + err.Error())
			}
		}
	}

	return copyMatrix
}
