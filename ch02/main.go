package main

import "fmt"

func main() {
	//var a int8
	//var b int16
	//var c int32
	//var d int64
	//
	//var ua uint8
	//var ub uint16
	//var uc uint32
	//var ud uint64
	//var e int //动态类型 根据操作系统定义的
	//
	//a = int8(b)
	//
	//var f1 float32 // 大约是3.4e38
	//var f2 float64 // 1.8e308
	//
	//f1 = 3
	//f2 = 3.14

	var c byte // 主要是用于存放字符的,存放字符串 阿斯克码
	c = 'a'
	c1 := 97
	fmt.Println(c)
	fmt.Printf("c=%c", c)
	fmt.Printf("\nc1=%c", c1)

	var c2 rune // 也是字符 国际化语言字符 中文韩文等等
	c2 = '穆'
	fmt.Printf("\nc2=%c", c2)

	var name string
	name = "wwww.xmaven.cn"
	fmt.Println(name)
}
