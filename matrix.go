package matrix

import "errors"

const (
	InvalidIndexError = "InvalidIndexError"
	NilMatrixObject   = "NilMatrixObject"
	InvalidMatrixSize = "InvalidMatrixSize"
)

type Matrix[T any] struct {
	cells    []T
	rowCount int
	colCount int
}

func NewZeroMatrix[T any](rows, columns int) *Matrix[T] {
	return &Matrix[T]{make([]T, rows*columns), rows, columns}
}

func NewMatrix[T any](data []T, rows, columns int) (*Matrix[T], error) {
	if len(data) != rows*columns {
		return nil, errors.New(InvalidMatrixSize)
	}
	return &Matrix[T]{data, rows, columns}, nil
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

func (m *Matrix[T]) RowData(row int) ([]T, error) {
	if m == nil {
		return []T{}, errors.New(NilMatrixObject)
	}
	if row < 0 || row >= m.rowCount {
		return []T{}, errors.New(InvalidIndexError)
	}

	res := make([]T, m.colCount)
	i, _ := m.index(row, 0)
	end, _ := m.index(row, m.colCount-1)

	for ; i <= end; i++ {
		res = append(res, m.cells[i])
	}

	return res, nil
}

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

/// TODO сдвиг строк-столбцов
/// транспонирование
/// сдвиг отдельных ячеек
/// проверка условия для строки - столбца
