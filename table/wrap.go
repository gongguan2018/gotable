package table

import (
)
//此函数执行换行操作
func wrap(str string) []string {
	var newstr string
	var strSlice []string
	divide := len(str) / 20
	for i := 0; i <= divide; i++ {
		if i == divide {
			newstr = str[0:]
		} else {
			newstr = str[0:15]
			str = str[16:]
		}
		strSlice = append(strSlice, newstr)
	}
	return strSlice 
}
