package table

import (
	"fmt"
)

/*
   **************************************************
   * Date:   2024-04-09                             *
   * Author: gongguan                               *
   * Email:  1542345123@qq.com                      *
   **************************************************
*/

/*
    设置全部列的对齐方式
    当只有表头或者表头和表内容都存在的时候,都是调用此函数
*/
func setAllColumnAlign(align, foregroundcolor, backgroundcolor int, tableAll []string, colorColumnIndex []int) {
	var tdslice [][]string
	tds := croppSlice(tdslice, tableAll) //裁剪切片,产过20个长度的就被裁剪为一个新切片
	res := checkrepeat(tds)              //去除而且切片中每个切片中重复的元素
	for _, v := range res {
		for k, v1 := range v {
			e := ""
			//颜色列索引切片长度不等于0,说明有要设置颜色的列
			if len(colorColumnIndex) != 0 {
				//检查当前循环中的k 是否在要设置颜色列的索引中
				b := checkIndex(k, colorColumnIndex)
				switch align {
				case leftAlignment:
					//b为true,说明需要设置颜色
					if b {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, foregroundcolor, backgroundcolor, v1, 4, " ")
						} else {
							e = fmt.Sprintf("%s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, foregroundcolor, backgroundcolor, v1, 24-len(v1), "")
						}
					} else {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%s%-*s", verticalLine, v1, 4, " ")
						} else {
							e = fmt.Sprintf("%s%s%-*s", verticalLine, v1, 24-len(v1), "")
						}

					}
				case rightAlignment:
					if b {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m", verticalLine, 4, " ", foregroundcolor, backgroundcolor, v1)
						} else {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m", verticalLine, 24-len(v1), " ", foregroundcolor, backgroundcolor, v1)
						}
					} else {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s%s", verticalLine, 4, " ", v1)
						} else {
							e = fmt.Sprintf("%s%-*s%s", verticalLine, 24-len(v1), " ", v1)
						}
					}
				case centerAlignment:
					if b {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 2, " ")
						} else {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 22-len(v1), " ")
						}
					} else {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 2, " ")
						} else {
							e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 22-len(v1), " ")
						}
					}
				}
			} else {
				//颜色列索引为0,说明全部列都要设置颜色
				switch align {
				case leftAlignment:
					if v1 == "" {
						e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
					} else if len(v1) == 20 {
						e = fmt.Sprintf("%s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, foregroundcolor, backgroundcolor, v1, 4, " ")
					} else {
						e = fmt.Sprintf("%s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, foregroundcolor, backgroundcolor, v1, 24-len(v1), "")
					}
				case rightAlignment:
					if v1 == "" {
						e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
					} else if len(v1) == 20 {
						e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m", verticalLine, 4, " ", foregroundcolor, backgroundcolor, v1)
					} else {
						e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m", verticalLine, 24-len(v1), " ", foregroundcolor, backgroundcolor, v1)
					}
				case centerAlignment:
					if v1 == "" {
						e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
					} else if len(v1) == 20 {
						e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 2, " ")
					} else {
						e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 22-len(v1), " ")
					}

				}
			}
			if k == (len(v) - 1) {
				e = e + verticalLine
			}
			fmt.Printf(e)
		}
		fmt.Println()
	}

}

