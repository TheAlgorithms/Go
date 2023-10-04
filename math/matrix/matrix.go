package matrix

import (
	"errors"
)

type Matrix[T any] struct {
	elements [][]T
	rows     int
	columns  int
}

// NewMatrix creates a new Matrix based on the provided arguments.
func New[T any](rows, columns int, initial T) *Matrix[T] {
	if rows < 0 || columns < 0 {
		return nil // Invalid dimensions
	}

	matrix := &Matrix[T]{
		elements: make([][]T, rows),
		rows:     rows,    // Set the rows field
		columns:  columns, // Set the columns field
	}

	for i := range matrix.elements {
		matrix.elements[i] = make([]T, columns)
		for j := range matrix.elements[i] {
			matrix.elements[i][j] = initial
		}
	}

	return matrix
}

// NewFromElements creates a new Matrix from the given elements.
func NewFromElements[T any](elements [][]T) (*Matrix[T], error) {
	if !IsValid(elements) {
		return nil, errors.New("rows have different numbers of columns")
	}
	rows := len(elements)
	if rows == 0 {
		return nil, nil // Empty matrix
	}

	columns := len(elements[0])
	matrix := &Matrix[T]{
		elements: make([][]T, rows),
		rows:     rows,    // Set the rows field
		columns:  columns, // Set the columns field
	}
	for i := range matrix.elements {
		matrix.elements[i] = make([]T, columns)
		copy(matrix.elements[i], elements[i])
	}

	return matrix, nil
}

func (m *Matrix[T]) Get(row, col int) (T, error) {
	if row < 0 || row >= m.rows || col < 0 || col >= m.columns {
		var zeroVal T
		return zeroVal, errors.New("index out of range")
	}
	return m.elements[row][col], nil
}

func (m *Matrix[T]) Set(row, col int, val T) error {
	if row < 0 || row >= m.rows || col < 0 || col >= m.columns {
		return errors.New("index out of bounds")
	}

	m.elements[row][col] = val
	return nil
}

func (m *Matrix[T]) Rows() int {
	return len(m.elements)
}

func (m *Matrix[T]) Columns() int {
	if len(m.elements) == 0 {
		return 0
	}
	return len(m.elements[0])
}
