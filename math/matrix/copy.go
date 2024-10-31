// copy.go
// description: Copy a matrix
// details: This function creates a new matrix with the same dimensions as the original matrix and copies all the elements from the original matrix to the new matrix.
// time complexity: O(n*m) where n and m are the dimensions of the matrix
// space complexity: O(n*m) where n and m are the dimensions of the matrix

package matrix

import "sync"

func (m Matrix[T]) Copy() (Matrix[T], error) {

	rows := m.Rows()
	columns := m.Columns()
	if rows == 0 || columns == 0 {
		return Matrix[T]{}, nil
	}
	zeroVal, err := m.Get(0, 0) // Get the zero value of the element type
	if err != nil {
		return Matrix[T]{}, err
	}
	copyMatrix := New(rows, columns, zeroVal)
	var wg sync.WaitGroup
	wg.Add(rows)
	errChan := make(chan error, 1)

	for i := 0; i < rows; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < columns; j++ {
				val, err := m.Get(i, j)
				if err != nil {
					select {
					case errChan <- err:
					default:
					}
					return
				}
				err = copyMatrix.Set(i, j, val)
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
