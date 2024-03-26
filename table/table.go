package table

import (
	"fmt"
	"strings"
)

type Table struct {
	maxlen []string
}

// main中调用此函数
func CreateTable(tableheader []string, row [][]string) *Table {
	if len(row) == 0 {
		_ = createTableHeader(tableheader)
		return nil
	} else {
		maxLengthSlice := maxlength(tableheader, row)
		maxLengthSlice.CreateTableRow(tableheader, row)
                return maxLengthSlice
	}
        return nil
}

// 创建表和列
//定义二维切片,然后把表头的切片和列切片都添加进去
func (t *Table) CreateTableRow(th []string, row [][]string) {
       totalSlice  := make([][]string,0,100)
       totalSlice = append(totalSlice,th)
       for _,v := range row {
           totalSlice = append(totalSlice,v)
       }
       //根据二维切片中元素数量可以知道表一共有多少行,如果长度为3,那么表头一行,内容2行
       for i :=0; i<= len(totalSlice);i++ { 
              for j := 0; j < len(t.maxlen); j++ {
                        c := fmt.Sprintf("%s%-*s", "+", len(t.maxlen[j])+4, "-")
                        d := strings.Replace(c, " ", "-", -1)
                        if j == (len(t.maxlen) - 1) {
                                d = d + "+"
                        }
                        fmt.Printf(d)
                }
                fmt.Println()
                if i != len(totalSlice) {
                        //打印中间竖线,2表示格式化2个字符宽度,此处是用空格替代,这2个字符加文本的字符长度刚好等于上面的-的长度,因此不会变形
                        for k, v := range totalSlice[i] {
                            //每一个表格的长度都以当前列中最大长度的内容为准,不足的使用空格补全,因此使用len(t.maxlen[k])来补全 
                            e := fmt.Sprintf("%s%*s%-*s%-*s","|",2,"",len(t.maxlen[k]),v,2,"")
                                    if k == (len(totalSlice[i]) - 1) {
                                        e = e + "|"
                                     }
                                   fmt.Printf(e)
                        }
                        fmt.Println()
                }
                
       }
       //根据t.maxlen切片中的字节长度来创建表格(通过比对表头和列中数据的最大长度得来了,否则表格无法对齐)
       
}

// 对比表头切片与列切片,获取到对应索引位置的元素的最大值组成一个新切片
func maxlength(tableheader []string, row [][]string) *Table {
	for _, rowslice := range row {
		tableheader = compare(tableheader, rowslice)
	}
	maxlengths := &Table{
		maxlen: tableheader,
	}
	return maxlengths
}

// 对切片中元素长度进行比较,获取到每个索引对应的最大长度值组成新切片
// eg: [888888,2,3]和["hello","beijing","100"] 最后组成的新切片就是[888888,"beijing","100"]
func compare(th, row []string) []string {
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

// 创建表头
// len(a[i])+2表示获取表头内容的字节长度+2个字节
// %-*s中的*可以动态获取格式化长度,格式化长度默认不足使用空格补齐,这里将空格替换为-
func createTableHeader(a []string) *Table {
	for j := 0; j <= 1; j++ {
		//打印上面一行
		for i := 0; i < len(a); i++ {
			c := fmt.Sprintf("%s%-*s", "+", len(a[i])+2, "-")
			d := strings.Replace(c, " ", "-", -1)
			if i == (len(a) - 1) {
				d = d + "+"
			}
			fmt.Printf(d)
		}
		fmt.Println()
		if j != 1 {
			//打印中间竖线,2表示格式化2个字符宽度,此处是用空格替代,这2个字符加文本的字符长度刚好等于上面的-的长度,因此不会变形
			for k, v := range a {
				e := fmt.Sprintf("%s%-*s%s", "|", 2, " ", v)
				if k == (len(a) - 1) {
					e = e + "|"
				}
				fmt.Printf(e)
			}
			fmt.Println()
		}
	}
	var table Table
	return &table
}

