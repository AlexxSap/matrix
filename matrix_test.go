package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func compareSlices[T comparable](act, exp []T) error {
	if len(act) != len(exp) {
		return errors.New("different len")
	}

	for i := 0; i < len(act); i++ {
		if act[i] != exp[i] {
			return fmt.Errorf("act: %v exp: %v at index: %d", act[i], exp[i], i)
		}
	}

	return nil
}

func Test_Index(t *testing.T) {
	m := NewZeroMatrix[int](3, 3)

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
	m := NewZeroMatrix[int](3, 3)

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

	m = NewZeroMatrix[int](3, 3)
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

	m = NewZeroMatrix[int](3, 3)
	_, _, err = m.pos(-1)
	if err.Error() != InvalidIndexError {
		t.Fatal("check invalid index fail")
	}

	_, _, err = m.pos(11)
	if err.Error() != InvalidIndexError {
		t.Fatal("check invalid index fail")
	}
}

func TestNewMatrixError(t *testing.T) {
	d := make([]int, 5)
	_, err := NewMatrix(d, 3, 2)

	if err.Error() != InvalidMatrixSize {
		t.Error("check invalid matrix size fail")
	}
}

func TestRowData(t *testing.T) {

	var m *Matrix[int]
	m = nil

	_, err := m.RowData(2)
	if err.Error() != NilMatrixObject {
		t.Fatal("check nil object fail")
	}

	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m, err = NewMatrix(d, 3, 3)
	if err != nil {
		t.Error(err)
	}

	_, err = m.RowData(-1)
	if err.Error() != InvalidIndexError {
		t.Error("check invalid index fail")
	}

	_, err = m.RowData(3)
	if err.Error() != InvalidIndexError {
		t.Error("check invalid index fail")
	}
	expRow := []int{1, 2, 3}
	row, err := m.RowData(0)
	if err != nil {
		t.Error(err)
	}
	if cmpRes := compareSlices(row, expRow); cmpRes != nil {
		t.Error(cmpRes)
	}

	strs := []string{"1", "2", "3", "4", "some 5", "or 6", "7", "8", "9"}
	ms, err := NewMatrix(strs, 3, 3)
	if err != nil {
		t.Error(err)
	}

	expStrings := []string{"4", "some 5", "or 6"}
	rowStrings, err := ms.RowData(1)
	if err != nil {
		t.Error(err)
	}
	if cmpRes := compareSlices(rowStrings, expStrings); cmpRes != nil {
		t.Error(cmpRes)
	}
}

func TestColumnData(t *testing.T) {

	var m *Matrix[int]
	_, err := m.ColumnData(2)
	if err.Error() != NilMatrixObject {
		t.Fatal("check nil object fail")
	}

	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m, err = NewMatrix(d, 3, 3)
	if err != nil {
		t.Error(err)
	}

	_, err = m.ColumnData(-1)
	if err.Error() != InvalidIndexError {
		t.Error("check invalid index fail")
	}

	_, err = m.ColumnData(3)
	if err.Error() != InvalidIndexError {
		t.Error("check invalid index fail")
	}
	expRow := []int{1, 4, 7}
	col, err := m.ColumnData(0)
	if err != nil {
		t.Error(err)
	}
	if cmpRes := compareSlices(col, expRow); cmpRes != nil {
		t.Error(cmpRes)
	}

	strs := []string{"1", "2", "3", "4", "some 5", "or 6", "7", "8", "9"}
	ms, err := NewMatrix(strs, 3, 3)
	if err != nil {
		t.Error(err)
	}

	expStrings := []string{"2", "some 5", "8"}
	colStrings, err := ms.ColumnData(1)
	if err != nil {
		t.Error(err)
	}
	if cmpRes := compareSlices(colStrings, expStrings); cmpRes != nil {
		t.Error(cmpRes)
	}
}
func TestAllOfRow(t *testing.T) {

	var m *Matrix[int]
	_, err := m.AllOfRow(2, func(cell int) bool { return false })
	if err.Error() != NilMatrixObject {
		t.Fatal("check nil object fail")
	}

	d := []int{1, 2, 3, 4, 8, 12, 7, 8, 9}
	m, err = NewMatrix(d, 3, 3)
	if err != nil {
		t.Error(err)
	}

	f := func(val int) bool {
		return val%3 == 0
	}

	_, err = m.AllOfColumn(-1, f)
	if err.Error() != InvalidIndexError {
		t.Error("check invalid index fail")
	}
	_, err = m.AllOfColumn(11, f)
	if err.Error() != InvalidIndexError {
		t.Error("check invalid index fail")
	}

	actual, err := m.AllOfColumn(1, f)
	if err != nil {
		t.Error(err)
	}

	if actual {
		t.Errorf("act: %t exp: %t", actual, false)
	}

	actual, err = m.AllOfColumn(2, f)
	if err != nil {
		t.Error(err)
	}

	if !actual {
		t.Errorf("act: %t exp: %t", actual, true)
	}

}

func TestShiftRowsDown(t *testing.T) {
	var m *Matrix[int]
	err := m.ShiftRowsDown()
	if err.Error() != NilMatrixObject {
		t.Fatal("check nil object fail")
	}

	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m, err = NewMatrix(d, 3, 3)
	if err != nil {
		t.Error(err)
	}

	err = m.ShiftRowsDown()
	if err != nil {
		t.Error(err)
	}

	exp := []int{0, 0, 0, 1, 2, 3, 4, 5, 6}
	if cmpRes := compareSlices(m.cells, exp); cmpRes != nil {
		t.Error(cmpRes)
	}

	err = m.ShiftRowsDown()
	if err != nil {
		t.Error(err)
	}

	exp = []int{0, 0, 0, 0, 0, 0, 1, 2, 3}
	if cmpRes := compareSlices(m.cells, exp); cmpRes != nil {
		t.Error(cmpRes)
	}

	err = m.ShiftRowsDown()
	if err != nil {
		t.Error(err)
	}

	exp = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	if cmpRes := compareSlices(m.cells, exp); cmpRes != nil {
		t.Error(cmpRes)
	}
}

func TestSet(t *testing.T) {
	var m *Matrix[int]
	err := m.Set(1, 1, 1)
	if err.Error() != NilMatrixObject {
		t.Fatal("check nil object fail")
	}

	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m, err = NewMatrix(d, 3, 3)
	if err != nil {
		t.Error(err)
	}

	m.Set(1, 1, 666)
	exp := []int{1, 2, 3, 4, 666, 6, 7, 8, 9}
	if cmpRes := compareSlices(m.cells, exp); cmpRes != nil {
		t.Error(cmpRes)
	}

	m.Set(1, 0, -666)
	exp = []int{1, 2, 3, -666, 666, 6, 7, 8, 9}
	if cmpRes := compareSlices(m.cells, exp); cmpRes != nil {
		t.Error(cmpRes)
	}
}
