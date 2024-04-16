package table

import (
	"testing"
)

func TestColumnNameCompare(t *testing.T) {
	columnName := []string{"hello", "world"}
	tableHeader := []string{"hello", "china", "beijing"}
	res := columnNameCompare(columnName, tableHeader)
	expected := []string{"world"}
	r := checksliceequal(res, expected)
	if r != 0 {
		t.Error("column set error, please input again!!")
	}
}

func TestSliceCompare(t *testing.T) {
	a := []string{"abcdefg", "345", "8888888888"}
	b := []string{"hello", "6666666", "5555555"}
	expected := []string{"abcdefg", "6666666", "8888888888"}
	res := sliceCompare(a, b)
	r := checksliceequal(res, expected)
	if r != 0 {
		t.Error("Failed to obtain the maximum index from two slices")
	}
}

func TestCheckColumnName(t *testing.T) {
	a := []string{"hello", "all"}
	expected := true
	r := checkColumnName(a)
	if r != expected {
		t.Errorf("The column name %s does not contain all", a)
	}
}
func TestGetColumnIndex(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	b := []string{"b", "d", "a", "f", "g"}
	expected := []int{2, 0, 1}
	res := getColumnIndex(a, b)
	r := checkintsliceequal(res, expected)
	if r != 0 {
		t.Errorf("Get index error!!!")
	}
}
func TestColumnInTableHeader(t *testing.T){
        a := "hello"
	b := []string{"abcd","world","hello"}
	expected := true
	r := columnInTableHeader(a,b)
        if r != expected {
             t.Errorf("The slice does not contain the %s field",a)
	}
}
