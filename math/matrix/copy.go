package matrix

import "sync"

func (m Matrix[T]) Copy() (Matrix[T], error) {
	rows := m.Rows()
	columns := m.Columns()
	zeroVal, _ := m.Get(0, 0) // Get the zero value of the element type

	copyMatrix := New(rows, columns, zeroVal)
	var wg sync.WaitGroup
	wg.Add(rows)
	errChan := make(chan error, 1)

	for i := 0; i < rows; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < columns; j++ {
				val, _ := m.Get(i, j)
				err := copyMatrix.Set(i, j, val)
				if err != nil {
					select {
					case errChan <- err:
					default:
					}
					return
				}
			}
		}(i)
	}

	wg.Wait()
	close(errChan)

	if err, ok := <-errChan; ok {
		return Matrix[T]{}, err
	}

	return copyMatrix, nil
}
