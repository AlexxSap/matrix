package matrix

import (
	"strconv"
	"testing"
)

func Test_Index(t *testing.T) {
	m := NewMatrix[int](3, 3)
	/*
		column \ row    0 1 2
						_ _ _
					0 | 0 1 2
					1 | 3 4 5
					2 | 6 7 8
	*/

	test := func(row, column, expected int) {
		t.Run(strconv.Itoa(expected), func(t *testing.T) {
			actual, err := m.index(row, column)
			if err != nil {
				t.Fatal(err)
			}
			if actual != expected {
				t.Errorf("act: %d, exp: %d", actual, expected)
			}
		})
	}

	test(0, 0, 0)
	test(0, 1, 1)
	test(0, 2, 2)
	test(1, 0, 3)
	test(1, 1, 4)
	test(1, 2, 5)
	test(2, 0, 6)
	test(2, 1, 7)
	test(2, 2, 8)
}

func Test_Pos(t *testing.T) {
	m := NewMatrix[int](3, 3)
	/*
		column \ row    0 1 2
						_ _ _
					0 | 0 1 2
					1 | 3 4 5
					2 | 6 7 8
	*/

	test := func(index, expRow, expCol int) {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			actRow, actCol, err := m.pos(index)
			if err != nil {
				t.Fatal(err)
			}
			if actRow != expRow || actCol != expCol {
				t.Errorf("act: %d, %d, exp: %d,%d", actRow, actCol, expRow, expCol)
			}
		})
	}

	test(0, 0, 0)
	test(1, 0, 1)
	test(2, 0, 2)
	test(3, 1, 0)
	test(4, 1, 1)
	test(5, 1, 2)
	test(6, 2, 0)
	test(7, 2, 1)
	test(8, 2, 2)
}

func Test_IndexErrors(t *testing.T) {
	var m *Matrix[int]
	m = nil

	_, err := m.index(0, 0)
	if err.Error() != NilMatrixObject {
		t.Fatal("check nil object fail")
	}

	m = NewMatrix[int](3, 3)
	_, err = m.index(-1, 0)
	if err.Error() != InvalidIndexError {
		t.Fatal("check invalid index fail")
	}

	_, err = m.index(1, -1)
	if err.Error() != InvalidIndexError {
		t.Fatal("check invalid index fail")
	}

	_, err = m.index(9, 0)
	if err.Error() != InvalidIndexError {
		t.Fatal("check invalid index fail")
	}

	_, err = m.index(1, 8)
	if err.Error() != InvalidIndexError {
		t.Fatal("check invalid index fail")
	}
}

func Test_PosError(t *testing.T) {
	var m *Matrix[int]
	m = nil

	_, _, err := m.pos(0)
	if err.Error() != NilMatrixObject {
		t.Fatal("check nil object fail")
	}

	m = NewMatrix[int](3, 3)
	_, _, err = m.pos(-1)
	if err.Error() != InvalidIndexError {
		t.Fatal("check invalid index fail")
	}

	_, _, err = m.pos(11)
	if err.Error() != InvalidIndexError {
		t.Fatal("check invalid index fail")
	}
}
