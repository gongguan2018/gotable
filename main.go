package main

import (
	"fmt"
	"github.com/gongguan2018/gotable/table"
)

func main() {
	tableHead := []string{"MemTotal", "MemUsed", "MemFree", "SwapTotal", "SwapUsed", "SwapFree"}
	row := [][]string{}
	row1 := []string{"guangdongshenzhen", "2123456789011121314151617", "3", "4", "5", "6"}
	row2 := []string{"chinabeijing", "8", "33", "i664", "heilongjiangshengdaqing", "6"}
	row3 := []string{"shanghai", "linux", "ubuntu", "amd64", "windows", "apple"}
	row = append(row, row1)
	row = append(row, row2)
	row = append(row, row3)
	//创建表,向函数传递表头和列
	if len(tableHead) != 0 {
		_ = table.CreateTable(tableHead, row)
	} else {
		fmt.Println("The table header content is empty,please check!")
	}
}
