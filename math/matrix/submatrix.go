package matrix

import (
	"context"
	"errors"
	"sync"
)

// SubMatrix extracts a submatrix from the current matrix.
func (m Matrix[T]) SubMatrix(rowStart, colStart, numRows, numCols int) (Matrix[T], error) {
	if rowStart < 0 || colStart < 0 || numRows < 0 || numCols < 0 {
		return Matrix[T]{}, errors.New("negative dimensions are not allowed")
	}

	if rowStart+numRows > m.rows || colStart+numCols > m.columns {
		return Matrix[T]{}, errors.New("submatrix dimensions exceed matrix bounds")
	}

	var zeroVal T
	if numRows == 0 || numCols == 0 {
		return New(numRows, numCols, zeroVal), nil // Return an empty matrix
	}

	subMatrix := New(numRows, numCols, zeroVal)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Make sure it's called to release resources even if no errors

	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for i := 0; i < numRows; i++ {
		i := i // Capture the loop variable for the goroutine
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numCols; j++ {
				select {
				case <-ctx.Done():
					return // Context canceled; return without an error
				default:
				}

				val, err := m.Get(rowStart+i, colStart+j)
				if err != nil {
					cancel()
					select {
					case errCh <- err:
					default:
					}
					return
				}

				err = subMatrix.Set(i, j, val)
				if err != nil {
					cancel()
					select {
					case errCh <- err:
					default:
					}
					return
				}
			}
		}()
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Check for any errors
	if err := <-errCh; err != nil {
		return Matrix[T]{}, err
	}

	return subMatrix, nil
}
