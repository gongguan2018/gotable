package table

import (
	"testing"
	"reflect"
)

// 测四两个int类型切片是否相等
func TestCheckIntSliceEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	expected := 0
	r := checkintsliceequal(a, b)
	if r != expected {
		t.Error("Two int type slices are not equal!!!")
	}
}
//检查两个string类型切片是否相等
func TestCheckSliceEqual(t *testing.T) {
	a := []string{"hello", "world"}
	b := []string{"hello", "world"}
	expected := 0
	r := checksliceequal(a, b)
	if r != expected {
		t.Error("Two string type slices are not equal!!!")
	}
}
//检查一个切片在另一个切片中的索引
func TestCheckSliceIndex(t *testing.T) {
      a := []string{"hello","world"}
      b := []string{"abcd","hello","china","world"}
      expected := []int{1,3}
      r := checkSliceIndex(a,b)
      res := checkintsliceequal(r,expected)
      if res != 0 {
            t.Error("Get Slice index Error!!!")
      }
}
//检查二维切片中是否有重复元素，如果有将后面的切片中重复元素改为空字符串
//使用reflect.DeepEqual()来判断两个切片是否相等
func TestCheckRepeat(t *testing.T) {
	slices := [][]string{
            []string{"hello","china","world","guang","dong"},
	    []string{"hello","chinas","guang","dong","beijing"},
	}
	expected := [][]string{
            []string{"hello","china","world","guang","dong"},
	    []string{"","chinas","","","beijing"},
	}
	res := checkrepeat(slices)
	r := reflect.DeepEqual(expected,res)
	if !r {
            t.Error("Error checking for duplicate slices!!!")
	}
}
