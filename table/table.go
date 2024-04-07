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
	} else {
		fmt.Println("Error:", "The table header content is empty,please check!")
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
	//	maxLengthSlice := t.maxlength(tableheader, row)
	//	t.maxlen = maxLengthSlice
		t.CreateTableRow(tableheader, row)
	}
}

// 创建表和列
// 定义二维切片,然后把表头的切片和列切片都添加进去
func (t *Table) CreateTableRow(tableHeader []string, row [][]string) {
	/*
              获取要设置的列在tableHeader中索引,为什么这个调用要放在这里呢,如果放在setSpecifyColumnAlignment函数或者后面的调用
	      中,由于需求,在后面的调用(croppSlice函数)有设置tableHeader为空,当tableHeader为空的时候,获取到的columnIndex也为空
	      当我们设置表中多个列向左对齐的时候,不只是表头向左对齐,内容也要向左对齐,此时是根据索引来设置的,表头和表内容的相应
	      位置索引一样,因此必须保证columnIndex的内容全局不可变,否则可能出现表头左对齐了,但是内容还是居中的情况
	*/
	columnIndex := checkSliceIndex(t.columnName, tableHeader)
	//二维切片,把表头切片和内容切片都添加进来
	var totalSlice [][]string
	totalSlice = append(totalSlice, tableHeader)
	for _, v := range row {
		totalSlice = append(totalSlice, v)
	}
	//根据二维切片中元素数量可以知道表一共有多少行,如果长度为3,那么表头一行,内容2行
	for i := 0; i <= len(totalSlice); i++ {
		//打印表格中最上面的Horizontalbar,表格的宽度要以manlen切片中每个索引的值的最大长度为宽,否则第一行内容10个字符,第二行20个字符就会出现无法对齐
		for j := 0; j < len(totalSlice[0]); j++ {
			c := fmt.Sprintf("%s%-*s", plussign, 24, Horizontalbar)
			d := strings.Replace(c, " ", Horizontalbar, -1)
			if j == (len(totalSlice[0]) - 1) {
				d = d + plussign
			}
			fmt.Printf(d)
		}
		fmt.Println()
		//当i的值等于len(totalSlice)的时候是不打印"|"的,eg: total长度为3,那么表头一行,内容两行,那么实际要打印四行的Horizontalbar和3行的"|"
		//因此当i=3(索引从0开始)的时候,只打印Horizontalbar,而不打印"|"
		if i != len(totalSlice) {
			if len(t.columnName) == 1 && strings.TrimSpace(t.columnName[0]) == "all" {
				//设置整个表的列对齐方式(包括表头和表内容)
				setAllColumnAlignment(t.alignment, totalSlice[i])
			} else {
				//给指定的列设置对齐方式
				setSpecifyColumnAlignment(t.alignment,  totalSlice[i],columnIndex)
			}

		}

	}

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
	columnIndex := checkSliceIndex(t.columnName, tableHeader)
	for j := 0; j <= 1; j++ {
		//此for循环主要用于打印"------"
		for i := 0; i < len(tableHeader); i++ {
			c := fmt.Sprintf("%s%-*s", plussign, 24, Horizontalbar)
			d := strings.Replace(c, " ", Horizontalbar, -1)
			if i == (len(tableHeader) - 1) {
				d = d + plussign
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
				setTableHeaderColumnAlignment(tableHeader, t.alignment, columnIndex)
			}
		}
	}
}
