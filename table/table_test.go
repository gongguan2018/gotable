package table

import (
	"testing"
)

type setalignmode struct {
	sliceint []int
	expected int
}

func TestSetAlignmentMode(t *testing.T) {
	table := &Table{}
	tsa := []setalignmode{
		setalignmode{
			sliceint: []int{},
			expected: 2,
		},
		setalignmode{
			sliceint: []int{0},
			expected: 0,
		},
		setalignmode{
			sliceint: []int{1},
			expected: 1,
		},
		setalignmode{
			sliceint: []int{2},
			expected: 2,
		},
	}
	for _, v := range tsa {
		if len(v.sliceint) == 0 {
			table.SetAlignmentMode()
			if table.alignment != v.expected {
				t.Error("Column alignment pattern mismatch!!!")
			}
		} else {
			table.SetAlignmentMode(v.sliceint[0])
			if table.alignment != v.expected {
				t.Error("Column alignment pattern mismatch!!!")
			}
		}
	}
}

type setaligncolumn struct {
	slicestr []string
	expected []string
}

func TestSetAlignmentColumn(t *testing.T) {
	table := &Table{
	  tableHeader : []string{"MemTotal", "MemUsedhelloworldchina", "Memfree", "SwapTotal", "SwapUsed", "SwapFree"},
	}
	tsa := []setaligncolumn{
		setaligncolumn{
			slicestr: []string{},
			expected: []string{"all"},
		},
		setaligncolumn{
			slicestr: []string{"test1", "all"},
			expected: []string{"all"},
		},
		setaligncolumn{
			slicestr: []string{"MemTotal", "SwapFree", "Memfree"},
			expected: []string{"MemTotal", "SwapFree", "Memfree"},
		},
	}
	for _, v := range tsa {
		if len(v.slicestr) == 0 {
			table.SetAlignmentColumn()
			if table.columnName[0] != v.expected[0] {
				t.Error("Column Name error!!!")
			}
		} else if checkColumnName(v.slicestr) {
			table.SetAlignmentColumn(v.slicestr...)
			if table.columnName[0] != v.expected[0] {
				t.Error("Column Name error!!!")
			}
		}else {
                       table.SetAlignmentColumn(v.slicestr...)
		       r := checksliceequal(table.columnName,v.expected)
		       if r != 0 {
                            t.Error("Column set error!!!")
		       }
		}
	}
}
