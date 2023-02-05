package main

import (
	"fmt"
	"strconv"
)

type MyInt int // 自定义类型，基于已有的类型自定义的一个类型

// 绑定在MyInt类型上面的string方法，string
func (mi MyInt) string() string {
	return strconv.Itoa(int(mi))
}

func main() {
	// type 关键字
	/*
		1.定义结构体
		2.定义接口
		3.定义类型别名
		4.类型定义
		5.类型判断
	*/
	// 别名实际上是为了更好的理解代码
	/*
		type MyInt = int // 类型的别名
		var i MyInt = 12
		var j int = 9
		fmt.Println(i + j) // 在编译的时候 类型别名会直接替换成int
		fmt.Printf("%T", i)
	*/

	var a interface{} = "abc"
	switch a.(type) {
	case string:
		fmt.Println("字符串")
	}

	var i MyInt = 12
	var j int = 9 // 既希望你是int类型，有希望你能有方法
	fmt.Println(i.string())
	fmt.Println(int(i) + j)
	fmt.Printf("%T\r\n", i)

}
