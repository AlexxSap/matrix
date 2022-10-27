package matrix

import "errors"

const (
	InvalidIndexError = "InvalidIndexError"
	NilMatrixObject   = "NilMatrixObject"
)

type Matrix[T any] struct {
	cells    []T
	rowCount int
	colCount int
}

func NewMatrix[T any](rows, columns int) *Matrix[T] {
	return &Matrix[T]{make([]T, rows*columns), rows, columns}
}

func (m *Matrix[T]) index(row, col int) (int, error) {
	if m == nil {
		return 0, errors.New(NilMatrixObject)
	}
	if row < 0 || col < 0 || row >= m.rowCount || col >= m.colCount {
		return 0, errors.New(InvalidIndexError)
	}
	return m.colCount*row + col, nil
}

func (m *Matrix[T]) pos(index int) (int, int, error) {
	if m == nil {
		return 0, 0, errors.New(NilMatrixObject)
	}
	if index < 0 || index >= len(m.cells) {
		return 0, 0, errors.New(InvalidIndexError)
	}
	return index / m.colCount, index - (index/m.colCount)*m.colCount, nil
}
