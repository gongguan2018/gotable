# gotable
Golang Table Library
这是一个使用Golang实现的表格库,可以实现一些简单的表格展示,功能如下:
1、表格内容换行： 当表格中文本内容超过20个字符的时候,将自动换行
2、表格内容对齐方式, 如下:
(1)、SetAlignmentMode(0) 表示左对齐
(2)、SetAlignmentMode(1) 表示右对齐
(3)、SetAlignmentMode(2) 表示剧中对齐
(4)、SetAlignmentMode()  什么都不设置默认表示剧中对齐
3、设置表格中要对齐的列: 
(1)、SetAlignmentColumn()                       什么都不输入默认表示设置全部列对齐
(2)、SetAlignmentColumn("MemTotal","all")       包含"all",默认也是设置全部列对齐
(3)、SetAlignmentColumn("all")                  只输入"all",默认也是设置全部列对齐
(4)、SetAlignmentColumn("MemTotal","MemFree")   设置MemTotal、MemFree对齐
4、设置表格中要设置颜色的列
(1)、SetColorColumn()                           什么都不输入默认表示设置全部列的颜色
(2)、SetColorColumn("all")                      只输入"all",表示设置全部列的颜色
(3)、SetColorColumn("MemTotal","all")           包含"all", 表示设置全部列的颜色
(4)、SetColorColumn("MemTotal","MemFree")       设置MemTotal, MemFree列的颜色
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
