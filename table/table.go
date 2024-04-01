package table

import (
	"fmt"
	"os"
	"strings"
)

type Table struct {
	tableHeader []string //表头内容
	maxlen      []string //表头和内容对应索引位置的值的最大长度组成的新切片,后续使用此长度创建表的宽度,这样才会对齐
	columnName  []string //表列名
	columnColor string   //列的颜色
	alignment   int      //对齐方式
}

// 初始化表
func InitTable(tableHeader []string) *Table {
	if len(tableHeader) != 0 {
		table := &Table{
			tableHeader: tableHeader,
		}
		return table
	} else {
		return &Table{} 
	}
}

// 设置对齐方式
func (t *Table) SetAlignment(alignment int, columnName ...string) {
	if len(t.tableHeader) != 0 {
	var columnNameSlice []string
	//如果要设置的列名为空,那么默认就是设置全部列名(不定参数的columnName实际为切片)
	if len(columnName) == 0 {
		align := checkColumnAlignment(alignment)
		if align == 5 {
			fmt.Println("Alignment parameter error,Can only select 0,1,2")
			return
		} else {
			t.alignment = align
			columnNameSlice = append(columnNameSlice, "all")
			t.columnName = columnNameSlice
		}

	} else {
		/*
			                          columnName不为空有可能出现多种情况,比如只输入了all,或者同时输入了all和其他列名或者没输入all而输入了其他列名
						  在这里我们设置只要切片中出现了all元素,那么就默认为设置全部列名
		*/
		b := checkColumnName(columnName)
		//如果b为true,说明columnName中是包含all的,因此视为设置全部的列
		if b {
			align := checkColumnAlignment(alignment)
			if align == 5 {
				fmt.Println("Alignment parameter error,Can only select 0,1,2")
				return
			} else {
				t.alignment = align
				columnNameSlice = append(columnNameSlice, "all")
				t.columnName = columnNameSlice

			}

		} else {
			//比较要设置的列名和表头中的列,是否有输入错误,返回包含错误列的切片
			s := columnNameCompare(columnName, t.tableHeader)
			if len(s) != 0 {
				fmt.Printf("Column Name %v Input Error,Please input again!!!\n", s)
				os.Exit(-1)
			} else {
				//len(s)等于0,此时要设置的列是正确的,开始设置
				align := checkColumnAlignment(alignment)
				if align == 5 {
					fmt.Println("Alignment parameter error,Can only select 0,1,2")
					return
				} else {
					t.alignment = align
					t.columnName = columnName
				}

			}
		}
	}
    }else {
	    fmt.Println("Error:","The table header content is empty,please check!")
	   os.Exit(-1) 
    }
}

// 设置颜色
func (t *Table) SetColor(columnName, columnColor string) {

}

// main中调用此函数
func (t *Table) CreateTable(tableheader []string, row [][]string) {
	//如果二维切片长度不为0说明除了表头外有数据,执行else逻辑
	if len(row) == 0 {
		t.createTableHeader(tableheader)
	} else {
		//取tableheader和row中切片的相同索引的值的最大长度
		maxLengthSlice := t.maxlength(tableheader, row)
		t.maxlen = maxLengthSlice
		t.CreateTableRow(tableheader, row)
	}
}

// 创建表和列
// 定义二维切片,然后把表头的切片和列切片都添加进去
func (t *Table) CreateTableRow(th []string, row [][]string) {
	//二维切片,把表头切片和内容切片都添加进来
	totalSlice := make([][]string, 0, 100)
	totalSlice = append(totalSlice, th)
	for _, v := range row {
		totalSlice = append(totalSlice, v)
	}
	//根据二维切片中元素数量可以知道表一共有多少行,如果长度为3,那么表头一行,内容2行
	for i := 0; i <= len(totalSlice); i++ {
		//打印表格中最上面的"-",表格的宽度要以manlen切片中每个索引的值的最大长度为宽,否则第一行内容10个字符,第二行20个字符就会出现无法对齐
		//maxlen切片中元素的个数也就是表的列的个数
		for j := 0; j < len(t.maxlen); j++ {
			//根据maxlen中的长度来创建表格,+4表示多加了4个空格
			c := fmt.Sprintf("%s%-*s", "+", len(t.maxlen[j])+4, "-")
			d := strings.Replace(c, " ", "-", -1)
			if j == (len(t.maxlen) - 1) {
				d = d + "+"
			}
			fmt.Printf(d)
		}
		fmt.Println()
		//当i的值等于len(totalSlice)的时候是不打印"|"的,eg: total长度为3,那么表头一行,内容两行,那么实际要打印四行的"-"和3行的"|"
		//因此当i=3(索引从0开始)的时候,只打印"-",而不打印"|"
		if i != len(totalSlice) {
				if len(t.columnName) == 1 && strings.TrimSpace(t.columnName[0]) == "all" {
					//设置整个表的列对齐方式(包括表头和表内容)
					setAllColumnAlignment(t.alignment, t.maxlen,totalSlice[i]) 
				} else {
					//给指定的列设置对齐方式
					setSpecifyColumnAlignment(t.columnName, t.alignment,t.maxlen, th,totalSlice[i])
				}

			fmt.Println()
		}

	}
	//根据t.maxlen切片中的字节长度来创建表格(通过比对表头和列中数据的最大长度得来了,否则表格无法对齐)

}

/*  
   对比表头切片与列切片,获取到对应索引位置的元素的最大长度值组成一个新切片
   第一轮循环后将比较结果依旧赋值给tableheader,这样tableheader的切片就变成了第一次比较后的结果,然后再次进行比较
   即使row有无数个切片，最终也会获取出最大值

*/
func (t *Table) maxlength(tableheader []string, row [][]string) []string {
	for _, rowslice := range row {
		tableheader = sliceCompare(tableheader, rowslice) //比较结果依旧赋值给tableheader,然后进行下一轮
	}
	return tableheader
}

// 创建表头,在没有表内容只有表头的时候调用此函数
func (t *Table) createTableHeader(tableHeader []string) {
	for j := 0; j <= 1; j++ {
		//此for循环主要用于打印"------"
		for i := 0; i < len(tableHeader); i++ {
			c := fmt.Sprintf("%s%-*s", "+", len(tableHeader[i])+4, "-")
			d := strings.Replace(c, " ", "-", -1)
			if i == (len(tableHeader) - 1) {
				d = d + "+"
			}
			fmt.Printf(d)
		}
		fmt.Println()
		if j != 1 {
			if len(t.columnName) == 1 && strings.TrimSpace(t.columnName[0]) == "all" {
				//设置表头全部列对齐方式
				setTableHeadAllColAlign(t.alignment, tableHeader)
			} else {
				//给指定的列设置对齐方式
				setTableHeaderColumnAlignment(t.columnName, t.alignment, tableHeader)
			}
			fmt.Println()
		}
	}
}
