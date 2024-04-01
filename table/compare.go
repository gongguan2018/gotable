package table

// 查找表头中的列中是否包含输入的要设置的列,如果不包含,说明要设置的列错误,不存在
func columnNameCompare(columnName []string, tableHeader []string) []string {
	var columnNotExists []string
	tableMap := make(map[string]bool)
	for _, v := range tableHeader {
		if !tableMap[v] {
			tableMap[v] = true
		}
	}
	for _, v := range columnName {
		if !tableMap[v] {
			columnNotExists = append(columnNotExists, v)
		}
	}
	return columnNotExists
}

// 对切片中元素长度进行比较,获取到每个索引对应的最大长度值组成新切片
// eg: [888888,2,3]和["hello","beijing","100"] 最后组成的新切片就是[888888,"beijing","100"]
func sliceCompare(th, row []string) []string {
	res := []string{}
	for k, v := range th {
		if len(v) >= len(row[k]) {
			res = append(res, v)
		} else {
			res = append(res, row[k])
		}
	}
	return res
}

// 检查列名切片中是否包含all元素
func checkColumnName(columnName []string) bool {
	columnNameMap := make(map[string]bool)
	for _, v := range columnName {
		if !columnNameMap[v] {
			columnNameMap[v] = true
		}
	}
	if !columnNameMap["all"] {
		return false
	} else {
		return true
	}
}
func getRepeat(column, tableHeader []string) []string {
	newMap := make(map[string]int)
	var newSlice []string
	for k, v := range tableHeader {
		_, ok := newMap[v]
		if !ok {
			newMap[v] = k
		}
	}
	for _, v := range column {
		_, ok := newMap[v]
		if ok {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// 获取要设置的字段在表头中的索引
func getColumnIndex(column, tableHeader []string) []int {
	newMap := make(map[string]int)
	var newSlice []int
	for k, v := range tableHeader {
		_, ok := newMap[v]
		if !ok {
			newMap[v] = k
		}
	}
	for _, v := range column {
		index, ok := newMap[v]
		if ok {
			newSlice = append(newSlice, index)
		}
	}
	return newSlice
}

// 检查tableHeader中字段是否在与要设置的列名匹配
func columnInTableHeader(s string, column []string) bool {
	newMap := make(map[string]bool)
	for _, v := range column {
		if !newMap[v] {
			newMap[v] = true
		}
	}
	if !newMap[s] {
		return false
	}
	return true
}
func checkIndex(index int, columnIndex []int) bool {
	indexMap := make(map[int]bool)
	for _, v := range columnIndex {
		if !indexMap[v] {
			indexMap[v] = true
		}
	}
	if indexMap[index] {
		return true
	} else {
		return false
	}
}
