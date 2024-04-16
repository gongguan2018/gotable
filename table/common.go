package table

import (
	"strings"
)

/*
   **************************************************
   * Date:   2024-04-09                             *
   * Author: gongguan                               *
   * Email:  1542345123@qq.com                      *
   **************************************************
*/

/*
    换行原理就是将切片中超过20个长度的部分裁剪出来,如果其余元素没超过,那么新的切片就包含空的元素和裁剪下来
    的超出的元素,也就是将一个切片变成了两个切片,一起放在二维切片里,分别打印出来即可
    将slice切片循环裁剪,如果元素的长度超过20,那么截取前20个元素,并将前20个元素替换为空字符串
    然后在修改此元素将空字符串去掉
    eg: ["hello","beijingshanghaiguangdong","world"]经过croppSlice后将被裁剪为二维切片
    [["hello","beijingshanghaiguang","world"],["hello","dong","world" ]] 
    
*/
func croppSlice(reslice [][]string, slice []string) [][]string {
	var newslice []string
	newslice = make([]string, len(slice))
//	newmap := make(map[int]int)
	for k, v := range slice {
		if len(v) > 20 {
			newslice[k] = v[0:20]
			s := strings.Replace(slice[k], v[0:20], "", 1)
			slice[k] = strings.TrimSpace(s)
		} else {
			newslice[k] = v
			slice[k] = ""
		}
//		if newmap[k] == 0 {
//			newmap[k] = len(slice[k]) //将截取后的slice切片长度添加到map中
//		}

	}
	reslice = append(reslice, newslice)
	b := checkSliceLength(slice) //检查slice切片中是否还有长度大于20的元素
	if b == 0 {                  //等0,说明切片中已经没有大于20的元素,返回二维切片reslice
		return reslice
	}
	return croppSlice(reslice, slice) //只有满足了上面的if循环,此递归函数才不会执行,否则一直调用
}

/*
    检查slice切片长度,如果切片中的元素小于20,那么就设置map中的值为0(key为切片元素对应的索引)
    如果大于20,那么设置map的值为1
    遍历map,求map中所有值的和,如果和为0,说明每个值都是0,进而说明切片中元素全部都是小于20
    如果和大于0,说明切片中还有元素长度大于20
*/
func checkSliceLength(slice []string) int {
	a := make([]int, 100)
	for k, v := range slice {
		if len(v) == 0 {
			a[k] = 0
		} else {
			a[k] = 1
		}
	}
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

/*
    检查二维切片中的每个切片是否有重复的元素,如果有重复,将后面重复的元素修改为空字符串
    经过checkrepeat处理后二维切片[["hello","beijingshanghaiguang","world"],["hello","dong","world" ]]
    将变成[["hello","beijingshanghaiguang","world"],["","dong","" ]],将这个新切片逐行打印即可实现换行
*/
func checkrepeat(slice [][]string) [][]string {
	newmap := make(map[string]bool)
	for _, v := range slice {
		for k, v1 := range v {
			if !newmap[v1] {
				newmap[v1] = true
			} else {
				v[k] = ""
			}
		}
	}
	return slice
}

// 检查一个切片元素在另一个切片中索引,返回int切片
func checkSliceIndex(column, tableHeader []string) []int {
	mapp := make(map[string]int)
	var newslice []int
	for k, v := range tableHeader {
		_, ok := mapp[v]
		if !ok {
			mapp[v] = k
		}
	}
	for _, v := range column {
		k, ok := mapp[v]
		if ok {
			newslice = append(newslice, k)
		}
	}
	return newslice
}
//检查两个切片是否相等
func checksliceequal(a, b []string) int {
	if len(a) != len(b) {
		return 2
	} else {
		maps := make(map[string]bool)
		for _, v := range a {
			if !maps[v] {
				maps[v] = true
			}
		}
		sum := 0
		for _, v := range b {
			if maps[v] {
				sum += 0
			} else {
				sum += 1
			}
		}
		return sum
	}
}

func checkintsliceequal(a, b []int) int {
        if len(a) != len(b) {
                return 2
        } else {
                maps := make(map[int]bool)
                for _, v := range a {
                        if !maps[v] {
                                maps[v] = true
                        }
                }
                sum := 0
                for _, v := range b {
                        if maps[v] {
                                sum += 0
                        } else {
                                sum += 1
                        }
                }
                return sum
        }
}
