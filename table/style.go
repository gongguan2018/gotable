package table


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
func setTableHeadAllColAlign(align int, tableHeader []string,colorColumnIndex []int,foregroundcolor,backgroundcolor int) {
        setAllColumnAlign(align,foregroundcolor,backgroundcolor,tableHeader,colorColumnIndex)
}

// 设置整个表的指定列的对齐方式(包含表头和表内容)
func setAllColumnAlignment(align int,  tableRow []string,colorColumnIndex []int,foregroundcolor,backgroundcolor int) {
	setAllColumnAlign(align,foregroundcolor,backgroundcolor,tableRow,colorColumnIndex)
}

/*
   设置表头指定列的对齐方式,可能一个列或者多个列,可能是第一例和最后一列要设置,也可能是第二列和最后一列要设置
   注: 这个函数设置的列只是表头,没有表内容,也就是在main.go中没有输入row,只有一个tableHeader
*/

func setTableHeaderColumnAlignment(tableHeader []string, align,foregroundcolor,backgroundcolor int, columnIndex,colorColumnIndex []int) {
        setSpecifyColumnAlign(tableHeader,align,foregroundcolor,backgroundcolor,columnIndex,colorColumnIndex)
}

/*
这个函数设置的整个表格的指定列对齐方式,包括表头和表内容,可能设置一列也可能设置多列
因为表中除了表头还有表格内容,因此就不能只根据列名来设置对齐方式,否则会出现第一列的表头被设置了左对齐,
但是下面的内容没变化,还是居中,因此此时就要根据索引来设置了,tableRow是每次循环totalSlice传递过来的切片,
每个tableRow都代表表中的一行,第一行是表头,第二行是表的内容第一行,以此类推
因此,可以先获取要设置的列在表头中的索引,tableRow通过此索引的获取的值也就是要设置的对齐的值
*/
func setSpecifyColumnAlignment(tableRow []string,align,foregroundcolor,backgroundcolor int,columnIndex,colorColumnIndex []int) {
        setSpecifyColumnAlign(tableRow,align,foregroundcolor,backgroundcolor,columnIndex,colorColumnIndex) 
}
