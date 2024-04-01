package table

import (
	"fmt"
)

// 检测对齐方式是否有错误,leftAlignment、rightAlignment、centerAlignment都是常量
func checkColumnAlignment(alignment int) int {
	switch alignment {
	case leftAlignment:
		return leftAlignment
	case rightAlignment:
		return rightAlignment
	case centerAlignment:
		return centerAlignment
	default:
		return 5
	}

}

// 设置表头全部列的对齐方式,不包含表内容
// 下面的for循环用于打印"|" 和表头内容
// 打印中间竖线,4表示格式化4个字符宽度,此处是用空格替代,这4个字符加文本的字符长度刚好等于上面的-的长度,因此不会变形
func setTableHeadAllColAlign(align int, tableHeader []string) {
	//获取表头中每个元素的长度,如果长度超过20,那么就要分行
	for k, v := range tableHeader {
			e := ""
			switch align {
			case leftAlignment:
				e = fmt.Sprintf("%s%s%-*s", "|", v, 4, " ")
			case rightAlignment:
				e = fmt.Sprintf("%s%-*s%s", "|", 4, " ", v)
			case centerAlignment:
				e = fmt.Sprintf("%s%-*s%s%-*s", "|", 2, " ", v, 2, " ")
			}
			if k == (len(tableHeader) - 1) {
				e = e + "|"
			}
			fmt.Printf(e)
	}
}

// 设置整个表的指定列的对齐方式(包含表头和表内容)
func setAllColumnAlignment(align int, maxlen, tableRow []string) {
	for k, v := range tableRow {
		e := ""
		switch align {
		case leftAlignment:
			e = fmt.Sprintf("%s%-*s%-*s", "|", len(maxlen[k]), v, 4, "")
		case rightAlignment:
			e = fmt.Sprintf("%s%-*s%*s", "|", 4, " ", len(maxlen[k]), v)
		case centerAlignment:
			e = fmt.Sprintf("%s%-*s%-*s%-*s", "|", 2, " ", len(maxlen[k]), v, 2, "")
		}
		if k == (len(tableRow) - 1) {
			e = e + "|"
		}
		fmt.Printf(e)
	}
}

/*
   设置表头指定列的对齐方式,可能一个列或者多个列,可能是第一例和最后一列要设置,也可能是第二列和最后一列要设置
   注: 这个函数设置的列只是表头,没有表内容,也就是在main.go中没有输入row,只有一个tableHeader
*/

func setTableHeaderColumnAlignment(column []string, align int, tableHeader []string) {
	for k, v := range tableHeader {
		e := ""
		switch align {
		case leftAlignment:
			//检查v是否在column切片中,如果在表示当前v的值就是要设置的列
			b := columnInTableHeader(v, column)
			if b {
				e = fmt.Sprintf("%s%s%-*s", "|", v, 4, " ")
			} else {
				e = fmt.Sprintf("%s%-*s%s%-*s", "|", 2, " ", v, 2, " ")
			}
		case rightAlignment:
			b := columnInTableHeader(v, column)
			if b {
				e = fmt.Sprintf("%s%-*s%s", "|", 4, " ", v)
			} else {
				e = fmt.Sprintf("%s%-*s%s%-*s", "|", 2, " ", v, 2, " ")
			}
		case centerAlignment:
			e = fmt.Sprintf("%s%-*s%s%-*s", "|", 2, " ", v, 2, " ")
		}
		if k == (len(tableHeader) - 1) {
			e = e + "|"
		}
		fmt.Printf(e)
	}
}

/*
这个函数设置的整个表格的指定列对齐方式,包括表头和表内容,可能设置一列也可能设置多列
因为我们传递过来的column是在main中设置的列名,也就是表头中的名字,但是实际我们还有表格内容,因此就不能只根据这个列名来设置对齐
否则会出现第一列的表头被设置了左对齐,但是下面的内容没变化,还是居中,因此此时就要根据索引来设置了
tableRow是每次循环totalSlice传递过来的切片,每个tableRow都代表表中的一行,第一行是表头,第二行是表的内容第一行,以此类推
因此,可以先获取要设置的列在表头中的索引,tableRow通过此索引的获取的值也就是要设置的对齐的值
*/
func setSpecifyColumnAlignment(column []string, align int, maxlen, tableHeader, tableRow []string) {
	//获取column元素在tableHeader中的索引
	columnIndex := getColumnIndex(column, tableHeader)
	for k, v := range tableRow {
		e := ""
		switch align {
		case leftAlignment:
			//判断k是否在columnIndex中,如果在,说明此时的v就是要设置的列
			b := checkIndex(k, columnIndex)
			if b {
				e = fmt.Sprintf("%s%-*s%-*s", "|", len(maxlen[k]), v, 4, "")
			} else {
				e = fmt.Sprintf("%s%-*s%-*s%-*s", "|", 2, " ", len(maxlen[k]), v, 2, "")
			}
		case rightAlignment:
			b := checkIndex(k, columnIndex)
			if b {
				e = fmt.Sprintf("%s%-*s%*s", "|", 4, " ", len(maxlen[k]), v)
			} else {
				e = fmt.Sprintf("%s%-*s%-*s%-*s", "|", 2, " ", len(maxlen[k]), v, 2, "")
			}
		case centerAlignment:
			e = fmt.Sprintf("%s%-*s%-*s%-*s", "|", 2, " ", len(maxlen[k]), v, 2, "")
		}
		if k == (len(tableRow) - 1) {
			e = e + "|"
		}
		fmt.Printf(e)
	}
}
