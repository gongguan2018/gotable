package main

import (
	"github.com/gongguan2018/gotable/table"
)

func main() {
	tableHead := []string{"MemTotal", "MemUsed", "MemFree", "SwapTotalhelloworldchina", "SwapUsed", "SwapFree"}
	row := [][]string{}
	row1 := []string{"guangdongshenzhen", "212345678901112131415161711111111111111111111", "3", "4", "5", "6"}
	row2 := []string{"chinabeijing", "8", "33", "i664", "heilongjiangshengdaqing", "6"}
	row3 := []string{"shanghai", "linux", "ubuntu", "amd64", "windows", "apple"}
	row = append(row, row1)
	row = append(row, row2)
	row = append(row, row3)
	//初始化
	t := table.InitTable(tableHead)
	//设置列名对齐方式(必选),0:left,1:right,2:center
	/*
            列名支持多种方式:
	    1、什么都不输入:           t.SetAlignment(0),表示将全部列设置为左对齐
	    2、只输入"all":            t.SetAlignment(0,"all"),也表示设置全部列为左对齐
	    3、输入的列名中包含"all":  t.SetAlignment(0,"all","MemTotal"),也是设置全部列为左对齐
	    4、不包含"all",包含指定列: t.SetAlignment(0,"MemTotal","MemFree"),设置对应列为左对齐
	*/
	t.SetAlignment(0, "MemUsed")
	//设置颜色
	//	t.SetColor()
	t.CreateTable(tableHead, row)
}
