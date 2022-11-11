package matrix

import (
	"errors"
)

const (
	InvalidIndexError = "InvalidIndexError"
	NilMatrixObject   = "NilMatrixObject"
	InvalidMatrixSize = "InvalidMatrixSize"
)

// Matrix represent simple square of any type data
type Matrix[T any] struct {
	cells    []T
	rowCount int
	colCount int
}

// NewZeroMatrix create default initialized matrix with specified size
func NewZeroMatrix[T any](rows, columns int) *Matrix[T] {
	return &Matrix[T]{make([]T, rows*columns), rows, columns}
}

// New Matrix create matrix from slice of data with spicified size
func NewMatrix[T any](data []T, rows, columns int) (*Matrix[T], error) {
	if len(data) != rows*columns {
		return nil, errors.New(InvalidMatrixSize)
	}
	return &Matrix[T]{data, rows, columns}, nil
}

// index convert square coords into slice index
func (m *Matrix[T]) index(row, col int) (int, error) {
	if m == nil {
		return 0, errors.New(NilMatrixObject)
	}
	if row < 0 || col < 0 || row >= m.rowCount || col >= m.colCount {
		return 0, errors.New(InvalidIndexError)
	}
	return m.colCount*row + col, nil
}

// pos convert slice index int square coords
func (m *Matrix[T]) pos(index int) (int, int, error) {
	if m == nil {
		return 0, 0, errors.New(NilMatrixObject)
	}
	if index < 0 || index >= len(m.cells) {
		return 0, 0, errors.New(InvalidIndexError)
	}
	return index / m.colCount, index - (index/m.colCount)*m.colCount, nil
}

// RowData get slice of values stored in spicified row
func (m *Matrix[T]) RowData(row int) ([]T, error) {
	if m == nil {
		return []T{}, errors.New(NilMatrixObject)
	}
	if row < 0 || row >= m.rowCount {
		return []T{}, errors.New(InvalidIndexError)
	}

	res := make([]T, 0, m.colCount)
	i, _ := m.index(row, 0)
	end, _ := m.index(row, m.colCount-1)

	for ; i <= end; i++ {
		res = append(res, m.cells[i])
	}

	return res, nil
}

// ColumnData get slice of values stored in spicified column
func (m *Matrix[T]) ColumnData(col int) ([]T, error) {
	if m == nil {
		return []T{}, errors.New(NilMatrixObject)
	}
	if col < 0 || col >= m.colCount {
		return []T{}, errors.New(InvalidIndexError)
	}

	res := make([]T, 0, m.rowCount)
	for i := 0; i < m.rowCount; i++ {
		index, _ := m.index(i, col)
		res = append(res, m.cells[index])
	}

	return res, nil
}

// AllOfRow check `f` for each value on `row`
func (m *Matrix[T]) AllOfRow(row int, f func(cell T) bool) (bool, error) {
	if m == nil {
		return false, errors.New(NilMatrixObject)
	}
	if row < 0 || row >= m.rowCount {
		return false, errors.New(InvalidIndexError)
	}

	r, _ := m.RowData(row)

	for _, cell := range r {
		if !f(cell) {
			return false, nil
		}
	}
	return true, nil
}

// AllOfColumn check `f` for each value on `col`
func (m *Matrix[T]) AllOfColumn(col int, f func(cell T) bool) (bool, error) {
	if m == nil {
		return false, errors.New(NilMatrixObject)
	}
	if col < 0 || col >= m.colCount {
		return false, errors.New(InvalidIndexError)
	}

	r, _ := m.ColumnData(col)

	for _, cell := range r {
		if !f(cell) {
			return false, nil
		}
	}
	return true, nil
}

// ShiftRowsDown shift all rows down to 1 row. First row make default values row.
func (m *Matrix[T]) ShiftRowsDown() error {
	if m == nil {
		return errors.New(NilMatrixObject)
	}

	for i := len(m.cells) - 1; i >= m.colCount; i-- {
		m.cells[i] = m.cells[i-m.colCount]
	}

	var def T
	for i := 0; i < m.colCount; i++ {
		m.cells[i] = def
	}

	return nil
}

// Set value `value` to cell [row, column]
func (m *Matrix[T]) Set(row, column int, value T) error {
	if m == nil {
		return errors.New(NilMatrixObject)
	}

	i, err := m.index(row, column)
	if err != nil {
		return err
	}

	m.cells[i] = value

	return nil
}

// Get `value` from matrix on [row,column]
func (m *Matrix[T]) Get(row, column int) (T, error) {
	var empty T
	if m == nil {
		return empty, errors.New(NilMatrixObject)
	}

	i, err := m.index(row, column)
	if err != nil {
		return empty, err
	}

	return m.cells[i], nil
}
