package table

import (
	"testing"
)

// 检查前景色
func TestCheckFGroundColor(t *testing.T) {
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	expected := true
	for _, v := range a {
		b := checkFGroundColor(v)
		if b != expected {
			t.Error("foregroundcolor set error!!!")
		}
	}
}

func TestCheckBGroundColor(t *testing.T) {
	a := []int{40, 41, 42, 43, 44, 45, 46, 47}
	expected := true
	for _, v := range a {
		b := checkBGroundColor(v)
		if b != expected {
			t.Error("backgroundcolor set error!!!")
		}
	}

}
