# gotable
gotable
这是一个使用Golang实现的表格库,可以实现一些简单的表格展示,功能如下:

1、表格内容换行： 当表格中文本内容超过20个字符的时候,将自动换行

2、表格内容对齐方式, 如下:

(1)、SetAlignmentMode(0) 表示左对齐

(2)、SetAlignmentMode(1) 表示右对齐

(3)、SetAlignmentMode(2) 表示剧中对齐

(4)、SetAlignmentMode()  什么都不设置默认表示居中对齐

3、设置表格中要对齐的列: 

(1)、SetAlignmentColumn()                       什么都不输入默认表示设置全部列对齐

(2)、SetAlignmentColumn("MemTotal","all")       包含"all",默认也是设置全部列对齐,不能为空字符串

(3)、SetAlignmentColumn("all")                  只输入"all",默认也是设置全部列对齐,不能为空字符串

(4)、SetAlignmentColumn("MemTotal","MemFree")   设置MemTotal、MemFree对齐,不能为空字符串

4、设置表格中要设置颜色的列

(1)、SetColorColumn()                           什么都不输入默认表示设置全部列的颜色

(2)、SetColorColumn("all")                      只输入"all",表示设置全部列的颜色,不能为空字符串

(3)、SetColorColumn("MemTotal","all")           包含"all", 表示设置全部列的颜色,不能为空字符串

(4)、SetColorColumn("MemTotal","MemFree")       设置MemTotal, MemFree列的颜色,不能为空字符串

5、设置表格中文字的前景色,也就是字体的颜色

(1)、SetForegroundColor()                       什么都不输入默认前景色为37,白色字体

(2)、SetForegroundColor(30)                     黑色

(3)、SetForegroundColor(31)                     红色

(4)、SetForegroundColor(32)                     绿色

(5)、SetForegroundColor(33)                     黄色

(6)、SetForegroundColor(34)                     蓝色

(7)、SetForegroundColor(35)                     紫红色

(8)、SetForegroundColor(36)                     青蓝色

(9)、SetForegroundColor(37)                     白色

6、设置文字背景色

(1)、SetBackgroundColor()                       什么都不输入默认背景色为40,黑色背景

(2)、SetForegroundColor(40)                     黑色

(3)、SetForegroundColor(41)                     红色

(4)、SetForegroundColor(42)                     绿色

(5)、SetForegroundColor(43)                     黄色

(6)、SetForegroundColor(44)                     蓝色

(7)、SetForegroundColor(45)                     紫红色

(8)、SetForegroundColor(46)                     青蓝色

(9)、SetForegroundColor(47)                     白色

演示例子1: MemTotal和MemUsedHelloworldchina设置字体绿色背景黑色,MemFree左对齐

![image](https://github.com/gongguan2018/gotable/assets/40058594/07ee146b-340c-43e4-bafb-a48b8a183da6)

演示例子2: 设置全部列为绿色字体红色背景,MemFree左对齐

![微信截图_20240409115146](https://github.com/gongguan2018/gotable/assets/40058594/a3bee646-8a9a-4e6a-9f14-c9c139983cd2)

演示例子3: 设置Memfree右对齐, MemTotal列为红色背景绿色字体


![20240409115723](https://github.com/gongguan2018/gotable/assets/40058594/c5e741d1-2ed1-4af3-abbe-fdde8db7d53f)


用法:

1、首先下载gotable库,命令如下：

```
   go get github.com/gongguan2018/gotable/table
```

或 如果通过go mod管理的项目， 可以在编辑完代码后执行如下命令也是可以的：

```
go mod tidy
```

2、创建一个table.go文件, 内容如下：

```
package main

import (
        "github.com/gongguan2018/gotable/table"
)

func main() {
        tableHead := []string{"country", "city", "street"}
        row := [][]string{}
        row1 := []string{"china", "shenzhen", "fuyong"}
        row2 := []string{"china", "daqing", "ranghulu"}
        row = append(row, row1)
        row = append(row, row2)
        //初始化表结构
        t := table.InitTable(tableHead)
        //设置对齐模式
        t.SetAlignmentMode()
        //设置对齐的列,默认所有
        t.SetAlignmentColumn()
        //设置要改变颜色的列,默认所有
        t.SetColorColumn()
        /*
                            设置前景色(非必填,默认为37),可选参数为:
                            30(黑色),31(红色),32(绿色),33(黄色),34(蓝色),35(紫红色),36(青蓝色),37(白色)
        */
        t.SetForegroundColor()
        /*
                            设置背景色(非必填,默认为40),可选参数为:
                            40(黑色),41(红色),42(绿色),43(黄色),44(蓝色),45(紫红色),46(青蓝色),47(白色)
        */
        t.SetBackgroundColor()
        //创建表
        t.CreateTable(tableHead, row)
}

```
3、最后运行table.go, 如下：

```
go run table.go
```
