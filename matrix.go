package matrix

type Matrix[T any] struct {
	cells []T
}

func NewMatrix[T any](rows, columns int) *Matrix[T] {
	return &Matrix[T]{make([]T, rows*columns)}
}

func (m *Matrix[T]) index(row, col int) int {

}

func (m *Matrix[T]) pos(index int) (int, int) {

}
