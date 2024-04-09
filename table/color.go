package table

/*
   **************************************************
   * Date:   2024-04-09                             *
   * Author: gongguan                               *
   * Email:  1542345123@qq.com                      *
   **************************************************
*/

// 定义前景色变量
var (
	fgroundcolor = []int{30, 31, 32, 33, 34, 35, 36, 37}
	bgroundcolor = []int{40, 41, 42, 43, 44, 45, 46, 47}
)

// 检查前景色
func checkFGroundColor(foregroundcolor int) bool {
	mapp := make(map[int]bool)
	for _, v := range fgroundcolor {
		if !mapp[v] {
			mapp[v] = true
		}
	}
	if mapp[foregroundcolor] {
		return true
	} else {
		return false
	}
}

// 检查背景色
func checkBGroundColor(backgroundcolor int) bool {
	mapp := make(map[int]bool)
	for _, v := range bgroundcolor {
		if !mapp[v] {
			mapp[v] = true
		}
	}
	if mapp[backgroundcolor] {
		return true
	} else {
		return false
	}
}
