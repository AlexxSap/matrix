package matrix

import (
	"errors"
)

const (
	InvalidIndexError = "InvalidIndexError"
	NilMatrixObject   = "NilMatrixObject"
	InvalidMatrixSize = "InvalidMatrixSize"
)

// PairIterator interface for iteraing on any collection with 2 values
type PairIterator interface {
	// Begin set iterator to begin
	Begin()
	// Next iterate to the next element or to the first element (if it is first call) and return true if element exists
	Next() bool
	// First get first value from pair
	First() int
	// Second get second value from pair
	Second() int
}

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

func calcIndex(row, col, maxCol int) int {
	return maxCol*row + col
}

// NewSquareMatrixFromPoints create square matrix of int with filled by
// `value` cells from `points`. Other cells filled by default value for type T.
func NewSquareMatrixFromPoints[T any](points PairIterator, value T) *Matrix[T] {
	max := 0
	calcMax := func(x, y int) {
		if x > max {
			max = x
		}
		if y > max {
			max = y
		}
	}

	points.Begin()
	for points.Next() {
		calcMax(points.First(), points.Second())
	}
	max++

	d := make([]T, max*max)
	points.Begin()
	for points.Next() {
		d[calcIndex(points.First(), points.Second(), max)] = value
	}

	m, _ := NewMatrix(d, max, max)
	return m
}

// Filtered get slice of points {row, column} represents matrix points which satisfy `f`
func (m *Matrix[T]) Filtered(f func(cell T) bool) ([]struct{ Row, Column int }, error) {

	if m == nil {
		return []struct{ Row, Column int }{}, errors.New(NilMatrixObject)
	}

	d := make([]struct{ Row, Column int }, 0)
	var value struct{ Row, Column int }
	for i, val := range m.cells {
		if f(val) {
			value.Row, value.Column, _ = m.pos(i)
			d = append(d, value)
		}
	}

	return d, nil
}

// index convert square coords into slice index
func (m *Matrix[T]) index(row, col int) (int, error) {
	if m == nil {
		return 0, errors.New(NilMatrixObject)
	}
	if row < 0 || col < 0 || row >= m.rowCount || col >= m.colCount {
		return 0, errors.New(InvalidIndexError)
	}

	return calcIndex(row, col, m.colCount), nil
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

// AnyOfPoints check if for any of `points` success functor `f`
func (m *Matrix[T]) AnyOfPoints(points PairIterator, f func(cell T) bool) (bool, error) {
	if m == nil {
		return false, errors.New(NilMatrixObject)
	}

	for points.Next() {
		i, err := m.index(points.First(), points.Second())
		if err != nil {
			return false, err
		}
		if f(m.cells[i]) {
			return true, nil
		}
	}

	return false, nil
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

// SetBatch set `value` to each point [row, column] from slice `points`
func (m *Matrix[T]) SetBatch(value T, points PairIterator) error {
	if m == nil {
		return errors.New(NilMatrixObject)
	}

	for points.Next() {
		i, err := m.index(points.First(), points.Second())
		if err != nil {
			return err
		}
		m.cells[i] = value
	}

	return nil
}

// Transpose transpose matrix
func (m *Matrix[T]) Transpose() error {
	if m == nil {
		return errors.New(NilMatrixObject)
	}

	newCells := make([]T, 0, len(m.cells))

	for col := 0; col < m.colCount; col++ {
		for row := 0; row < m.rowCount; row++ {
			d, _ := m.index(row, col)
			newCells = append(newCells, m.cells[d])
		}
	}
	m.cells = newCells
	m.colCount, m.rowCount = m.rowCount, m.colCount

	return nil
}

// MirrorRows reverse row order
func (m *Matrix[T]) MirrorRows() error {
	if m == nil {
		return errors.New(NilMatrixObject)
	}

	for bRow, eRow := 0, m.rowCount-1; bRow < eRow; bRow, eRow = bRow+1, eRow-1 {
		for c := 0; c < m.colCount; c++ {
			b, _ := m.index(bRow, c)
			e, _ := m.index(eRow, c)
			m.cells[b], m.cells[e] = m.cells[e], m.cells[b]
		}
	}

	return nil
}

// MirrorColumns reverse column order
func (m *Matrix[T]) MirrorColumns() error {
	if m == nil {
		return errors.New(NilMatrixObject)
	}

	for bCol, eCol := 0, m.colCount-1; bCol < eCol; bCol, eCol = bCol+1, eCol-1 {
		for row := 0; row < m.rowCount; row++ {
			b, _ := m.index(row, bCol)
			e, _ := m.index(row, eCol)
			m.cells[b], m.cells[e] = m.cells[e], m.cells[b]
		}
	}

	return nil
}

// Rotate rotate matrix to 90 grad
func (m *Matrix[T]) Rotate() error {
	if m == nil {
		return errors.New(NilMatrixObject)
	}

	err := m.Transpose()
	if err != nil {
		return err
	}

	err = m.MirrorColumns()
	if err != nil {
		return err
	}

	return nil
}
