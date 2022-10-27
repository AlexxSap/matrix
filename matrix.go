package matrix

type Matrix[T any] struct {
	cells    []T
	rowCount int
	colCount int
}

func NewMatrix[T any](rows, columns int) *Matrix[T] {
	return &Matrix[T]{make([]T, rows*columns), rows, columns}
}

func (m *Matrix[T]) index(row, col int) int {
	return m.colCount*row + col
}

func (m *Matrix[T]) pos(index int) (int, int) {
	return index / m.colCount, index - (index/m.colCount)*m.colCount
}
