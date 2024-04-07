package table

import (
	"fmt"
)
//设置全部列的对齐方式
func setAllColumnAlign(align int, tableAll []string) {
	var tdslice [][]string
	tds := croppSlice(tdslice, tableAll)
	res := checkrepeat(tds) //res是一个二维切片
	for _, v := range res {
		for k, v1 := range v {
			e := ""
			switch align {
			case leftAlignment:
				if v1 == "" {
					e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
				} else if len(v1) == 20 {
					e = fmt.Sprintf("%s%s%-*s", verticalLine, v1, 4, " ")
				} else {
					e = fmt.Sprintf("%s%s%-*s", verticalLine, v1, 24-len(v1), "")
				}
			case rightAlignment:
				if v1 == "" {
					e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
				} else if len(v1) == 20 {
					e = fmt.Sprintf("%s%-*s%s", verticalLine, 4, " ", v1)
				} else {
					e = fmt.Sprintf("%s%-*s%s", verticalLine, 24-len(v1), " ", v1)
				}
			case centerAlignment:
				if v1 == "" {
					e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
				} else if len(v1) == 20 {
					e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 2, " ")
				} else {
					e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 22-len(v1), " ")
				}
			}
			if k == (len(tableAll) - 1) {
				e = e + verticalLine
			}
			fmt.Printf(e)
		}
		fmt.Println()
	}

}

// 设置表中指定列的对齐方式(包含表头和表数据)
func setSpecifyColumnAlign(tableColumn []string,align int ,columnIndex []int){
	var tdslice [][]string
	tds := croppSlice(tdslice, tableColumn)
	res := checkrepeat(tds)
	for _, v := range res {
		for k1, v1 := range v {
			e := ""
			b := checkIndex(k1, columnIndex)
			switch align {
			case leftAlignment:
				if b {
					if v1 == "" {
						e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
					} else if len(v1) == 20 {
						e = fmt.Sprintf("%s%s%-*s", verticalLine, v1, 4, " ")
					} else {
						e = fmt.Sprintf("%s%s%-*s", verticalLine, v1, 24-len(v1), "")
					}
				} else {
					if v1 == "" {
						e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
					} else if len(v1) == 20 {
						e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 2, " ")
					} else {
						e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 22-len(v1), " ")
					}
				}
			case rightAlignment:
				if b {
					if v1 == "" {
						e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
					} else if len(v1) == 20 {
						e = fmt.Sprintf("%s%-*s%s", verticalLine, 4, " ", v1)
					} else {
						e = fmt.Sprintf("%s%-*s%s", verticalLine, 24-len(v1), " ", v1)
					}
				} else {
					if v1 == "" {
						e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
					} else if len(v1) == 20 {
						e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 2, " ")
					} else {
						e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 22-len(v1), " ")
					}
				}
			case centerAlignment:
				if v1 == "" {
					e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
				} else if len(v1) == 20 {
					e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 2, " ")
				} else {
					e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 22-len(v1), " ")
				}
			}
			if k1 == (len(v) - 1) {
				e = e + verticalLine
			}
			fmt.Printf(e)
		}
		fmt.Println()
	}
}