/*
    设置表中指定列的对齐方式,如果只有表头或者表头和表内容都存在的时候,都会调用此函数
*/
func setSpecifyColumnAlign(tableColumn []string, align, foregroundcolor, backgroundcolor int, columnIndex, colorColumnIndex []int) {
	var tdslice [][]string
	tds := croppSlice(tdslice, tableColumn)
	res := checkrepeat(tds)
	for _, v := range res {
		for k1, v1 := range v {
			e := ""
			//检查当前循环的索引是否在要设置的对齐的列的索引切片中
			b := checkIndex(k1, columnIndex)
			switch align {
			case leftAlignment:
				//如果b为true,说明此时的 k 就为要设置对齐的索引
				if b {
					//在判断下要设置的颜色列索引切片中是否有元素,!=0说明有元素
					if len(colorColumnIndex) != 0 {
						//检测当前循环索引 k 是否为要设置的颜色索引
						b1 := checkIndex(k1, colorColumnIndex)
						//如果b1为true,说明当前索引也是要设置的颜色列
						if b1 {
							if v1 == "" {
								e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
							} else if len(v1) == 20 {
								e = fmt.Sprintf("%s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, foregroundcolor, backgroundcolor, v1, 4, " ")
							} else {
								e = fmt.Sprintf("%s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, foregroundcolor, backgroundcolor, v1, 24-len(v1), "")
							}
						} else {
							if v1 == "" {
								e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
							} else if len(v1) == 20 {
								e = fmt.Sprintf("%s%s%-*s", verticalLine, v1, 4, " ")
							} else {
								e = fmt.Sprintf("%s%s%-*s", verticalLine, v1, 24-len(v1), "")
							}
						}
					} else {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, foregroundcolor, backgroundcolor, v1, 4, " ")
						} else {
							e = fmt.Sprintf("%s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, foregroundcolor, backgroundcolor, v1, 24-len(v1), "")
						}
					}
				} else {
					if len(colorColumnIndex) != 0 {
						b1 := checkIndex(k1, colorColumnIndex)
						if b1 {
							if v1 == "" {
								e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
							} else if len(v1) == 20 {
								e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 2, " ")
							} else {
								e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 22-len(v1), " ")
							}
						} else {
							if v1 == "" {
								e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
							} else if len(v1) == 20 {
								e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 2, " ")
							} else {
								e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 22-len(v1), " ")
							}
						}
					} else {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 2, " ")
						} else {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 22-len(v1), " ")
						}
					}
				}
			case rightAlignment:
				if b {
					if len(colorColumnIndex) != 0 {
						b1 := checkIndex(k1, colorColumnIndex)
						if b1 {
							if v1 == "" {
								e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
							} else if len(v1) == 20 {
								e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m", verticalLine, 4, " ", foregroundcolor, backgroundcolor, v1)
							} else {
								e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m", verticalLine, 24-len(v1), " ", foregroundcolor, backgroundcolor, v1)
							}
						} else {
							if v1 == "" {
								e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
							} else if len(v1) == 20 {
								e = fmt.Sprintf("%s%-*s%s", verticalLine, 4, " ", v1)
							} else {
								e = fmt.Sprintf("%s%-*s%s", verticalLine, 24-len(v1), " ", v1)
							}
						}
					} else {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, " ")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m", verticalLine, 4, " ", foregroundcolor, backgroundcolor, v1)
						} else {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m", verticalLine, 24-len(v1), " ", foregroundcolor, backgroundcolor, v1)
						}
					}
				} else {
					if len(colorColumnIndex) != 0 {
						b1 := checkIndex(k1, colorColumnIndex)
						if b1 {
							if v1 == "" {
								e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
							} else if len(v1) == 20 {
								e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 2, " ")
							} else {
								e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 22-len(v1), " ")
							}
						} else {
							if v1 == "" {
								e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
							} else if len(v1) == 20 {
								e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 2, " ")
							} else {
								e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 22-len(v1), " ")
							}
						}
					} else {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 2, " ")
						} else {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 22-len(v1), " ")
						}
					}
				}
			case centerAlignment:
				if len(colorColumnIndex) != 0 {
					b1 := checkIndex(k1, colorColumnIndex)
					if b1 {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 2, " ")
						} else {
							e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 22-len(v1), " ")
						}
					} else {
						if v1 == "" {
							e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
						} else if len(v1) == 20 {
							e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 2, " ")
						} else {
							e = fmt.Sprintf("%s%-*s%s%-*s", verticalLine, 2, " ", v1, 22-len(v1), " ")
						}
					}
				} else {
					if v1 == "" {
						e = fmt.Sprintf("%s%-*s", verticalLine, 24, "")
					} else if len(v1) == 20 {
						e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 2, " ")
					} else {
						e = fmt.Sprintf("%s%-*s\033[1;%d;%dm%s\033[0m%-*s", verticalLine, 2, " ", foregroundcolor, backgroundcolor, v1, 22-len(v1), " ")
					}
				}
			}
			if k1 == (len(v) - 1) {
				e = e + verticalLine
			}
			fmt.Printf(e)
		}
		fmt.Println()
	}
}
