package main

import (
	"github.com/gongguan2018/gotable/table"
)
/*
   **************************************************
   * Date:   2024-04-09                             *
   * Author: gongguan                               *
   * Email:  1542345123@qq.com                      *
   **************************************************
*/
func main() {
	tableHead := []string{"MemTotal", "MemUsedhelloworldchina", "Memfree", "SwapTotal", "SwapUsed", "SwapFree"}
	row := [][]string{}
	row1 := []string{"guangdongshenzhen", "212345678901112131415161711111111", "javagocpythonk8sdocker", "4", "5", "6"}
	row2 := []string{"chinabeijing", "8", "33", "i664", "heilongjiangshengdaqing", "6"}
	row3 := []string{"shanghai", "linux", "ubuntu", "amd64", "windows", "apple"}
	row = append(row, row1)
	row = append(row, row2)
	row = append(row, row3)
	//初始化
	t := table.InitTable(tableHead)
	/*
	    设置列名对齐方式(非必填,默认为居中),可选择:
	    0:left,1:right,2:center(近似居中)
	*/
	t.SetAlignmentMode(1)
	/*
            设置列名的对齐方式(非必填),支持多种方式:
	    1、什么都不输入:           t.SetAlignmentColumn(),                     表示将全部列设置对齐
	    2、只输入"all":            t.SetAlignmentColumn("all"),                表示设置全部列对齐
	    3、输入的列名中包含"all":  t.SetAlignmentColumn("all","MemTotal"),     表示设置全部列对齐
	    4、不包含"all",包含指定列: t.SetAlignmentColumn("MemTotal","MemFree"), 表示设置对应列对齐
	    注: 只能输入表头中的元素,不可输入表内容中的元素
	*/
	t.SetAlignmentColumn("Memfree")
	/*
	    SetColorColumn(): 设置将哪个列设置颜色,非必填,有多种设置方式,如下:
	    1、什么都不输入:              SetColorColumn(),                     表示设置所有列的颜色
	    2、只输入"all":               SetColorColumn("all"),                表示设置所有列的颜色 
	    3、输入的列中包含"all":       SetColorColumn("MemTotal","all")      表示设置所有列的颜色
	    4、不包含"all",包含指定列:    SetColorColumn("MemTotal","SwarUsed") 表示设置指定列的颜色
	    注: 只能输入表头中的元素,不可输入表内容中的元素
	*/
	t.SetColorColumn("MemTotal")
	/*
            设置前景色(非必填,默认为37),可选参数为:
	    30(黑色),31(红色),32(绿色),33(黄色),34(蓝色),35(紫红色),36(青蓝色),37(白色)
	*/
	t.SetForegroundColor(32)
	/*
            设置背景色(非必填,默认为40),可选参数为:
	    40(黑色),41(红色),42(绿色),43(黄色),44(蓝色),45(紫红色),46(青蓝色),47(白色)
	*/
	t.SetBackgroundColor(41)
	//创建表
	t.CreateTable(tableHead, row)
}
